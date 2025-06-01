package game_scene

import (
	"math/rand"

	"github.com/Dobefu/spaceship-game/internal/particles/star"
	"github.com/Dobefu/spaceship-game/internal/player"
	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type GameScene struct {
	scene.Scene
}

func (s *GameScene) Init() {
	player := player.NewPlayer(vectors.Vector2{X: 0, Y: 0})

	s.AddGameObject(player)
	s.Camera.SetTarget(player)

	for range 200 {
		randX := rand.Float32() * 640
		randY := rand.Float32() * 640
		randZ := rand.Float32() * 3

		s.AddGameObject(star.NewStar(vectors.Vector3{
			X: randX,
			Y: randY,
			Z: randZ,
		}))
	}
}
