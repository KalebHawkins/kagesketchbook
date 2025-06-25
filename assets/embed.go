package assets

import "embed"

//go:embed shaders/*.kage
var ShaderFS embed.FS

func LoadShader(fs embed.FS, file string) ([]byte, error) {
	return fs.ReadFile("shaders/" + file)
}
