package gfx

import (
	gl "github.com/wibbe/glh/gl32c"
	"github.com/wibbe/glh/math"
)

type uniformSetFunc func(loc gl.Int, values []float32)

var uniformSetFuncs [UNIFORM_TYPE_COUNT]uniformSetFunc
var uniformSize = map[int]int{
	UNIFORM_TYPE_FLOAT: 1,
	UNIFORM_TYPE_VEC3:  3,
	UNIFORM_TYPE_MAT4:  16,
}

func init() {
	uniformSetFuncs[UNIFORM_TYPE_FLOAT] = func(loc gl.Int, values []float32) {
		gl.Uniform1fv(loc, 1, (*gl.Float)(&values[0]))
	}

	uniformSetFuncs[UNIFORM_TYPE_VEC3] = func(loc gl.Int, values []float32) {
		gl.Uniform3fv(loc, 1, (*gl.Float)(&values[0]))
	}

	uniformSetFuncs[UNIFORM_TYPE_MAT4] = func(loc gl.Int, values []float32) {
		gl.UniformMatrix4fv(loc, 1, gl.FALSE, (*gl.Float)(&values[0]))
	}
}

type gl32Uniform struct {
	location gl.Int
	dataType int
	changed  bool
	value    []float32
	name     string
}

func newGL32Uniform(location gl.Int, dataType int, name string) *gl32Uniform {
	return &gl32Uniform{location, dataType, false, make([]float32, uniformSize[dataType]), name}
}

func (u *gl32Uniform) SetFloat(val float32) {
	if u.dataType == UNIFORM_TYPE_FLOAT {
		u.value[0] = val
		u.changed = true
	}
}

func (u *gl32Uniform) SetVector3(vec math.Vector3) {
	if u.dataType == UNIFORM_TYPE_VEC3 {
		u.value[0] = vec.X
		u.value[1] = vec.Y
		u.value[2] = vec.Z
		u.changed = true
	}
}

func (u *gl32Uniform) SetMatrix4(vec math.Matrix4) {
	if u.dataType == UNIFORM_TYPE_MAT4 {
		u.changed = true
	}
}

func (u *gl32Uniform) PreRender(ctx Context) {
	if u.changed {
		u.changed = false
		set := uniformSetFuncs[u.dataType]
		set(u.location, u.value)
	}
}

func (u *gl32Uniform) PostRender(ctx Context) {
	// Empty implementation
}

func (u *gl32Uniform) Name() string  { return u.name }
func (u *gl32Uniform) Type() int     { return u.dataType }
func (u *gl32Uniform) Priority() int { return GL_PRIORITY_UNIFORM }
