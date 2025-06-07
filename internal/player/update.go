package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/input"
)

func (p *Player) Update() (err error) {
	p.HandleMovement()

	p.shootCooldown = math.Max(p.shootCooldown-1, 0)

	if input.GlobalInput.GetButtonA().IsHeld && p.shootCooldown <= 0 {
		p.Shoot()
		p.shootCooldown = 10
	}

	return nil
}
