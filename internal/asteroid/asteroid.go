package asteroid

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/fastrand"
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Asteroid struct {
	game_object.GameObject

	rotation float64
	scale    float64

	asteroidModelPath       vector.Path
	asteroidModelPathPoints []vectors.Vector2
	asteroidModelCenter     vectors.Vector2

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewAsteroid() (asteroid *Asteroid) {
	asteroid = &Asteroid{
		scale: 1,
	}

	asteroid.asteroidModelPathPoints = []vectors.Vector2{}

	const pointCount = 24
	for i := range pointCount {
		randX := float64(fastrand.Rand.Next()>>28) / 4
		randY := float64(fastrand.Rand.Next()>>28) / 4

		asteroid.asteroidModelPathPoints = append(asteroid.asteroidModelPathPoints, vectors.Vector2{
			X: math.Cos((float64(i)/pointCount)*(math.Pi*2)) * (randX + 20),
			Y: math.Sin((float64(i)/pointCount)*(math.Pi*2)) * (randY + 20),
		})
	}

	for _, points := range asteroid.asteroidModelPathPoints {
		asteroid.asteroidModelCenter.X += points.X
		asteroid.asteroidModelCenter.Y += points.Y
	}

	asteroid.asteroidModelCenter.X /= float64(len(asteroid.asteroidModelPathPoints))
	asteroid.asteroidModelCenter.Y /= float64(len(asteroid.asteroidModelPathPoints))

	return asteroid
}
