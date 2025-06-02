package scene

import (
	"github.com/Dobefu/spaceship-game/internal/camera"
	"github.com/Dobefu/spaceship-game/internal/game_object"
)

type IScene interface {
	Init()
	GetGameObjects() []game_object.IGameObject
	AddGameObject(gameObject game_object.IGameObject)
	SetCamera(camera camera.Camera)
	GetCamera() (camera *camera.Camera)
}

type Scene struct {
	IScene

	Camera      camera.Camera
	gameObjects []game_object.IGameObject
}
