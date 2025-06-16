package asteroid

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	game_object.GameObject

	rotation float64
	scale    float64

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewAsteroid(position vectors.Vector3) (asteroid *Asteroid) {
	asteroid = &Asteroid{
		scale: 1,
	}

	asteroid.SetPosition(position)
	asteroid.SetIsActive(true)

	return asteroid
}
