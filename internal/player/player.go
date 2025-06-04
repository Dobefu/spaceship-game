package player

import (
	"github.com/Dobefu/spaceship-game/internal/bullet"
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	game_object.GameObject

	bulletPool    *[]*bullet.Bullet
	shootCooldown float64

	fireScale float64

	velocity vectors.Vector3
	scale    float64
	rotation float64

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewPlayer(
	position vectors.Vector3,
	bulletPool *[]*bullet.Bullet,
) (player *Player) {
	player = &Player{
		bulletPool: bulletPool,

		velocity: vectors.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		},
		scale:    1,
		rotation: 0,
	}

	player.SetPosition(position)
	player.SetIsActive(true)

	return player
}
