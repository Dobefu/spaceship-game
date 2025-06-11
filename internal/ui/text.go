package ui

import (
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	game_object.GameObject

	halign text.Align
	valign text.Align

	text     string
	fontSize float64
}

func NewText(
	position vectors.Vector2,
	halign text.Align,
	valign text.Align,
	text string,
	fontSize float64,
) *Text {
	textObj := &Text{
		halign:   halign,
		valign:   valign,
		text:     text,
		fontSize: fontSize,
	}

	textObj.SetPosition(position.ToVector3())
	textObj.SetIsActive(true)

	return textObj
}

func (t *Text) Update() (err error) {
	return nil
}

func (t *Text) Draw(screen *ebiten.Image) {
	position := t.GetPosition()
	posX := float32(position.X)
	posY := float32(position.Y)

	op := &text.DrawOptions{}
	op.PrimaryAlign = t.halign
	op.SecondaryAlign = t.valign
	op.GeoM.Translate(float64(posX), float64(posY))
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(
		screen,
		t.text,
		&text.GoTextFace{
			Source: fontSrc,
			Size:   t.fontSize,
		},
		op,
	)
}
