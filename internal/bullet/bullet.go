package bullet

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type Bullet struct {
	game_object.GameObject

	isActive bool
}

func NewBullet(position vectors.Vector2) (bullet *Bullet) {
	bullet = &Bullet{isActive: false}
	bullet.SetPosition(position)

	return bullet
}
