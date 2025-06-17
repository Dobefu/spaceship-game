package scene

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (s *Scene) GetPauseScreenGameObjects() []interfaces.GameObject {
	return s.pauseScreenGameObjects
}
