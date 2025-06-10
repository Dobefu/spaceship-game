package main_menu_scene

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/Dobefu/spaceship-game/internal/ui"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type MainMenuScene struct {
	scene.Scene
}

func (s *MainMenuScene) Init() {
	s.AddGameObject(
		ui.NewButton(
			vectors.Vector2{
				X: (float64(globals.GlobalValues.Width) / 2) - 75,
				Y: (float64(globals.GlobalValues.Height) / 2) - 25,
			},
			150,
			50,
		),
	)
}
