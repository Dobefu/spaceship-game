package player

import (
	"github.com/Dobefu/spaceship-game/internal/input"
)

func (p *Player) Update() (err error) {
	p.HandleMovement()

	if input.GlobalInput.GetButtonA().IsPressed {
		p.Shoot()
	}

	return nil
}
