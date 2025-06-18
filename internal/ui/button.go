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

	text   string
	action func(b *Button)

	state buttonState
}

func NewButton(
	position vectors.Vector2,
	width float32,
	height float32,
	halign text.Align,
	valign text.Align,
	text string,
	action func(b *Button),
) *Button {
	button := &Button{
		width:  width,
		height: height,
		halign: halign,
		valign: valign,
		text:   text,
		action: action,
	}

	button.SetPosition(position.ToVector3())
	button.SetIsActive(true)

	return button
}

func (b *Button) Update() (err error) {
	cx, cy := ebiten.CursorPosition()

	position := b.getAlignedPosition()
	posX := position.X
	posY := position.Y

	if float64(cx) >= posX &&
		float64(cx) < posX+float64(b.width) &&
		float64(cy) >= posY &&
		float64(cy) < posY+float64(b.height) {

		if b.state == buttonStateActive && inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			b.action(b)
		}

		b.state = buttonStateHover

		if inpututil.MouseButtonPressDuration(ebiten.MouseButton0) > 0 {
			b.state = buttonStateActive
		}

		return
	}

	b.state = buttonStateNormal
	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	var bg uint8 = 100

	if b.state != buttonStateNormal {
		bg = 90
	}

	position := b.getAlignedPosition()
	posX := float32(position.X)
	posY := float32(position.Y)

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

func (b *Button) getAlignedPosition() vectors.Vector2 {
	position := b.GetPosition()

	posX := position.X
	posY := position.Y

	if b.halign == text.AlignCenter {
		posX -= float64(b.width / 2)
	}
	if b.valign == text.AlignCenter {
		posY -= float64(b.height / 2)
	}
	if b.halign == text.AlignEnd {
		posX -= float64(b.width)
	}
	if b.valign == text.AlignEnd {
		posY -= float64(b.height)
	}

	return vectors.Vector2{X: posX, Y: posY}
}

func (b *Button) SetText(text string) {
	b.text = text
}
