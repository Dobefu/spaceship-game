package camera

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
)

func (c *Camera) Update() {
	if c.target == nil {
		return
	}

	gameHeight := float64(globals.GlobalValues.Height)
	gameWidth := float64(globals.GlobalValues.Width)

	targetX := c.target.GetPosition().X - (gameWidth / 2)
	targetY := c.target.GetPosition().Y - (gameHeight / 2)

	c.position.X += (targetX - c.position.X) / 10
	c.position.Y += (targetY - c.position.Y) / 10
}
