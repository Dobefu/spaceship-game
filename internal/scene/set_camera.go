package scene

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (s *Scene) SetCamera(camera interfaces.Camera) {
	s.Camera = camera
}
