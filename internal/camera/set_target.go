package camera

import "github.com/Dobefu/spaceship-game/internal/game_object"

func (c *Camera) SetTarget(target game_object.IGameObject) {
	c.target = target
}
