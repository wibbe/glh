package gfx

import (
	gl "github.com/wibbe/glh/gl32c"
)

type glShader struct {
	program gl.Int
}

func (s *glShader) PreRender(ctx *glContext) {
}

func (s *glShader) PostRender(ctx *glContext) {
}

func (s *glShader) Uniforms() Parameters {
	return nil
}

func (s *glShader) Priority() int { return GL_PRIORITY_SHADER }
