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

		x := (points.X - modelCenter.X)
		y := (points.Y - modelCenter.Y)

		modelPath.LineTo(
			((x*cos-y*sin)+p.position.X)*p.scale,
			((x*sin+y*cos)+p.position.Y)*p.scale,
		)
	}

	modelPath.Close()

	p.vertices, p.indices = modelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, whiteSubImage, drawTrianglesOptions)
}
