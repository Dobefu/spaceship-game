package player

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	game_object.IGameObject

	X float32
	Y float32
}

func (p Player) Update() (err error) {
	return nil
}

func (p Player) Draw(screen *ebiten.Image) {
}
