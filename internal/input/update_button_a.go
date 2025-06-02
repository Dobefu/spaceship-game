package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (i *Input) updateButtonA() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton0) {
		if !i.buttonA.IsPressed {
			i.buttonA.IsPressed = true
		}

		if i.buttonA.IsHeld {
			i.buttonA.IsPressed = false
		}

		i.buttonA.IsHeld = true
		return
	}

	i.buttonA.IsHeld = false
}
