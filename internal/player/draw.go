package player

import (
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	blackImage = ebiten.NewImage(1, 1)
	whiteImage = ebiten.NewImage(1, 1)
	redImage   = ebiten.NewImage(1, 1)

	strokeOptions        *vector.StrokeOptions
	drawTrianglesOptions *ebiten.DrawTrianglesOptions
)

func init() {
	blackImage.Set(0, 0, color.Black)
	whiteImage.Set(0, 0, color.White)
	redImage.Set(0, 0, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	strokeOptions = &vector.StrokeOptions{
		Width: globals.GlobalValues.OutlineWidth,
	}

	drawTrianglesOptions = &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.drawFireModel(screen)
	p.drawPlayerModel(screen)
}
