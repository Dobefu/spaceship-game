package scene

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
)

func (s *Scene) AddGameObject(gameObject game_object.IGameObject) {
	s.gameObjects = append(s.gameObjects, gameObject)
}
