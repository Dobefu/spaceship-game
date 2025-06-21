package player

import (
	"github.com/Dobefu/spaceship-game/internal/bullet"
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/particles/smoke"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	game_object.GameObject

	bulletPool    *[]*bullet.Bullet
	shootCooldown float64

	fireScale      float64
	smokeParticles *[]*smoke.Smoke

	velocity vectors.Vector3
	scale    float64
	rotation float64

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewPlayer(
	position vectors.Vector3,
) (player *Player) {
	s := globals.GlobalValues.Game.GetScene()

	var bulletPool []*bullet.Bullet

	for range 10 {
		b := bullet.NewBullet()
		s.AddGameObject(b)

		bulletPool = append(bulletPool, b)
	}

	var smokeParticles []*smoke.Smoke

	for range 60 {
		particle := smoke.NewSmoke()
		s.AddGameObject(particle)

		smokeParticles = append(smokeParticles, particle)
	}

	player = &Player{
		bulletPool:     &bulletPool,
		smokeParticles: &smokeParticles,

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
