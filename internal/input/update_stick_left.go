package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (i *Input) updateStickLeft() {
	// Reset the axes.
	i.stickLeft.Horizontal = 0
	i.stickLeft.Vertical = 0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyA) ||
		ebiten.IsKeyPressed(ebiten.KeyH) ||
		ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton14) {
		i.stickLeft.Horizontal -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) ||
		ebiten.IsKeyPressed(ebiten.KeyD) ||
		ebiten.IsKeyPressed(ebiten.KeyL) ||
		ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton15) {
		i.stickLeft.Horizontal += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) ||
		ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyK) {
		i.stickLeft.Vertical -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) ||
		ebiten.IsKeyPressed(ebiten.KeyJ) ||
		ebiten.IsKeyPressed(ebiten.KeyS) {
		i.stickLeft.Vertical += 1
	}

	i.stickLeft.Horizontal += ebiten.GamepadAxisValue(0, 0)

	if ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton6) {
		i.stickLeft.Vertical = 1
	}

	if ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton7) {
		i.stickLeft.Vertical = -1

		if ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton6) {
			i.stickLeft.Vertical = 0
		}
	}

	math.Max(i.stickLeft.Horizontal, math.Max(0, 1))
	math.Max(i.stickLeft.Vertical, math.Max(0, 1))

	if math.Abs(i.stickLeft.Horizontal) < deadzone {
		i.stickLeft.Horizontal = 0
	}

	if math.Abs(i.stickLeft.Vertical) < deadzone {
		i.stickLeft.Vertical = 0
	}
}
