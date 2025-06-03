package bullet

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (b *Bullet) Fire(from vectors.Vector2, angle float64) {
	b.SetIsActive(true)
	b.SetPosition(from)

	b.velocity = vectors.Vector2{
		X: -math.Cos(angle+math.Pi/2) * 10,
		Y: -math.Sin(angle+math.Pi/2) * 10,
	}
}
