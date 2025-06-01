package scene

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
)

type scene interface {
	AddGameObject(gameObject *game_object.GameObject)
}

type Scene struct {
	scene

	gameObjects []game_object.GameObject
}
