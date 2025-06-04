package player

import (
	"image"
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	blackImage    = ebiten.NewImage(3, 3)
	blackSubImage = blackImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	strokeOptions        *vector.StrokeOptions
	drawTrianglesOptions *ebiten.DrawTrianglesOptions
)

func init() {
	blackImage.Fill(color.Black)
	whiteImage.Fill(color.White)
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
