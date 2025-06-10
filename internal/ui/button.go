package ui

import (
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	game_object.GameObject

	width  float32
	height float32
}

func NewButton(position vectors.Vector2, width float32, height float32) *Button {
	button := &Button{
		width:  width,
		height: height,
	}

	button.SetPosition(position.ToVector3())
	button.SetIsActive(true)

	return button
}

func (b *Button) Update() (err error) {
	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	position := b.GetPosition()

	vector.DrawFilledRect(
		screen,
		float32(position.X),
		float32(position.Y),
		b.width,
		b.height,
		color.Gray{Y: 100},
		false,
	)

	op := &text.DrawOptions{}
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	op.GeoM.Translate(position.X+float64(b.width/2)+2, position.Y+float64(b.height/2)+2)
	op.ColorScale.ScaleWithColor(color.Black)

	text.Draw(
		screen,
		"TEST",
		&text.GoTextFace{
			Source: fontSrc,
			Size:   16,
		},
		op,
	)

	op.GeoM.Translate(-2, -2)
	op.ColorScale.Reset()
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(
		screen,
		"TEST",
		&text.GoTextFace{
			Source: fontSrc,
			Size:   16,
		},
		op,
	)
}
