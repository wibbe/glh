package gfx

type Parameter interface {
	PreRender(ctx Context)
	PostRender(ctx Context)
}

type Material struct {
	parameters map[string]Parameter
}

func NewMaterial() *Material {
	mat := Material{}
	mat.parameters = make(map[string]Parameter)
	return &mat
}

func (m *Material) AddParameter(name string, para Parameter) {
	m.parameters[name] = para
}

// -- Shader --

type Shader struct {
}
