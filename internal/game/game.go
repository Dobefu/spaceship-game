package game

import (
	"log"

	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/options"
	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/Dobefu/spaceship-game/internal/scenes/game_scene"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	game = &Game{}
)

type IGame interface {
	SetScene(scene *scene.IScene)
}

type Game struct {
	IGame

	scene scene.IScene
}

func (g *Game) Update() (err error) {
	input.GlobalInput.Update()

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
	return options.GlobalOptions.Height, options.GlobalOptions.Width
}

func Run() {
	ebiten.SetWindowSize(options.GlobalOptions.Height, options.GlobalOptions.Width)
	ebiten.SetWindowTitle("Spaceship Game")

	game.SetScene(&game_scene.GameScene{})

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}
