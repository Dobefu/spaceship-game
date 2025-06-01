package game

import (
	"github.com/Dobefu/spaceship-game/internal/camera"
	"github.com/Dobefu/spaceship-game/internal/scene"
)

func (g *Game) SetScene(scene scene.IScene) {
	g.scene = scene
	g.scene.SetCamera(camera.Camera{})
	g.scene.Init()
}
