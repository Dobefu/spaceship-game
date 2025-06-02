package vectors

type IVector3 interface {
	Add(vec Vector3)
	Sub(vec Vector3)
	Mul(vec Vector3)
	Div(vec Vector3)
}

type Vector3 struct {
	IVector3

	X float64
	Y float64
	Z float64
}

func (v *Vector3) Add(vec Vector3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

func (v *Vector3) Sub(vec Vector3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

func (v *Vector3) Mul(vec Vector3) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
}

func (v *Vector3) Div(vec Vector3) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
}
