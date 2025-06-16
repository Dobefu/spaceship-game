package asteroid

import (
	"image/color"
	"math"

	"github.com/Dobefu/spaceship-game/internal/fastrand"
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	asteroidModelPath       vector.Path
	asteroidModelPathPoints []vectors.Vector2
	asteroidModelCenter     vectors.Vector2

	whiteImage = ebiten.NewImage(1, 1)

	strokeOptions        *vector.StrokeOptions
	drawTrianglesOptions *ebiten.DrawTrianglesOptions
)

func init() {
	asteroidModelPathPoints = []vectors.Vector2{}

	const pointCount = 24
	for i := range pointCount {
		randX := float64(fastrand.Rand.Next()>>28) / 4
		randY := float64(fastrand.Rand.Next()>>28) / 4

		asteroidModelPathPoints = append(asteroidModelPathPoints, vectors.Vector2{
			X: math.Cos((float64(i)/pointCount)*(math.Pi*2)) * (randX + 20),
			Y: math.Sin((float64(i)/pointCount)*(math.Pi*2)) * (randY + 20),
		})
	}

	for _, points := range asteroidModelPathPoints {
		asteroidModelCenter.X += points.X
		asteroidModelCenter.Y += points.Y
	}

	asteroidModelCenter.X /= float64(len(asteroidModelPathPoints))
	asteroidModelCenter.Y /= float64(len(asteroidModelPathPoints))
}

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
	asteroidModelPath = vector.Path{}
	position := a.GetPosition()
	camera := globals.GlobalValues.Game.GetScene().GetCamera()

	for _, points := range asteroidModelPathPoints {
		sin := math.Sin(a.rotation)
		cos := math.Cos(a.rotation)

		x := (points.X - asteroidModelCenter.X)
		y := (points.Y - asteroidModelCenter.Y)

		screenPos := camera.WorldToScreenPosition2D(vectors.Vector3{
			X: ((x*cos - y*sin) + position.X) * a.scale,
			Y: ((x*sin + y*cos) + position.Y) * a.scale,
			Z: position.Z,
		})

		asteroidModelPath.LineTo(
			float32(screenPos.X),
			float32(screenPos.Y),
		)
	}

	asteroidModelPath.Close()

	a.vertices, a.indices = asteroidModelPath.AppendVerticesAndIndicesForStroke(a.vertices[:0], a.indices[:0], strokeOptions)
	screen.DrawTriangles(a.vertices, a.indices, whiteImage, drawTrianglesOptions)
}
