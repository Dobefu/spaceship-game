package scene

import (
	"math/rand"

	"github.com/Dobefu/spaceship-game/internal/asteroid"
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/particles/star"
	"github.com/Dobefu/spaceship-game/internal/player"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type GameScene struct {
	Scene
}

func (s *GameScene) Init() {
	s.SetCanPause(true)

	player := player.NewPlayer(vectors.Vector3{X: 0, Y: 0, Z: 1000})

	s.AddGameObject(player)
	s.Camera.SetTarget(player)

	for range 100 {
		randX := rand.Float64() * float64(globals.GlobalValues.Width) * 3
		randY := rand.Float64() * float64(globals.GlobalValues.Height) * 3
		randZ := rand.Float64() * -3

		s.AddGameObject(star.NewStar(vectors.Vector3{
			X: randX,
			Y: randY,
			Z: randZ,
		}))
	}

	s.AddGameObject(asteroid.NewAsteroidSpawner())
}
