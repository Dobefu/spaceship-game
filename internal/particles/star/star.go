package star

import (
	"image/color"
	"math"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	game_object.GameObject
}

func NewStar(position vectors.Vector3) (star *Star) {
	star = &Star{}
	star.SetPosition(position)
	star.SetIsActive(true)

	return star
}

func (s *Star) Update() (err error) {
	return nil
}

func (s *Star) Draw(screen *ebiten.Image) {
	camera := globals.GlobalValues.Game.GetScene().GetCamera()

	gameHeight := float64(globals.GlobalValues.Height)
	gameWidth := float64(globals.GlobalValues.Width)

	position := s.GetPosition()

	screenPos := camera.WorldToScreenPosition(vectors.Vector2{
		X: position.X,
		Y: position.Y,
	})

	vector.DrawFilledRect(
		screen,
		float32(math.Mod(math.Mod(screenPos.X, gameWidth)+gameWidth, gameWidth)*position.Z),
		float32(math.Mod(math.Mod(screenPos.Y, gameHeight)+gameHeight, gameHeight)*position.Z),
		float32(position.Z),
		float32(position.Z),
		color.White,
		false,
	)
}
