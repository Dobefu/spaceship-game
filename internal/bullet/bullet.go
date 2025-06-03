package bullet

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type Bullet struct {
	game_object.GameObject
}

func NewBullet(position vectors.Vector2) (bullet *Bullet) {
	bullet = &Bullet{}
	bullet.SetPosition(position)
	bullet.SetIsActive(false)

	return bullet
}
