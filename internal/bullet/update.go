package bullet

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
)

func (b *Bullet) Update() (err error) {
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	screenPos := camera.WorldToScreenPosition(*b.GetPosition())

	gameHeight := float64(globals.GlobalValues.Height)
	gameWidth := float64(globals.GlobalValues.Width)

	// If the bullet is off-screen, deactivate it.
	if screenPos.X < -b.size/2 || screenPos.X > gameWidth+b.size/2 || screenPos.Y < -b.size/2 || screenPos.Y > gameHeight+b.size/2 {
		b.SetIsActive(false)
		return nil
	}

	b.GetPosition().Add(b.velocity)

	return nil
}
