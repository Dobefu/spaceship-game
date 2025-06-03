package input

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (i *Input) GetStickLeft() interfaces.Axis {
	return i.stickLeft
}
