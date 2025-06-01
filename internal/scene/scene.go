package scene

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
)

type IScene interface {
	Init()
	GetGameObjects() []game_object.IGameObject
	AddGameObject(gameObject game_object.IGameObject)
}

type Scene struct {
	IScene

	gameObjects []game_object.IGameObject
}
