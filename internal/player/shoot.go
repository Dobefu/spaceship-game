package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (p *Player) Shoot() {
	for _, b := range *p.bulletPool {
		// Skip already active bullets.
		if b.GetIsActive() {
			continue
		}

		from := vectors.Vector3{
			X: math.Cos(p.rotation - math.Pi/2),
			Y: math.Sin(p.rotation - math.Pi/2),
			Z: 0,
		}

		from.Normalize()
		from.Mul(vectors.Vector3{X: 50, Y: 50, Z: 0})
		from.Add(p.GetPosition())

		b.Fire(from, p.rotation, p.velocity)
		break
	}
}
