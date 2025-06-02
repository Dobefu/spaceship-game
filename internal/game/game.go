package game

import (
	"log"

	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/Dobefu/spaceship-game/internal/scenes/game_scene"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	game       = &Game{}
	gameHeight = 640
	gameWidth  = 640
)

type IGame interface {
	SetScene(scene *scene.IScene)
}

type Game struct {
	IGame

	scene scene.IScene
	input input.Input
}

func (g *Game) Update() (err error) {
	g.input.Update()

	if g.scene == nil {
		return nil
	}

	gameObjects := g.scene.GetGameObjects()

	for _, gameObject := range gameObjects {
		err = gameObject.Update()

		if err != nil {
			return err
		}
	}

	camera := g.scene.GetCamera()
	camera.Update()

	// TODO: Not this.
	g.scene.SetCamera(camera)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}

	gameObjects := g.scene.GetGameObjects()
	camera := g.scene.GetCamera()

	for _, gameObject := range gameObjects {
		gameObject.Draw(screen, camera.GetPosition())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameHeight, gameWidth
}

func Run() {
	ebiten.SetWindowSize(gameHeight, gameWidth)
	ebiten.SetWindowTitle("Spaceship Game")

	game.input = input.GlobalInput
	game.SetScene(&game_scene.GameScene{})

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}
