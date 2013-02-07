package gfx

import (
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

// The smallest building block in a material
type Parameter interface {
	Priority() int
	PreRender(ctx Context)
	PostRender(ctx Context)
}

type Parameters []Parameter

// Implementation of the sort.Interface interface
func (p Parameters) Len() int           { return len(p) }
func (p Parameters) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Parameters) Less(i, j int) bool { return p[i].Priority() > p[j].Priority() }

// Represents a shader
type Shader interface {
	Parameter
	Uniforms() Parameters
}

// Represents the value of a uniform in a shader
type Uniform interface {
	Parameter
	Type() int
	Name() string
	SetFloat(float32)
	SetVector3(math.Vector3)
	SetMatrix4(math.Matrix4)
}

type Geometry interface {
}

type Context interface {
	Clear(colorBuffer, depthBuffer, stencilBuffer bool)
	UseCamera(camera *Camera)

	PushMaterial(material *Material, mergeStartegy int)
	PopMaterial()

	Draw(geom Geometry)
}
