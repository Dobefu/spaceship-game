package game_object

import "github.com/hajimehoshi/ebiten"

type gameObject interface {
	Update() (err error)
	Draw(screen *ebiten.Image)
}

type GameObject struct {
	gameObject

	X float32
	Y float32
}
