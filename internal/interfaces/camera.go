package interfaces

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type Camera interface {
	SetTarget(target GameObject)
	GetPosition() (position vectors.Vector2)
	Update()
	WorldToScreenPosition2D(worldPos vectors.Vector3) vectors.Vector2
	IsPositionWithinBounds(pos vectors.Vector2, margin float64) bool
}
