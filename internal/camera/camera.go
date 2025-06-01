package camera

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type ICamera interface {
	SetTarget(target game_object.IGameObject)
	Update()
}

type Camera struct {
	ICamera

	position vectors.Vector2
	target   game_object.IGameObject
}
