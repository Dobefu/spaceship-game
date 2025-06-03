package camera

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (c *Camera) WorldToScreenPosition(world vectors.Vector2) vectors.Vector2 {
	return vectors.Vector2{
		X: world.X - c.position.X,
		Y: world.Y - c.position.Y,
	}
}
