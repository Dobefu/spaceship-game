package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

func (p *Player) Update() (err error) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= .1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += .1
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		sin := float32(math.Sin(p.rotation - math.Pi/2))
		cos := float32(math.Cos(p.rotation - math.Pi/2))

		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.velocity.Add(vectors.Vector2{
				X: cos * .1,
				Y: sin * .1,
			})
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.velocity.Sub(vectors.Vector2{
				X: cos * .1,
				Y: sin * .1,
			})
		}
	}

	p.position.Add(p.velocity)
	p.velocity.Mul(vectors.Vector2{
		X: .99,
		Y: .99,
	})

	return nil
}
