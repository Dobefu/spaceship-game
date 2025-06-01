package vectors

type IVector2 interface {
	Add(vec Vector2)
	Sub(vec Vector2)
	Mul(vec Vector2)
	Div(vec Vector2)
}

type Vector2 struct {
	IVector2

	X float32
	Y float32
}

func (v *Vector2) Add(vec Vector2) {
	v.X += vec.X
	v.Y += vec.Y
}

func (v *Vector2) Sub(vec Vector2) {
	v.X -= vec.X
	v.Y -= vec.Y
}

func (v *Vector2) Mul(vec Vector2) {
	v.X *= vec.X
	v.Y *= vec.Y
}

func (v *Vector2) Div(vec Vector2) {
	v.X /= vec.X
	v.Y /= vec.Y
}
