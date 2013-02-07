package gfx

type uniformSetFunc func(loc gl.Int, values []float32)

var uniformSetFuncs [UNIFORM_TYPE_COUNT]uniformSetFunc

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

type glUniform struct {
	value    []float32
	dataType int
	location gl.Int
	changed  bool
	name     string
}

func (u *glUniform) SetFloat(val float32) {
	if u.dataType == UNIFORM_TYPE_FLOAT {
		u.value[0] = val
		u.changed = true
	}
}

func (u *glUniform) SetVector3(vec math.Vector3) {
	if u.dataType == UNIFORM_TYPE_VEC3 {
		u.value[0] = vec.X
		u.value[1] = vec.Y
		u.value[2] = vec.Z
		u.changed = true
	}
}

func (u *glUniform) SetMatrix4(vec math.Matrix4) {
	if u.dataType == UNIFORM_TYPE_MAT4 {
		u.changed = true
	}
}

func (u *glUniform) PreRender(ctx *glContext) {
	if u.changed {
		u.changed = false
		set := uniformSetFuncs[u.dataType]
		set(u.location, u.value)
	}
}

func (u *glUniform) PostRender(ctx *glContext) {
	// Empty implementation
}

func (u *glUniform) Name() string  { return u.name }
func (u *glUniform) Type() int     { return u.dataType }
func (u *glUniform) Priority() int { return GL_PRIORITY_UNIFORM }
