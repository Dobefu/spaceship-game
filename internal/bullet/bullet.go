package bullet

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type IBullet interface {
	Fire(from vectors.Vector2, angle float64)
}

type Bullet struct {
	IBullet
	game_object.GameObject

	velocity vectors.Vector2
}

func NewBullet() (bullet *Bullet) {
	bullet = &Bullet{
		velocity: vectors.Vector2{},
	}

	bullet.SetIsActive(false)

	return bullet
}
