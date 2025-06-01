package scene

import (
	"github.com/Dobefu/spaceship-game/internal/camera"
)

func (s *Scene) SetCamera(camera camera.ICamera) {
	s.camera = camera
}
