package interfaces

import (
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type Bullet interface {
	Fire(from vectors.Vector2, angle float64)
}
