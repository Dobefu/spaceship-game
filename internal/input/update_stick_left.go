package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (i *Input) updateStickLeft() {
	// Reset the axes.
	i.stickLeft.Horizontal = 0
	i.stickLeft.Vertical = 0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		i.stickLeft.Horizontal -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		i.stickLeft.Horizontal += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		i.stickLeft.Vertical -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		i.stickLeft.Vertical += 1
	}

	math.Max(i.stickLeft.Horizontal, math.Max(0, 1))
	math.Max(i.stickLeft.Vertical, math.Max(0, 1))

	// TODO: Handle joysticks and deadzones.
}
