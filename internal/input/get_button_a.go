package input

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (i *Input) GetButtonA() interfaces.Button {
	return i.buttonA
}
