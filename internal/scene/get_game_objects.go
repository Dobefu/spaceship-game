package scene

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (s *Scene) GetGameObjects() []interfaces.GameObject {
	return s.gameObjects
}
