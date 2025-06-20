package scene

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/ui"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (s *Scene) SetCanPause(canPause bool) {
	s.canPause = canPause

	// Create the pause UI gameObjects if they're not already present.
	if len(s.GetPauseScreenGameObjects()) <= 0 {
		s.AddPauseScreenGameObject(
			ui.NewText(
				vectors.Vector2{
					X: float64(globals.GlobalValues.Width) / 2,
					Y: float64(globals.GlobalValues.Height)/2 - 75,
				},
				text.AlignCenter,
				text.AlignCenter,
				"Paused",
				32,
			),
		)

		s.AddPauseScreenGameObject(
			ui.NewButton(
				vectors.Vector2{
					X: float64(globals.GlobalValues.Width) / 2,
					Y: float64(globals.GlobalValues.Height) / 2,
				},
				175,
				50,
				text.AlignCenter,
				text.AlignStart,
				"Continue",
				func(b *ui.Button) {
					s.SetIsPaused(false)
				},
			),
		)

		s.AddPauseScreenGameObject(
			ui.NewButton(
				vectors.Vector2{
					X: float64(globals.GlobalValues.Width) / 2,
					Y: float64(globals.GlobalValues.Height)/2 + 75,
				},
				175,
				50,
				text.AlignCenter,
				text.AlignStart,
				"Main menu",
				func(b *ui.Button) {
					globals.GlobalValues.Game.SetScene(&MainMenuScene{})
				},
			),
		)
	}
}
