package game_object

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type IGameObject interface {
	Update() (err error)
	Draw(screen *ebiten.Image)
	GetPosition() (position vectors.Vector2)
}

type GameObject struct {
	IGameObject

	position vectors.Vector2
}
