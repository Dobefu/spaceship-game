package scene

import (
	"github.com/Dobefu/spaceship-game/internal/camera"
	"github.com/Dobefu/spaceship-game/internal/game_object"
)

type IScene interface {
	Init()
	GetGameObjects() []game_object.IGameObject
	AddGameObject(gameObject game_object.IGameObject)
	SetCamera(camera camera.ICamera)
	GetCamera() (camera camera.ICamera)
}

type Scene struct {
	IScene

	camera      camera.ICamera
	gameObjects []game_object.IGameObject
}
