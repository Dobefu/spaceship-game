package game_object

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type IGameObject interface {
	Update(offset vectors.Vector2) (err error)
	Draw(screen *ebiten.Image, offset vectors.Vector2)
	GetPosition() (position *vectors.Vector2)
	SetIsActive(isActive bool)
	GetIsActive() (isActive bool)
}

type GameObject struct {
	IGameObject

	position vectors.Vector2
	isActive bool
}
