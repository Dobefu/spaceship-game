package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (p *Player) Draw(screen *ebiten.Image) {
	modelPath = vector.Path{}
	camera := globals.GlobalValues.Game.GetScene().GetCamera()

	for _, points := range modelPathPoints {
		sin := math.Sin(p.rotation)
		cos := math.Cos(p.rotation)

		x := (points.X - modelCenter.X)
		y := (points.Y - modelCenter.Y)

		screenPos := camera.WorldToScreenPosition(vectors.Vector2{
			X: ((x*cos - y*sin) + p.GetPosition().X) * p.scale,
			Y: ((x*sin + y*cos) + p.GetPosition().Y) * p.scale,
		})

		modelPath.LineTo(
			float32(screenPos.X),
			float32(screenPos.Y),
		)
	}

	modelPath.Close()

	p.vertices, p.indices = modelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, whiteSubImage, drawTrianglesOptions)
}
