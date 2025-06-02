package star

import (
	"image/color"
	"math"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	game_object.GameObject

	position vectors.Vector3
}

func NewStar(position vectors.Vector3) *Star {
	return &Star{
		position: position,
	}
}

func (s *Star) Update() (err error) {
	return nil
}

func (s *Star) Draw(screen *ebiten.Image, offset vectors.Vector2) {
	// TODO: Get the actual canvas size instead of hardcoding it here.
	vector.DrawFilledRect(
		screen,
		float32(math.Mod(math.Mod(s.position.X-offset.X, 640)+640, 640)*s.position.Z),
		float32(math.Mod(math.Mod(s.position.Y-offset.Y, 640)+640, 640)*s.position.Z),
		float32(s.position.Z),
		float32(s.position.Z),
		color.White,
		false,
	)
}
