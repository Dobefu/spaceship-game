package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (p *Player) Update(offset vectors.Vector2) (err error) {
	p.HandleMovement()

	p.shootCooldown = math.Max(p.shootCooldown-1, 0)

	if input.GlobalInput.GetButtonA().IsHeld && p.shootCooldown <= 0 {
		p.Shoot()
		p.shootCooldown = 30
	}

	return nil
}
