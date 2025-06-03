package bullet

import (
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/options"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (b *Bullet) Draw(screen *ebiten.Image, offset vectors.Vector2) {
	// TODO: Move this to the GameObject struct.
	if !b.isActive {
		return
	}

	vector.StrokeCircle(
		screen,
		float32(b.GetPosition().X-offset.X),
		float32(b.GetPosition().X-offset.Y),
		5,
		options.GlobalOptions.OutlineWidth,
		color.White,
		true,
	)
}
