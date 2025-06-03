package game_scene

import (
	"math/rand"

	"github.com/Dobefu/spaceship-game/internal/bullet"
	"github.com/Dobefu/spaceship-game/internal/particles/star"
	"github.com/Dobefu/spaceship-game/internal/player"
	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type GameScene struct {
	scene.Scene
}

func (s *GameScene) Init() {
	var bulletPool []*bullet.Bullet

	for range 10 {
		b := bullet.NewBullet()
		s.AddGameObject(b)

		bulletPool = append(bulletPool, b)
	}

	player := player.NewPlayer(vectors.Vector3{X: 0, Y: 0, Z: 1000}, &bulletPool)

	s.AddGameObject(player)
	s.Camera.SetTarget(player)

	for range 200 {
		randX := rand.Float64() * 640
		randY := rand.Float64() * 640
		randZ := rand.Float64() * 3

		s.AddGameObject(star.NewStar(vectors.Vector3{
			X: randX,
			Y: randY,
			Z: randZ,
		}))
	}
}
