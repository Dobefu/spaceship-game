package scene

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

type Scene struct {
	interfaces.Scene

	Camera      interfaces.Camera
	gameObjects []interfaces.GameObject
}
