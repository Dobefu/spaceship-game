package camera

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (c *Camera) IsPositionWithinBounds(pos vectors.Vector2, margin float64) bool {
	gameHeight := float64(globals.GlobalValues.Height)
	gameWidth := float64(globals.GlobalValues.Width)

	return pos.X < -margin ||
		pos.X > gameWidth+margin ||
		pos.Y < -margin ||
		pos.Y > gameHeight+margin
}
