package shaders

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
)

var (
	//go:embed bloom.go
	bloomSrc []byte
	Bloom    *ebiten.Shader
)

func init() {
	var err error
	Bloom, err = ebiten.NewShader(bloomSrc)

	if err != nil {
		log.Fatal(err)
	}
}
