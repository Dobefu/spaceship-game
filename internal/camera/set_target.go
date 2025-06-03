package camera

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/options"
)

func (c *Camera) SetTarget(target game_object.IGameObject) {
	c.target = target

	gameHeight := float64(options.GlobalOptions.Height)
	gameWidth := float64(options.GlobalOptions.Width)

	c.position.X = c.target.GetPosition().X - (gameWidth / 2)
	c.position.Y = c.target.GetPosition().Y - (gameHeight / 2)
}
