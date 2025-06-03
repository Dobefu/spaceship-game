package scene

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (s *Scene) AddGameObject(gameObject interfaces.GameObject) {
	s.gameObjects = append(s.gameObjects, gameObject)
	s.InvalidateDepthOrder()
}
