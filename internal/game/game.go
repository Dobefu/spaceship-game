package game

import (
	"log"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/interfaces"
	"github.com/Dobefu/spaceship-game/internal/scenes"
	"github.com/Dobefu/spaceship-game/internal/shaders"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	interfaces.Game

	scene interfaces.Scene
}

func (g *Game) Update() (err error) {
	input.GlobalInput.Update()

	if g.scene == nil {
		return nil
	}

	if g.scene.GetCanPause() && input.GlobalInput.GetButtonStart().IsPressed {
		g.scene.SetIsPaused(!g.scene.GetIsPaused())
	}

	var gameObjects []interfaces.GameObject

	if g.scene.GetIsPaused() {
		gameObjects = g.scene.GetPauseScreenGameObjects()
	} else {
		gameObjects = g.scene.GetGameObjects()
	}

	for _, gameObject := range gameObjects {
		if !gameObject.GetIsActive() {
			continue
		}

		err = gameObject.Update()

		if err != nil {
			return err
		}
	}

	if !g.scene.GetIsPaused() {
		camera := g.scene.GetCamera()
		camera.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}

	gameObjects := g.scene.GetGameObjects()

	if g.scene.GetIsPaused() {
		gameObjects = append(gameObjects, g.scene.GetPauseScreenGameObjects()...)
	}

	for _, gameObject := range gameObjects {
		if !gameObject.GetIsActive() {
			continue
		}

		gameObject.Draw(screen)
	}
}

func (g *Game) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	shaderOptions := &ebiten.DrawRectShaderOptions{}
	shaderOptions.Images[0] = offscreen
	shaderOptions.GeoM = geoM

	os := offscreen.Bounds().Size()

	screen.DrawRectShader(
		os.X,
		os.Y,
		shaders.Bloom,
		shaderOptions,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func Run() {
	ebiten.SetWindowSize(globals.GlobalValues.Width, globals.GlobalValues.Height)
	ebiten.SetWindowTitle("Spaceship Game")

	globals.GlobalValues.Game = &Game{}
	globals.GlobalValues.Game.SetScene(&scenes.MainMenuScene{})

	gameOptions := &ebiten.RunGameOptions{
		DisableHiDPI: true,
	}

	err := ebiten.RunGameWithOptions(globals.GlobalValues.Game, gameOptions)

	if err != nil {
		log.Fatal(err)
	}
}
