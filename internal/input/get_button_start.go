package input

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

func (i *Input) GetButtonStart() interfaces.Button {
	return i.buttonStart
}
