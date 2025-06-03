package game_object

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (g *GameObject) GetPosition() (position vectors.Vector3) {
	return g.position
}
