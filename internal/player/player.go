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
	modelCenter     [2]float32

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

	for _, points := range modelPathPoints {
		modelCenter[0] += points[0]
		modelCenter[1] += points[1]
	}

	modelCenter[0] /= float32(len(modelPathPoints))
	modelCenter[1] /= float32(len(modelPathPoints))

	whiteImage.Fill(color.White)
	strokeOptions = &vector.StrokeOptions{
		Width: 2,
	}

	drawTrianglesOptions = &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	}
}

type Player struct {
	game_object.GameObject

	x        float32
	y        float32
	scale    float32
	rotation float64

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewPlayer(x float32, y float32) *Player {
	return &Player{
		x:     x,
		y:     y,
		scale: 1,
	}
}
