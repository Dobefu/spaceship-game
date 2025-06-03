package bullet

import (
	"github.com/Dobefu/spaceship-game/internal/options"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

func (b *Bullet) Update(offset vectors.Vector2) (err error) {
	gameHeight := float64(options.GlobalOptions.Height)
	gameWidth := float64(options.GlobalOptions.Width)
	posX := b.GetPosition().X - offset.X
	posY := b.GetPosition().Y - offset.Y

	// If the bullet is off-screen, deactivate it.
	if posX < -b.size/2 || posX > gameWidth+b.size/2 || posY < -b.size/2 || posY > gameHeight+b.size/2 {
		b.SetIsActive(false)
		return nil
	}

	b.GetPosition().Add(b.velocity)

	return nil
}
