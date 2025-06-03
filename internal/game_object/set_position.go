package game_object

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (g *GameObject) SetPosition(position vectors.Vector2) {
	g.position = position
}
