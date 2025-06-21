package asteroid

import "github.com/Dobefu/spaceship-game/internal/globals"

func (a *Asteroid) Update() (err error) {
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	screenPos := camera.WorldToScreenPosition2D(a.GetPosition())

	// If the asteroid is off-screen, deactivate it.
	if camera.IsPositionWithinBounds(screenPos, 10*a.scale) {
		a.SetIsActive(false)
		return nil
	}

	return nil
}
