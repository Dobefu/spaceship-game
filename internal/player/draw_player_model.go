package player

import (
	"math"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	playerModelPath       vector.Path
	playerModelPathPoints []vectors.Vector2
	playerModelCenter     vectors.Vector2
)

func init() {
	playerModelPathPoints = []vectors.Vector2{
		{X: 25, Y: 0},
		{X: 50, Y: 50},
		{X: 25, Y: 37},
		{X: 0, Y: 50},
	}

	for _, points := range playerModelPathPoints {
		playerModelCenter.X += points.X
		playerModelCenter.Y += points.Y
	}

	playerModelCenter.X /= float64(len(playerModelPathPoints))
	playerModelCenter.Y /= float64(len(playerModelPathPoints))
}

func (p *Player) drawPlayerModel(screen *ebiten.Image) {
	playerModelPath = vector.Path{}
	camera := globals.GlobalValues.Game.GetScene().GetCamera()
	position := p.GetPosition()

	for _, points := range playerModelPathPoints {
		sin := math.Sin(p.rotation)
		cos := math.Cos(p.rotation)

		x := (points.X - playerModelCenter.X)
		y := (points.Y - playerModelCenter.Y)

		screenPos := camera.WorldToScreenPosition(vectors.Vector2{
			X: ((x*cos - y*sin) + position.X) * p.scale,
			Y: ((x*sin + y*cos) + position.Y) * p.scale,
		})

		playerModelPath.LineTo(
			float32(screenPos.X),
			float32(screenPos.Y),
		)
	}

	playerModelPath.Close()

	p.vertices, p.indices = playerModelPath.AppendVerticesAndIndicesForFilling(p.vertices[:0], p.indices[:0])
	screen.DrawTriangles(p.vertices, p.indices, blackSubImage, drawTrianglesOptions)
	p.vertices, p.indices = playerModelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, whiteSubImage, drawTrianglesOptions)
}
