package gfx

import ()

const (
	MATERIAL_MERGE   = 0
	MATERIAL_REPLACE = 1
)

type Context interface {
	Clear()
	UseCamera(camera *Camera)

	SetShader(shader Shader)
	SetUniform(uniform Uniform)

	PushMaterial(mat *Material, mergeStartegy int)
	PopMaterial()
}
