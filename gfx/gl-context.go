package gfx

import (
	gl "github.com/wibbe/glh/gl32c"
)

type materialPair struct {
	mat      *Material
	strategy int
}

type glContext struct {
	materialStack []materialPair
	camera        *Camera
}

func NewGL32Context() *glContext {
	return &glContext{
		materialStack: make([]materialPair, 0, 16),
	}
}

func (ctx *glContext) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (ctx *glContext) PushMaterial(mat *Material, mergeStrategy int) {
	ctx.materialStack = append(ctx.materialStack, materialPair{mat, mergeStrategy})
}

func (ctx *glContext) PopMaterial() {
}

func (ctx *glContext) UseCamera(camera *Camera) {
	ctx.camera = camera
}
