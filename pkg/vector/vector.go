package vector

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Vector struct {
	x, y float32
}

func Zero() Vector {
	return Vector{
		x: 0,
		y: 0,
	}
}

func New(x, y float32) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

func FromRL(v rl.Vector2) Vector {
	return Vector{
		x: v.X,
		y: v.Y,
	}
}

func Mouse() Vector {
	return New(float32(rl.GetMouseX()), float32(rl.GetMouseY()))
}

func MouseWorld(camera *rl.Camera2D) Vector {
	return FromRL(rl.GetScreenToWorld2D(rl.GetMousePosition(), *camera))
}

func (v Vector) X() float32 {
	return v.x
}

func (v Vector) Y() float32 {
	return v.y
}

func (v Vector) SetX(x float32) Vector {
	v.x = x

	return v
}

func (v Vector) SetY(y float32) Vector {
	v.y = y
	return v
}

func (v Vector) Dot(other Vector) float32 {
	return v.x*other.x + v.y*other.y
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v Vector) Scale(s float32) Vector {
	return Vector{
		x: v.x * s,
		y: v.y * s,
	}
}

func (v Vector) Norm() Vector {
	len := v.Length()
	if len != 0 {
		return v.Scale(1 / len)
	}

	return v
}

func (v Vector) Length() float32 {
	return float32(math.Sqrt(float64(v.x*v.x + v.y*v.y)))
}

func (v Vector) Project(other Vector) Vector {
	return other.Scale(v.Dot(other) / other.Dot(other))
}

func (v Vector) AngleTo(other Vector) float32 {
	return float32(math.Atan2(float64(other.y-v.y), float64(other.x-v.x)))
}

func (v Vector) RL() rl.Vector2 {
	return rl.Vector2{X: v.x, Y: v.y}
}

func (v Vector) Draw() {
	rl.DrawLine(0, 0, int32(v.x), int32(v.y), rl.Red)
	rl.DrawCircle(int32(v.x), int32(v.y), 5, rl.Blue)
}

func (v Vector) Lerp(other Vector, t float32) Vector {
	return v.Add(other.Sub(v).Scale(min(t, 1)))
}
