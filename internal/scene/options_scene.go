package scene

import (
	"fmt"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/ui"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type OptionsScene struct {
	Scene
}

func (s *OptionsScene) Init() {
	s.AddGameObject(
		ui.NewText(
			vectors.Vector2{
				X: float64(globals.GlobalValues.Width) / 2,
				Y: float64(globals.GlobalValues.Height) / 4,
			},
			text.AlignCenter,
			text.AlignCenter,
			"Options",
			40,
		),
	)

	s.AddGameObject(
		ui.NewButton(
			vectors.Vector2{
				X: float64(globals.GlobalValues.Width) / 2,
				Y: float64(globals.GlobalValues.Height) / 2,
			},
			250,
			50,
			text.AlignCenter,
			text.AlignStart,
			fmt.Sprintf("Shaders: %v", globals.Options.EnableShaders),
			func(b *ui.Button) {
				globals.Options.EnableShaders = !globals.Options.EnableShaders
				b.SetText(fmt.Sprintf("Shaders: %v", globals.Options.EnableShaders))

				go func() {
					var enableShaders byte = 0

					if globals.Options.EnableShaders {
						enableShaders = 1
					}

					_ = globals.DataManager.SaveObjectProp("settings", "enable-shaders", []byte{enableShaders})
				}()
			},
		),
	)

	s.AddGameObject(
		ui.NewButton(
			vectors.Vector2{
				X: float64(globals.GlobalValues.Width) / 2,
				Y: float64(globals.GlobalValues.Height)/2 + 75,
			},
			250,
			50,
			text.AlignCenter,
			text.AlignStart,
			"Back",
			func(b *ui.Button) {
				globals.GlobalValues.Game.SetScene(&MainMenuScene{})
			},
		),
	)
}
