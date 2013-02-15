package gfx

import (
	"errors"
	"fmt"
	"github.com/wibbe/glh/color"
	gl "github.com/wibbe/glh/gl32c"
	_ "github.com/wibbe/glh/math"
)

const (
	GL_PRIORITY_SHADER = iota
	GL_PRIORITY_UNIFORM
)

type materialPair struct {
	mat      *Material
	strategy int
}

type gl32Context struct {
	materialStack []materialPair
	camera        *Camera
}

func NewGL32Context() (*gl32Context, error) {
	if err := gl.Init(); err != nil {
		return nil, errors.New(fmt.Sprintf("Could not create OpenGL context: %s", err))
	}

	return &gl32Context{
		materialStack: make([]materialPair, 0, 16),
	}, nil
}

func (ctx *gl32Context) NewShader(name, vertexCode, fragmentCode string) (Shader, error) {
	return newGL32Shader(name, vertexCode, fragmentCode)
}

func (ctx *gl32Context) NewBuffer(bufferType, usage int) Buffer {
	return nil
}

func (ctx *gl32Context) SetClearColor(c color.Color) {
	gl.ClearColor(gl.Float(c.R), gl.Float(c.G), gl.Float(c.B), gl.Float(c.A))
}

func (ctx *gl32Context) Clear(colorBuffer, depthBuffer, stencilBuffer bool) {
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

func (ctx *gl32Context) PushMaterial(mat *Material, mergeStrategy int) {
	ctx.materialStack = append(ctx.materialStack, materialPair{mat, mergeStrategy})
}

func (ctx *gl32Context) PopMaterial() {
}

func (ctx *gl32Context) UseCamera(camera *Camera) {
	ctx.camera = camera
}

func (ctx *gl32Context) Draw(geom Geometry) {
	ctx.preRenderMaterials()

	ctx.postRenderMaterials()
}

func (ctx *gl32Context) preRenderMaterials() {
	for _, mat := range ctx.materialStack {
		mat.mat.PreRender(ctx)
	}
}

func (ctx *gl32Context) postRenderMaterials() {
	for _, mat := range ctx.materialStack {
		mat.mat.PostRender(ctx)
	}
}
