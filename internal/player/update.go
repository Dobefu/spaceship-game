package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (p *Player) Update() (err error) {
	leftAxis := input.GlobalInput.GetLeftStick()
	p.rotation += leftAxis.Horizontal * .1

	if leftAxis.Vertical != 0 {
		sin := math.Sin(p.rotation + math.Pi/2)
		cos := math.Cos(p.rotation + math.Pi/2)

		p.velocity.Add(vectors.Vector2{
			X: float32(cos * leftAxis.Vertical * .1),
			Y: float32(sin * leftAxis.Vertical * .1),
		})
	}

	position := p.GetPosition()
	position.Add(p.velocity)
	p.SetPosition(position)

	p.velocity.Mul(vectors.Vector2{
		X: .99,
		Y: .99,
	})

	return nil
}
