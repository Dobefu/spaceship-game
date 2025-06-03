package game

import (
	"github.com/Dobefu/spaceship-game/internal/camera"
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (g *Game) SetScene(scene interfaces.Scene) {
	g.scene = scene

	g.scene.SetCamera(&camera.Camera{})
	g.scene.Init()
}
