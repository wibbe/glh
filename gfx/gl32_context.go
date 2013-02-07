package gfx

import (
	gl "github.com/wibbe/glh/gl32c"
	"github.com/wibbe/glh/math"
)

const (
	GL_PRIORITY_SHADER = iota
	GL_PRIORITY_UNIFORM
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

func (ctx *glContext) Clear(colorBuffer, depthBuffer, stencilBuffer bool) {
	var bits gl.Bitfield = 0

	if colorBuffer {
		bits = bits | gl.COLOR_BUFFER_BIT
	}
	if depthBuffer {
		bits = bits | gl.DEPTH_BUFFER_BIT
	}
	if stencilBuffer {
		bits = bits | gl.STENCIL_BUFFER_BIT
	}

	gl.Clear(bits)
}

func (ctx *glContext) PushMaterial(mat *Material, mergeStrategy int) {
	ctx.materialStack = append(ctx.materialStack, materialPair{mat, mergeStrategy})
}

func (ctx *glContext) PopMaterial() {
}

func (ctx *glContext) UseCamera(camera *Camera) {
	ctx.camera = camera
}

func (ctx *glContext) Draw(geom Geometry) {
	ctx.applyMaterials()
}

func (ctx *glContext) applyMaterials() {
}
