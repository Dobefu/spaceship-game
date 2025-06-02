package input

var (
	GlobalInput Input
	deadzone    = .1
)

type IInput interface {
	Update()
	GetStickLeft() Axis
	GetButtonA() Button

	updateStickLeft()
	updateButtonA()
}

type Input struct {
	IInput

	stickLeft Axis
	buttonA   Button
}
