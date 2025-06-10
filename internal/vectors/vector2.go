package vectors

import (
	"math"
)

type IVector2 interface {
	Add(vec Vector2)
	Sub(vec Vector2)
	Mul(vec Vector2)
	Div(vec Vector2)
	Normalize() (vec Vector2)
}

type Vector2 struct {
	IVector2

	X float64
	Y float64
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

func (v *Vector2) Normalize() (vec Vector2) {
	vec = *v
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)

	if magnitude != 0 {
		vec.X /= magnitude
		vec.Y /= magnitude
	}

	return vec
}
