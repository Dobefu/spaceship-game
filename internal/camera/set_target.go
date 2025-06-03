package camera

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (c *Camera) SetTarget(target interfaces.GameObject) {
	c.target = target

	gameHeight := float64(globals.GlobalValues.Height)
	gameWidth := float64(globals.GlobalValues.Width)

	c.position.X = c.target.GetPosition().X - (gameWidth / 2)
	c.position.Y = c.target.GetPosition().Y - (gameHeight / 2)
}
