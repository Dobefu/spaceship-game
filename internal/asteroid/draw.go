package asteroid

import (
	"image/color"
	"math"

	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	whiteImage = ebiten.NewImage(1, 1)

	strokeOptions        *vector.StrokeOptions
	drawTrianglesOptions *ebiten.DrawTrianglesOptions
)

func init() {
	whiteImage.Set(0, 0, color.White)

	strokeOptions = &vector.StrokeOptions{
		Width: globals.GlobalValues.OutlineWidth,
	}

	drawTrianglesOptions = &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	}
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	a.asteroidModelPath = vector.Path{}
	position := a.GetPosition()
	camera := globals.GlobalValues.Game.GetScene().GetCamera()

	for _, points := range a.asteroidModelPathPoints {
		sin := math.Sin(a.rotation)
		cos := math.Cos(a.rotation)

		x := (points.X - a.asteroidModelCenter.X)
		y := (points.Y - a.asteroidModelCenter.Y)

		screenPos := camera.WorldToScreenPosition2D(vectors.Vector3{
			X: ((x*cos - y*sin) + position.X) * a.scale,
			Y: ((x*sin + y*cos) + position.Y) * a.scale,
			Z: position.Z,
		})

		a.asteroidModelPath.LineTo(
			float32(screenPos.X),
			float32(screenPos.Y),
		)
	}

	a.asteroidModelPath.Close()

	a.vertices, a.indices = a.asteroidModelPath.AppendVerticesAndIndicesForStroke(a.vertices[:0], a.indices[:0], strokeOptions)
	screen.DrawTriangles(a.vertices, a.indices, whiteImage, drawTrianglesOptions)
}
