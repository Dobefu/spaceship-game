package game

import (
	"github.com/Dobefu/spaceship-game/internal/camera"
	"github.com/Dobefu/spaceship-game/internal/scene"
)

func (g *Game) SetScene(scene scene.IScene) {
	g.scene = scene
	g.scene.Init()
	g.scene.SetCamera(camera.Camera{})
}
