package gfx

import (
	"github.com/wibbe/glh/color"
	"github.com/wibbe/glh/math"
)

const (
	MATERIAL_MERGE = iota
	MATERIAL_REPLACE
)

const (
	UNIFORM_TYPE_FLOAT = iota
	UNIFORM_TYPE_VEC3
	UNIFORM_TYPE_MAT4
	UNIFORM_TYPE_COUNT
)

const (
	BUFFER_USAGE_STATIC = iota
	BUFFER_USGAE_DYNAMIC
)

const (
	BUFFER_TYPE_VERTEX = iota
	BUFFER_TYPE_INDEX
)

// The smallest building block in a material
type Parameter interface {
	Priority() int
	Name() string
	PreRender(ctx Context)
	PostRender(ctx Context)
}

type Parameters []Parameter

// Implementation of the sort.Interface interface
func (p Parameters) Len() int           { return len(p) }
func (p Parameters) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Parameters) Less(i, j int) bool { return p[i].Priority() > p[j].Priority() }

// Represents the value of a uniform in a shader
type Uniform interface {
	Parameter
	Type() int
	SetFloat(float32)
	SetVector3(math.Vector3)
	SetMatrix4(math.Matrix4)
}

// Represents a shader
type Shader interface {
	Parameter
	Uniforms() Parameters
}

type Buffer interface {
	Bind(Context)
	Uploadf32(data []float32)
	Uploadi16(data []uint16)
}

type Context interface {
	NewShader(name, vertexCode, fragmentCode string) (Shader, error)
	NewBuffer(bufferType, usage int) Buffer

	SetClearColor(color color.Color)

	Clear(colorBuffer, depthBuffer, stencilBuffer bool)
	UseCamera(camera *Camera)

	PushMaterial(material *Material, mergeStartegy int)
	PopMaterial()

	Draw(geom Geometry)
}
