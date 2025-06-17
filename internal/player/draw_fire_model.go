package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	fireModelPath        vector.Path
	fireModelPathPoints  []vectors.Vector2
	fireModelPathOffsets []vectors.Vector2
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

	// Pre-fill the offsets slice, to get allocations out of the way.
	for range fireModelPathPoints {
		fireModelPathOffsets = append(
			fireModelPathOffsets,
			vectors.Vector2{X: 0, Y: 0},
		)
	}
}

func (p *Player) drawFireModel(screen *ebiten.Image) {

	if p.fireScale <= 0.0 {
		return
	}

	fireModelPath = vector.Path{}
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	position := p.GetPosition()

	for i, points := range fireModelPathPoints {
		sin := math.Sin(p.rotation)
		cos := math.Cos(p.rotation)

		x := ((points.X - playerModelCenter.X) * p.fireScale) * fireModelPathOffsets[i].X
		y := ((points.Y-8)*p.fireScale + 8) * fireModelPathOffsets[i].Y

		screenPos := camera.WorldToScreenPosition2D(vectors.Vector3{
			X: ((x*cos - y*sin) + position.X) * p.scale,
			Y: ((x*sin + y*cos) + position.Y) * p.scale,
			Z: position.Z,
		})

		fireModelPath.LineTo(
			float32(screenPos.X),
			float32(screenPos.Y),
		)
	}

	fireModelPath.Close()

	p.vertices, p.indices = fireModelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, redImage, drawTrianglesOptions)
}
