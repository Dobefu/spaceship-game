package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (p *Player) HandleMovement() {
	leftAxis := input.GlobalInput.GetStickLeft()
	p.rotation += leftAxis.Horizontal * .1

	if leftAxis.Vertical != 0 {
		sin := math.Sin(p.rotation + math.Pi/2)
		cos := math.Cos(p.rotation + math.Pi/2)

		p.velocity.Add(vectors.Vector2{
			X: cos * leftAxis.Vertical * .1,
			Y: sin * leftAxis.Vertical * .1,
		})
	}

	p.GetPosition().Add(p.velocity)

	p.velocity.Mul(vectors.Vector2{
		X: .99,
		Y: .99,
	})

}
