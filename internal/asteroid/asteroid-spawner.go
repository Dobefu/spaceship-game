package asteroid

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

type AsteroidSpawner struct {
	game_object.GameObject

	asteroidPool *[]*Asteroid
}

func NewAsteroidSpawner() (asteroidSpawner *AsteroidSpawner) {
	s := globals.GlobalValues.Game.GetScene()

	var asteroidPool []*Asteroid

	for range 100 {
		a := NewAsteroid()
		s.AddGameObject(a)

		asteroidPool = append(asteroidPool, a)
	}

	asteroidSpawner = &AsteroidSpawner{
		asteroidPool: &asteroidPool,
	}

	asteroidSpawner.SetIsActive(true)

	return asteroidSpawner
}

func (a *AsteroidSpawner) Draw(screen *ebiten.Image) {
	// noop
}

func (a *AsteroidSpawner) Update() (err error) {
	return nil
}
