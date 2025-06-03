package player

import (
	"image"
	"image/color"

	"github.com/Dobefu/spaceship-game/internal/bullet"
	"github.com/Dobefu/spaceship-game/internal/game_object"
	"github.com/Dobefu/spaceship-game/internal/globals"
	"github.com/Dobefu/spaceship-game/internal/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	modelPath       vector.Path
	modelPathPoints []vectors.Vector2
	modelCenter     vectors.Vector2

	blackImage    = ebiten.NewImage(3, 3)
	blackSubImage = blackImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
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

	blackImage.Fill(color.Black)
	whiteImage.Fill(color.White)
	strokeOptions = &vector.StrokeOptions{
		Width: globals.GlobalValues.OutlineWidth,
	}

	drawTrianglesOptions = &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	}
}

type Player struct {
	game_object.GameObject

	bulletPool    *[]*bullet.Bullet
	shootCooldown float64

	velocity vectors.Vector3
	scale    float64
	rotation float64

	vertices []ebiten.Vertex
	indices  []uint16
}

func NewPlayer(
	position vectors.Vector3,
	bulletPool *[]*bullet.Bullet,
) (player *Player) {
	player = &Player{
		bulletPool: bulletPool,

		velocity: vectors.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		},
		scale:    1,
		rotation: 0,
	}

	player.SetPosition(position)
	player.SetIsActive(true)

	return player
}
