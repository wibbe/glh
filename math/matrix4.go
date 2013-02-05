package math

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
