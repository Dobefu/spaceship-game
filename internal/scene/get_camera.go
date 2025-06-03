package scene

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (s *Scene) GetCamera() (camera interfaces.Camera) {
	return s.Camera
}
