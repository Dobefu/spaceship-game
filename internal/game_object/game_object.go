package game_object

import "github.com/hajimehoshi/ebiten/v2"

type IGameObject interface {
	Update() (err error)
	Draw(screen *ebiten.Image)
}

type GameObject struct {
	IGameObject

	X float32
	Y float32
}
