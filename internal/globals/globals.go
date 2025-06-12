package globals

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	GlobalValues Globals
)

type Globals struct {
	Game   interfaces.Game
	Height int
	Width  int
	Scale  int

	OutlineWidth float32
}

func init() {
	scale := int(ebiten.Monitor().DeviceScaleFactor())

	GlobalValues = Globals{
		Game:   nil,
		Height: 640,
		Width:  640,
		Scale:  scale,

		OutlineWidth: 1.5,
	}
}
