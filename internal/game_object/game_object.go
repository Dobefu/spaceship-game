package game_object

type gameObject interface{}

type GameObject struct {
	gameObject

	X float32
	Y float32
}
