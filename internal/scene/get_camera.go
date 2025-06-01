package scene

import "github.com/Dobefu/spaceship-game/internal/camera"

func (s *Scene) GetCamera() (camera camera.ICamera) {
	return s.camera
}
