package interfaces

type Input interface {
	Update()
	GetStickLeft() Axis
	GetButtonStart() Button
	GetButtonA() Button

	updateStickLeft()
	updateButtonStart()
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
