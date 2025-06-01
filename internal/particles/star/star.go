package star

import (
	"image/color"

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

func (s *Star) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		s.position.X,
		s.position.Y,
		s.position.Z,
		s.position.Z,
		color.White,
		false,
	)
}
