package gfx

type Parameter interface {
	Priority() int
	PreRender(ctx Context)
	PostRender(ctx Context)
}

type parameters []Parameter

func (p parameters) Len() int           { return len(p) }
func (p parameters) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p parameters) Less(i, j int) bool { return p[i].Priority() > p[j].Priority() }

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
