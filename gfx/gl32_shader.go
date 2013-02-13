package gfx

import (
	"errors"
	"fmt"
	gl "github.com/wibbe/glh/gl32c"
	"strings"
)

var validUniformTypes = map[gl.Enum]bool{
	gl.FLOAT:      true,
	gl.FLOAT_VEC3: true,
	gl.FLOAT_MAT4: true,
}

var glToUniform = map[gl.Enum]int{
	gl.FLOAT:      UNIFORM_TYPE_FLOAT,
	gl.FLOAT_VEC3: UNIFORM_TYPE_VEC3,
	gl.FLOAT_MAT4: UNIFORM_TYPE_MAT4,
}
var currentActiveShader gl.Uint = 0

type gl32Shader struct {
	name     string
	program  gl.Uint
	uniforms Parameters
}

func compileShader(code string, shaderType gl.Enum) (gl.Uint, error) {
	shader := gl.CreateShader(shaderType)

	glCode := gl.GLString(code)
	defer gl.GLStringFree(glCode)

	gl.ShaderSource(shader, 1, &glCode, nil)
	gl.CompileShader(shader)

	var status gl.Int
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == 0 {
		var logLength gl.Int
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		msg := gl.GLStringAlloc(gl.Sizei(logLength))
		defer gl.GLStringFree(msg)

		gl.GetShaderInfoLog(shader, gl.Sizei(logLength), nil, msg)
		gl.DeleteShader(shader)

		return 0, errors.New(gl.GoString(msg))
	}

	return shader, nil
}

func newGL32Shader(name, vertexCode, fragmentCode string) (*gl32Shader, error) {

	vertexShader, err := compileShader(vertexCode, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compileShader(fragmentCode, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	shader := &gl32Shader{name, gl.CreateProgram(), nil}
	gl.AttachShader(shader.program, vertexShader)
	gl.AttachShader(shader.program, fragmentShader)

	gl.BindAttribLocation(shader.program, ATTRIB_POSITION, gl.GLString("position"))
	gl.BindAttribLocation(shader.program, ATTRIB_NORMAL, gl.GLString("normal"))
	gl.BindAttribLocation(shader.program, ATTRIB_UV0, gl.GLString("uv0"))
	gl.BindAttribLocation(shader.program, ATTRIB_UV1, gl.GLString("uv1"))
	gl.BindAttribLocation(shader.program, ATTRIB_UV2, gl.GLString("uv2"))
	gl.BindAttribLocation(shader.program, ATTRIB_UV3, gl.GLString("uv3"))
	gl.BindAttribLocation(shader.program, ATTRIB_COLOR, gl.GLString("color"))

	gl.LinkProgram(shader.program)

	var status gl.Int
	gl.GetProgramiv(shader.program, gl.LINK_STATUS, &status)
	if status != 1 {
		gl.DeleteShader(vertexShader)
		gl.DeleteShader(fragmentShader)
		gl.DeleteProgram(shader.program)

		return nil, errors.New(fmt.Sprintf("Could not link shader program '%s'", name))
	}

	shader.buildUniformList()

	return shader, nil
}

func (s *gl32Shader) buildUniformList() {
	s.uniforms = make(Parameters, 4)

	s.PreRender(nil)

	var uniformCount gl.Int
	gl.GetProgramiv(s.program, gl.ACTIVE_UNIFORMS, &uniformCount)
	uniformName := gl.GLStringAlloc(128)
	defer gl.GLStringFree(uniformName)

	textureUnit := 0

	for i := 0; i < int(uniformCount); i++ {
		var uniformType gl.Enum
		var uniformSize gl.Int
		var nameLength gl.Sizei

		gl.GetActiveUniform(s.program, gl.Uint(i), 128, &nameLength, &uniformSize, &uniformType, uniformName)

		// Ignore build in uniforms
		if !strings.HasPrefix(gl.GoString(uniformName), "gl_") {
			location := gl.GetUniformLocation(s.program, uniformName)

			if uniformType == gl.SAMPLER_2D {
				gl.Uniform1i(location, gl.Int(textureUnit))
				textureUnit++
			}

			if validUniformTypes[uniformType] {
				s.uniforms = append(s.uniforms, newGL32Uniform(location, glToUniform[uniformType], gl.GoStringN(uniformName, nameLength)))
			}
		}
	}
}

func (s *gl32Shader) PreRender(Context) {
	if currentActiveShader != s.program {
		currentActiveShader = s.program
		gl.UseProgram(s.program)
	}
}

func (s *gl32Shader) PostRender(Context) {
}

func (s *gl32Shader) Uniforms() Parameters {
	return s.uniforms
}

func (s *gl32Shader) Priority() int { return GL_PRIORITY_SHADER }
func (s *gl32Shader) Name() string  { return s.name }
