package math

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float32
}

func (v Vector3) Clone() Vector3 {
	return Vector3{v.X, v.Y, v.Z}
}

func (self Vector3) Add(vec Vector3) Vector3 {
	return Vector3{self.X + vec.X, self.Y + vec.Y, self.Z + vec.Z}
}

func (self Vector3) Sub(vec Vector3) Vector3 {
	return Vector3{self.X - vec.X, self.Y - vec.Y, self.Z - vec.Z}
}

func (self Vector3) Mul(scalar float32) Vector3 {
	return Vector3{self.X * scalar, self.Y * scalar, self.Z * scalar}
}

func (v Vector3) Length() float32 {
	return float32(math.Sqrt(float64((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))))
}

func (v1 Vector3) Dot(v2 Vector3) float32 {
	return (v1.X * v2.X) + (v1.Y * v2.Y) + (v1.Z * v2.Z)
}

func (a Vector3) Cross(b Vector3) Vector3 {
	return Vector3{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

func (v Vector3) String() string {
	return fmt.Sprintf("[%f, %f, %f]", v.X, v.Y, v.Z)
}
func (self Vector3) Normalized() Vector3 {
	d := 1.0 / self.Length()
	return Vector3{self.X * d, self.Y * d, self.Z * d}
}
