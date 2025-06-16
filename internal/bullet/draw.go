package bullet

import (
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (b *Bullet) Draw(screen *ebiten.Image) {
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	screenPos := camera.WorldToScreenPosition2D(b.GetPosition())

	vector.StrokeCircle(
		screen,
		float32(screenPos.X),
		float32(screenPos.Y),
		5,
		globals.GlobalValues.OutlineWidth,
		color.White,
		true,
	)
}
