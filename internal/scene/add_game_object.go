package scene

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
)

func (s *Scene) AddGameObject(gameObject *game_object.GameObject) {
	s.gameObjects = append(s.gameObjects, *gameObject)
}
