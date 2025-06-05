package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/fastrand"
	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (p *Player) HandleMovement() {
	leftAxis := input.GlobalInput.GetStickLeft()
	p.rotation += leftAxis.Horizontal * .1

	if leftAxis.Vertical != 0 {
		sin := math.Sin(p.rotation + math.Pi/2)
		cos := math.Cos(p.rotation + math.Pi/2)

		p.velocity.Add(vectors.Vector3{
			X: cos * leftAxis.Vertical * .1,
			Y: sin * leftAxis.Vertical * .1,
			Z: 0,
		})
	}

	if input.GlobalInput.GetStickLeft().Vertical < 0 {
		p.fireScale = math.Min(p.fireScale-input.GlobalInput.GetStickLeft().Vertical*0.1, 1)

		if p.fireScale >= 1 {
			for _, particle := range *p.smokeParticles {
				// Skip already active smoke particles.
				if particle.GetIsActive() {
					continue
				}

				from := vectors.Vector3{
					X: math.Cos(p.rotation - math.Pi/2),
					Y: math.Sin(p.rotation - math.Pi/2),
					Z: 0,
				}

				from.Normalize()
				from.Mul(vectors.Vector3{X: -45, Y: -45, Z: 0})
				from.Add(p.GetPosition())
				from.Add(vectors.Vector3{
					X: (float64(fastrand.Rand.Next()>>24) - 127) / 25.5,
					Y: (float64(fastrand.Rand.Next()>>24) - 127) / 25.5,
				})

				particle.SetIsActive(true)
				particle.SetScale(5 + ((float32(fastrand.Rand.Next()>>24) - 127) / 127))
				particle.SetShade(uint8(fastrand.Rand.Next() >> 28))
				particle.SetPosition(from)
				particle.SetVelocity(vectors.Vector3{
					X: (from.X - p.GetPosition().X) / 10,
					Y: (from.Y - p.GetPosition().Y) / 10,
				})

				break
			}
		}
	} else {
		p.fireScale = math.Max(p.fireScale-0.1, 0.0)
	}

	position := p.GetPosition()
	position.Add(p.velocity)
	p.SetPosition(position)

	p.velocity.Mul(vectors.Vector3{
		X: .99,
		Y: .99,
		Z: .99,
	})

}
