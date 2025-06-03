package player

import (
	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (p *Player) Update(offset vectors.Vector2) (err error) {
	p.HandleMovement()

	if input.GlobalInput.GetButtonA().IsPressed {
		p.Shoot()
	}

	return nil
}
