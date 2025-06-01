package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (p *Player) Draw(screen *ebiten.Image) {
	modelPath = vector.Path{}

	for _, points := range modelPathPoints {
		sin := float32(math.Sin(p.rotation))
		cos := float32(math.Cos(p.rotation))

		modelPath.LineTo(
			(((points[0]-50)*cos-(points[1]-50)*sin)+p.x)*p.scale,
			(((points[0]-50)*sin+(points[1]-50)*cos)+p.y)*p.scale,
		)
	}

	modelPath.Close()

	p.vertices, p.indices = modelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, whiteSubImage, drawTrianglesOptions)
}
