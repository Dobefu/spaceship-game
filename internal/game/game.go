package game

import (
	"log"

	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	game       = &Game{}
	gameHeight = 640
	gameWidth  = 640
)

type Game struct {
	scene scene.Scene
}

func (g *Game) Update() (err error) {
	gameObjects := g.scene.GetGameObjects()

	for _, gameObject := range gameObjects {
		err = gameObject.Update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gameObjects := g.scene.GetGameObjects()

	for _, gameObject := range gameObjects {
		gameObject.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameHeight, gameWidth
}

func Run() {
	ebiten.SetWindowSize(gameHeight, gameWidth)
	ebiten.SetWindowTitle("Spaceship Game")

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}
