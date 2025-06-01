package scene

import "github.com/Dobefu/spaceship-game/internal/camera"

func (s *Scene) GetCamera() (camera camera.Camera) {
	return s.Camera
}
