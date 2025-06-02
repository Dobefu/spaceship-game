package input

var (
	GlobalInput Input
)

type IInput interface {
	Update()
}

type Input struct {
	IInput

	LeftStick Axis
}

func (i *Input) Update() {
}
