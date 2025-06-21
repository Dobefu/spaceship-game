package asteroid

import (
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	game_object.GameObject

	rotation float64
	scale    float64

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewAsteroid() (asteroid *Asteroid) {
	asteroid = &Asteroid{
		scale: 1,
	}

	return asteroid
}
