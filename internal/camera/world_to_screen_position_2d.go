package camera

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (c *Camera) WorldToScreenPosition2D(worldPos vectors.Vector3) vectors.Vector2 {
	if worldPos.Z >= 0 {
		return vectors.Vector2{
			X: (worldPos.X - c.position.X),
			Y: (worldPos.Y - c.position.Y),
		}
	}

	scale := 1 / (1 + math.Abs(worldPos.Z)*0.5)

	return vectors.Vector2{
		X: worldPos.X - (c.position.X * scale),
		Y: worldPos.Y - (c.position.Y * scale),
	}
}
