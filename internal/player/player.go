package player

import (
	"image"
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	modelPath       vector.Path
	modelPathPoints [][2]float32

	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	strokeOptions        *vector.StrokeOptions
	drawTrianglesOptions *ebiten.DrawTrianglesOptions
)

func init() {
	modelPathPoints = [][2]float32{
		{50, 0},
		{100, 100},
		{50, 75},
		{0, 100},
	}

	whiteImage.Fill(color.White)
	strokeOptions = &vector.StrokeOptions{
		Width: 2,
	}

	drawTrianglesOptions = &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	}
}

type Player struct {
	game_object.IGameObject

	X float32
	Y float32

	vertices []ebiten.Vertex
	indices  []uint16
}

func (p Player) Update() (err error) {
	return nil
}

func (p Player) Draw(screen *ebiten.Image) {
	modelPath = vector.Path{}

	for _, points := range modelPathPoints {
		modelPath.LineTo(points[0]+p.X, points[1]+p.Y)
	}

	modelPath.Close()

	p.vertices, p.indices = modelPath.AppendVerticesAndIndicesForStroke(p.vertices[:0], p.indices[:0], strokeOptions)
	screen.DrawTriangles(p.vertices, p.indices, whiteSubImage, drawTrianglesOptions)
}
