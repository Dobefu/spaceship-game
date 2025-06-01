package scene

import "github.com/Dobefu/spaceship-game/internal/game_object"

func (s *Scene) GetGameObjects() []game_object.GameObject {
	return s.gameObjects
}
