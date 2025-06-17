package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (i *Input) updateButtonStart() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton9) {
		if !i.buttonStart.IsPressed {
			i.buttonStart.IsPressed = true
		}

		if i.buttonStart.IsHeld {
			i.buttonStart.IsPressed = false
		}

		i.buttonStart.IsHeld = true
		return
	}

	i.buttonStart.IsHeld = false
}
