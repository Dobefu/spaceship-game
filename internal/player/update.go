package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/fastrand"
	"github.com/Dobefu/spaceship-game/internal/input"
)

func (p *Player) Update() (err error) {
	p.HandleMovement()

	for i := range fireModelPathPoints {
		randX := float64(fastrand.Rand.Next()>>24)/1270 + 1
		randY := float64(fastrand.Rand.Next()>>24)/1000 + 1

		fireModelPathOffsets[i].X = randX
		fireModelPathOffsets[i].Y = randY
	}

	p.shootCooldown = math.Max(p.shootCooldown-1, 0)

	if input.GlobalInput.GetButtonA().IsHeld && p.shootCooldown <= 0 {
		p.Shoot()
		p.shootCooldown = 10
	}

	return nil
}
