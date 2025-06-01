package scene

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
)

type scene interface {
	GetGameObjects() []game_object.GameObject
	AddGameObject(gameObject *game_object.GameObject)
}

type Scene struct {
	scene

	gameObjects []game_object.GameObject
}
