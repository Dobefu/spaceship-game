package smoke

import (
	"image/color"
	"math"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Smoke struct {
	game_object.GameObject

	scale    float32
	velocity vectors.Vector3
}

func NewSmoke() (smoke *Smoke) {
	smoke = &Smoke{}
	smoke.SetIsActive(true)

	return smoke
}

func (s *Smoke) SetScale(scale float32) {
	s.scale = scale
}

func (s *Smoke) SetVelocity(velocity vectors.Vector3) {
	s.velocity = velocity
}

func (s *Smoke) Update() (err error) {
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	screenPos := camera.WorldToScreenPosition(s.GetPosition().ToVector2())

	// If the smoke particle is off-screen, deactivate it.
	if camera.IsPositionWithinBounds(screenPos, float64(s.scale)/2) {
		s.SetIsActive(false)
		return nil
	}

	s.scale = float32(math.Max(0, float64(s.scale)-.1))

	if s.scale <= 0 {
		s.SetIsActive(false)
	}

	position := s.GetPosition()
	position.Add(s.velocity)
	s.SetPosition(position)

	return nil
}

func (s *Smoke) Draw(screen *ebiten.Image) {
	camera := globals.GlobalValues.Game.GetScene().GetCamera()

	position := s.GetPosition()

	screenPos := camera.WorldToScreenPosition(vectors.Vector2{
		X: position.X,
		Y: position.Y,
	})

	vector.StrokeCircle(
		screen,
		float32(screenPos.X),
		float32(screenPos.Y),
		s.scale,
		globals.GlobalValues.OutlineWidth,
		color.Gray{5},
		false,
	)
}
