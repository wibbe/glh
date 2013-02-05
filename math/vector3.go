
package math

import (
    "math"
    "fmt"
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

func (self Vector3) Dot(vec Vector3) float32 {
    return (self.X * vec.X) + (self.Y * vec.Y) + (self.Z * vec.Z)
}

func (v Vector3) Length() float32 {
    return float32(math.Sqrt(float64((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))))
}

func (v Vector3) String() string {
    return fmt.Sprintf("[%f, %f, %f]", v.X, v.Y, v.Z)
}

func (self Vector3) Normalized() Vector3 {
    d := 1.0 / self.Length()
    return Vector3{self.X * d, self.Y * d, self.Z * d}
}

