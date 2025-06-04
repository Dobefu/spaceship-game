package bullet

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (b *Bullet) Fire(
	from vectors.Vector3,
	angle float64,
	velocity vectors.Vector3,
) {
	b.SetIsActive(true)
	b.SetPosition(from)

	velocity.Add(vectors.Vector3{
		X: -math.Cos(angle+math.Pi/2) * 15,
		Y: -math.Sin(angle+math.Pi/2) * 15,
		Z: 0,
	})

	b.velocity = velocity
}
