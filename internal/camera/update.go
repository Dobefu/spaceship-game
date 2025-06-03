package camera

import "github.com/Dobefu/spaceship-game/internal/options"

func (c *Camera) Update() {
	if c.target == nil {
		return
	}

	gameHeight := float64(options.GlobalOptions.Height)
	gameWidth := float64(options.GlobalOptions.Width)

	targetX := c.target.GetPosition().X - (gameWidth / 2)
	targetY := c.target.GetPosition().Y - (gameHeight / 2)

	c.position.X += (targetX - c.position.X) / 10
	c.position.Y += (targetY - c.position.Y) / 10
}
