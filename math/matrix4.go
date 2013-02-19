package math

import (
	internalMath "math"
)

type Matrix4 [16]float32

func Identity() Matrix4 {
	return [16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (self *Matrix4) SetPosition(pos Vector3) {
	self[12] = pos.X
	self[13] = pos.Y
	self[14] = pos.Z
}

func (m *Matrix4) Perspective(fov, aspect, near, far float32) {
	xymax := near * float32(internalMath.Tan(float64(fov*(internalMath.Pi/90.0))))

	ymin := -xymax
	xmin := -xymax

	width := xymax - xmin
	height := xymax - ymin

	depth := far - near
	q := -(far + near) / depth
	qn := -2.0 * (far * near) / depth

	w := 2.0 * near / width
	w = w / aspect
	h := 2.0 * near / height

	m[0] = w
	m[4] = 0
	m[8] = 0
	m[12] = 0
	m[1] = 0
	m[5] = h
	m[9] = 0
	m[13] = 0
	m[2] = 0
	m[6] = 0
	m[10] = q
	m[14] = qn
	m[3] = 0
	m[7] = 0
	m[11] = -1
	m[15] = 0
}

func (m *Matrix4) Orthographic(left, right, top, bottom, near, far float32) {
	w := right - left
	h := top - bottom
	p := far - near
	x := (right + left) / w
	y := (top + bottom) / h
	z := (far + near) / p

	m[0] = 2 / w
	m[4] = 0
	m[8] = 0
	m[12] = -x
	m[1] = 0
	m[5] = 2 / h
	m[9] = 0
	m[13] = -y
	m[2] = 0
	m[6] = 0
	m[10] = -2 / p
	m[14] = -z
	m[3] = 0
	m[7] = 0
	m[11] = 0
	m[15] = 1
}
