package star

import (
	"image/color"
	"math"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/options"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	game_object.GameObject

	position vectors.Vector3
}

func NewStar(position vectors.Vector3) (star *Star) {
	star = &Star{
		position: position,
	}

	star.SetIsActive(true)

	return star
}

func (s *Star) Update() (err error) {
	return nil
}

func (s *Star) Draw(screen *ebiten.Image, offset vectors.Vector2) {
	gameHeight := float64(options.GlobalOptions.Height)
	gameWidth := float64(options.GlobalOptions.Width)

	vector.DrawFilledRect(
		screen,
		float32(math.Mod(math.Mod(s.position.X-offset.X, gameWidth)+gameWidth, gameWidth)*s.position.Z),
		float32(math.Mod(math.Mod(s.position.Y-offset.Y, gameHeight)+gameHeight, gameHeight)*s.position.Z),
		float32(s.position.Z),
		float32(s.position.Z),
		color.White,
		false,
	)
}
