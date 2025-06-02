package input

var (
	GlobalInput Input
	deadzone    = .05
)

type IInput interface {
	Update()
	GetLeftStick() Axis

	updateLeftStick()
}

type Input struct {
	IInput

	leftStick Axis
}
