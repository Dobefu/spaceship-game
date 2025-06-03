package game_object

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (g *GameObject) SetPosition(position vectors.Vector3) {
	if g.position.Z == position.Z {
		g.position = position
		return
	}

	g.position = position

	if globals.GlobalValues.Game == nil || globals.GlobalValues.Game.GetScene() == nil {
		return
	}

	globals.GlobalValues.Game.GetScene().InvalidateDepthOrder()
}
