package engine

import (
	"fmt"
	"log"
	"time"

	"github.com/KalebHawkins/kageskestchbook/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Config struct {
	ScreenWidth  int
	ScreenHeight int
	WindowTitle  string
}

type Game struct {
	Config      *Config
	Shaders     map[string]*ebiten.Shader
	ShaderKeys  []string
	StartTime   time.Time
	idx         int
	Uniforms    map[string]interface{}
	ToggleState float32
}

func NewGame(cfg *Config) *Game {
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.WindowTitle)

	g := &Game{
		Config:     cfg,
		StartTime:  time.Now(),
		Shaders:    map[string]*ebiten.Shader{},
		ShaderKeys: []string{},
		Uniforms:   map[string]interface{}{},
	}

	if err := g.loadShaders(); err != nil {
		log.Fatal(err)
	}

	return g
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		g.idx = (g.idx + 1) % len(g.Shaders)
	}

	g.Uniforms["Time"] = time.Since(g.StartTime).Seconds()
	g.Uniforms["Toggle"] = g.ToggleState

	mx, my := ebiten.CursorPosition()
	g.Uniforms["Mouse"] = []float32{float32(mx), float32(my)}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.ToggleState == 1.0 {
			g.ToggleState = 0.0
		} else {
			g.ToggleState = 1.0
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = g.Uniforms

	op.GeoM.Scale(float64(g.Config.ScreenWidth), float64(g.Config.ScreenHeight))
	screen.DrawRectShader(g.Config.ScreenWidth, g.Config.ScreenHeight, g.Shaders[g.ShaderKeys[g.idx]], op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Press 'Tab' to shift through shaders...\nShader File: %s", g.ShaderKeys[g.idx]))
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func (g *Game) loadShaders() error {
	shaderFiles, err := assets.ShaderFS.ReadDir("shaders")
	if err != nil {
		return err
	}

	for _, sf := range shaderFiles {
		shaderData, err := assets.LoadShader(assets.ShaderFS, sf.Name())
		if err != nil {
			return err
		}

		shader, err := ebiten.NewShader(shaderData)
		if err != nil {
			return err
		}

		g.Shaders[sf.Name()] = shader
		g.ShaderKeys = append(g.ShaderKeys, sf.Name())
	}

	return nil
}
