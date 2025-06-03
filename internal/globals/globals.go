package globals

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
)

var (
	GlobalValues = Globals{
		Game:   nil,
		Height: 640,
		Width:  640,

		OutlineWidth: 1.5,
	}
)

type Globals struct {
	Game   interfaces.Game
	Height int
	Width  int

	OutlineWidth float32
}
