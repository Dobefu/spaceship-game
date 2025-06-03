package bullet

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/interfaces"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type Bullet struct {
	interfaces.Bullet
	game_object.GameObject

	size     float64
	velocity vectors.Vector3
}

func NewBullet() (bullet *Bullet) {
	bullet = &Bullet{
		size:     5,
		velocity: vectors.Vector3{},
	}

	bullet.SetIsActive(false)

	return bullet
}
