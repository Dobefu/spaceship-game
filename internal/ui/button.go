package ui

import (
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type buttonState int

const (
	buttonStateNormal = iota
	buttonStateHover
	buttonStateActive
)

type Button struct {
	game_object.GameObject

	width  float32
	height float32
	halign text.Align
	valign text.Align

	text string

	state buttonState
}

func NewButton(
	position vectors.Vector2,
	width float32,
	height float32,
	text string,
	halign text.Align,
	valign text.Align,
) *Button {
	button := &Button{
		width:  width,
		height: height,
		halign: halign,
		valign: valign,
		text:   text,
	}

	button.SetPosition(position.ToVector3())
	button.SetIsActive(true)

	return button
}

func (b *Button) Update() (err error) {
	cx, cy := ebiten.CursorPosition()
	position := b.GetPosition()

	b.state = buttonStateNormal

	if float64(cx) >= position.X &&
		float64(cx) < position.X+float64(b.width) &&
		float64(cy) >= position.Y &&
		float64(cy) < position.Y+float64(b.height) {

		b.state = buttonStateHover

		if inpututil.MouseButtonPressDuration(ebiten.MouseButton0) > 0 {
			b.state = buttonStateActive
		}
	}
	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	position := b.GetPosition()
	var bg uint8 = 100

	if b.state != buttonStateNormal {
		bg = 90
	}

	posX := float32(position.X)
	posY := float32(position.Y)

	if b.halign == text.AlignCenter {
		posX -= b.width / 2
	}
	if b.valign == text.AlignCenter {
		posY -= b.height / 2
	}
	if b.halign == text.AlignEnd {
		posX -= b.width
	}
	if b.valign == text.AlignEnd {
		posY -= b.height
	}

	vector.DrawFilledRect(
		screen,
		posX,
		posY,
		b.width,
		b.height,
		color.Gray{Y: bg},
		false,
	)

	op := &text.DrawOptions{}
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	op.GeoM.Translate(float64(posX+(b.width/2)+2), float64(posY+(b.height/2)+2))

	if b.state != buttonStateActive {
		op.ColorScale.ScaleWithColor(color.Black)

		text.Draw(
			screen,
			b.text,
			&text.GoTextFace{
				Source: fontSrc,
				Size:   16,
			},
			op,
		)

		op.GeoM.Translate(-2, -2)
	}
	op.ColorScale.Reset()
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(
		screen,
		b.text,
		&text.GoTextFace{
			Source: fontSrc,
			Size:   16,
		},
		op,
	)
}
