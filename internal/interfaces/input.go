package interfaces

type Input interface {
	Update()
	GetStickLeft() Axis
	GetButtonA() Button

	updateStickLeft()
	updateButtonA()
}

type Axis struct {
	Horizontal float64
	Vertical   float64
}

type Button struct {
	IsPressed bool
	IsHeld    bool
}
