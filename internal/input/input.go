package input

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

var (
	GlobalInput Input
	deadzone    = .1
)

type Input struct {
	interfaces.Input

	stickLeft   interfaces.Axis
	buttonStart interfaces.Button
	buttonA     interfaces.Button
}
