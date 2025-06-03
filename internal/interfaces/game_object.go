package interfaces

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	Update() (err error)
	Draw(screen *ebiten.Image)
	GetPosition() (position *vectors.Vector2)
	SetIsActive(isActive bool)
	GetIsActive() (isActive bool)
}
