
package glh

import "math"

type Vector3 struct {
    X, Y, Z float32
}

func (v *Vector3) Length() float32 {
    return float32(math.Sqrt(float64((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))))
}

