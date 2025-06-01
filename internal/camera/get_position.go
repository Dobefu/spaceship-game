package camera

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (c *Camera) GetPosition() (position vectors.Vector2) {
	return c.position
}
