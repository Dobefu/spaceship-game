package scene

import (
	"sort"

	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (s *Scene) GetGameObjects() []interfaces.GameObject {
	if s.isDepthOrderDirty {
		sort.Slice(s.gameObjects, func(i int, j int) bool {
			return s.gameObjects[i].GetPosition().Z < s.gameObjects[j].GetPosition().Z
		})

		s.isDepthOrderDirty = false
	}

	return s.gameObjects
}
