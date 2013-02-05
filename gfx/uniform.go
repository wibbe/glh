package gfx

import (
	gl "github.com/wibbe/glh/gl32c"
	"github.com/wibbe/glh/math"
)

const (
	UNIFORM_TYPE_VEC3 = iota
	UNIFORM_TYPE_VEC4
	UNIFORM_TYPE_MAT4
)

type Uniform struct {
	value    []float32
	dataType int
	location gl.Int
	changed  bool
}

func NewUniform(dataType int) *Uniform {
	size := 0
	switch dataType {
	case UNIFORM_TYPE_VEC3:
		size = 3
	case UNIFORM_TYPE_VEC4:
		size = 4
	case UNIFORM_TYPE_MAT4:
		size = 16
	}

	return &Uniform{make([]float32, size), dataType, 0, true}
}

func (u *Uniform) Priority() int {
	return MAT_PRIORITY_UNIFORM
}

func (u *Uniform) PreRender(ctx Context) {
	if u.changed {
		u.changed = false

		switch u.dataType {
		case UNIFORM_TYPE_VEC3:
			gl.Uniform3fv(u.location, 1, (*gl.Float)(&u.value[0]))
		case UNIFORM_TYPE_VEC4:
			gl.Uniform4fv(u.location, 1, (*gl.Float)(&u.value[0]))
		case UNIFORM_TYPE_MAT4:
			gl.UniformMatrix4fv(u.location, 1, gl.FALSE, (*gl.Float)(&u.value[0]))
		}
	}
}

func (u *Uniform) PostRender(ctx Context) {
}

func (u *Uniform) SetVec3(vec math.Vector3) {
	if u.dataType == UNIFORM_TYPE_VEC3 {
		u.value[0] = vec.X
		u.value[1] = vec.Y
		u.value[2] = vec.Z
	}
}
