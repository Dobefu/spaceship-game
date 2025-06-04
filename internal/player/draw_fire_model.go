package player

import (
	"math"
	"math/rand"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/input"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	fireModelPath       vector.Path
	fireModelPathPoints []vectors.Vector2
)

func init() {
	fireModelPathPoints = []vectors.Vector2{
		{X: 25, Y: 6},
		{X: 33, Y: 10},
		{X: 36, Y: 16},
		{X: 36, Y: 22},
		{X: 30, Y: 30},
		{X: 25, Y: 41},
		{X: 20, Y: 30},
		{X: 14, Y: 22},
		{X: 14, Y: 16},
		{X: 17, Y: 10},
	}
}

func (p *Player) drawFireModel(screen *ebiten.Image) {
	if input.GlobalInput.GetStickLeft().Vertical < 0 {
		p.fireScale = math.Min(p.fireScale-input.GlobalInput.GetStickLeft().Vertical*0.1, 1)
	} else {
		p.fireScale = math.Max(p.fireScale-0.1, 0.0)
	}

	if p.fireScale <= 0.0 {
		return
	}

	fireModelPath = vector.Path{}
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	position := p.GetPosition()

	for _, points := range fireModelPathPoints {
		sin := math.Sin(p.rotation)
		cos := math.Cos(p.rotation)

		x := ((points.X - playerModelCenter.X) * p.fireScale) * (rand.Float64()*.1 + 1)
		y := ((points.Y-8)*p.fireScale + 8) * (rand.Float64()*.2 + 1)

		screenPos := camera.WorldToScreenPosition(vectors.Vector2{
			X: ((x*cos - y*sin) + position.X) * p.scale,
			Y: ((x*sin + y*cos) + position.Y) * p.scale,
		})

		fireModelPath.LineTo(
			float32(screenPos.X),
			float32(screenPos.Y),
		)
	}

	fireModelPath.Close()

	p.vertices, p.indices = fireModelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, whiteSubImage, drawTrianglesOptions)
}
