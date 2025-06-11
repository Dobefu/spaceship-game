package scenes

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/Dobefu/spaceship-game/internal/ui"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type MainMenuScene struct {
	scene.Scene
}

func (s *MainMenuScene) Init() {
	s.AddGameObject(
		ui.NewText(
			vectors.Vector2{
				X: float64(globals.GlobalValues.Width) / 2,
				Y: float64(globals.GlobalValues.Height) / 4,
			},
			text.AlignCenter,
			text.AlignCenter,
			"Title Screen",
			40,
		),
	)

	s.AddGameObject(
		ui.NewButton(
			vectors.Vector2{
				X: float64(globals.GlobalValues.Width) / 2,
				Y: float64(globals.GlobalValues.Height) / 2,
			},
			150,
			50,
			text.AlignCenter,
			text.AlignStart,
			"Start",
			func() {
				globals.GlobalValues.Game.SetScene(&GameScene{})
			},
		),
	)

	s.AddGameObject(
		ui.NewButton(
			vectors.Vector2{
				X: float64(globals.GlobalValues.Width) / 2,
				Y: float64(globals.GlobalValues.Height)/2 + 75,
			},
			150,
			50,
			text.AlignCenter,
			text.AlignStart,
			"Options",
			func() {
				globals.GlobalValues.Game.SetScene(&OptionsScene{})
			},
		),
	)
}
