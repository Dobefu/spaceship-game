package player

import (
	"image"
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/bullet"
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/options"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	modelPath       vector.Path
	modelPathPoints []vectors.Vector2
	modelCenter     vectors.Vector2

	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	strokeOptions        *vector.StrokeOptions
	drawTrianglesOptions *ebiten.DrawTrianglesOptions
)

func init() {
	modelPathPoints = []vectors.Vector2{
		{X: 25, Y: 0},
		{X: 50, Y: 50},
		{X: 25, Y: 37},
		{X: 0, Y: 50},
	}

	for _, points := range modelPathPoints {
		modelCenter.X += points.X
		modelCenter.Y += points.Y
	}

	modelCenter.X /= float64(len(modelPathPoints))
	modelCenter.Y /= float64(len(modelPathPoints))

	whiteImage.Fill(color.White)
	strokeOptions = &vector.StrokeOptions{
		Width: options.GlobalOptions.OutlineWidth,
	}

	drawTrianglesOptions = &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	}
}

type Player struct {
	game_object.GameObject

	bulletPool *[]bullet.Bullet

	velocity vectors.Vector2
	scale    float64
	rotation float64

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewPlayer(
	position vectors.Vector2,
	bulletPool *[]bullet.Bullet,
) (player *Player) {
	player = &Player{
		bulletPool: bulletPool,

		velocity: vectors.Vector2{
			X: 0,
			Y: 0,
		},
		scale:    1,
		rotation: 0,
	}

	player.SetPosition(position)
	player.SetIsActive(true)

	return player
}
