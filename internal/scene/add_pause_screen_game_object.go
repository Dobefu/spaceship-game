package scene

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (s *Scene) AddPauseScreenGameObject(gameObject interfaces.GameObject) {
	s.pauseScreenGameObjects = append(s.pauseScreenGameObjects, gameObject)
}
