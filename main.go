package main

import (
	"log"

	"github.com/KalebHawkins/kageskestchbook/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	cfg := &engine.Config{
		ScreenWidth:  640,
		ScreenHeight: 480,
		WindowTitle:  "Kage Sketchbook",
	}

	if err := ebiten.RunGame(engine.NewGame(cfg)); err != nil {
		log.Fatal(err)
	}
}
