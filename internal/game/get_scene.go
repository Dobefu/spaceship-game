package game

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (g *Game) GetScene() (scene interfaces.Scene) {
	return g.scene
}
