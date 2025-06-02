package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (i *Input) updateLeftStick() {
	// Reset the axes.
	i.leftStick.Horizontal = 0
	i.leftStick.Vertical = 0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		i.leftStick.Horizontal -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		i.leftStick.Horizontal += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		i.leftStick.Vertical -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		i.leftStick.Vertical += 1
	}

	math.Max(i.leftStick.Horizontal, math.Max(0, 1))
	math.Max(i.leftStick.Vertical, math.Max(0, 1))

	// TODO: Handle joysticks and deadzones.
}
