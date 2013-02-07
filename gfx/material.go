package gfx

const (
	MAT_PRIORITY_SHADER = iota
	MAT_PRIORITY_UNIFORM
)

type Material struct {
	params parameters
}

func NewMaterial() *Material {
	mat := Material{}
	//mat.params = make(map[string]Parameter)
	return &mat
}

func (m *Material) AddParameter(name string, para Parameter) {
	//m.params[name] = para
}

func (m *Material) Lookup(name string) Parameter {

	return nil
}
