package bullet

import (
	"github.com/Dobefu/spaceship-game/internal/globals"
)

func (b *Bullet) Update() (err error) {
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	screenPos := camera.WorldToScreenPosition2D(b.GetPosition())

	// If the bullet is off-screen, deactivate it.
	if camera.IsPositionWithinBounds(screenPos, b.size/2) {
		b.SetIsActive(false)
		return nil
	}

	position := b.GetPosition()
	position.Add(b.velocity)
	b.SetPosition(position)

	return nil
}
