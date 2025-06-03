package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (p *Player) Draw(screen *ebiten.Image, offset vectors.Vector2) {
	modelPath = vector.Path{}

	for _, points := range modelPathPoints {
		sin := math.Sin(p.rotation)
		cos := math.Cos(p.rotation)

		x := (points.X - modelCenter.X)
		y := (points.Y - modelCenter.Y)

		modelPath.LineTo(
			float32((((x*cos-y*sin)+p.GetPosition().X)*p.scale)-offset.X),
			float32((((x*sin+y*cos)+p.GetPosition().Y)*p.scale)-offset.Y),
		)
	}

	modelPath.Close()

	p.vertices, p.indices = modelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, whiteSubImage, drawTrianglesOptions)
}
