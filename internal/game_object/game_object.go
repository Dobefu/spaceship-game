package game_object

import (
	"github.com/Dobefu/spaceship-game/internal/interfaces"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type GameObject struct {
	interfaces.GameObject

	position vectors.Vector2
	isActive bool
}
