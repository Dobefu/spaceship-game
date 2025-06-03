package camera

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type Camera struct {
	interfaces.Camera

	position vectors.Vector2
	target   interfaces.GameObject
}
