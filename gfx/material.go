package gfx

import (
	"sort"
)

const (
	MAT_PRIORITY_SHADER = iota
	MAT_PRIORITY_UNIFORM
)

type Material struct {
	params Parameters
}

func NewMaterial() *Material {
	mat := Material{make(Parameters, 8)}
	return &mat
}

func (m *Material) PreRender(ctx Context) {
	for _, param := range m.params {
		param.PreRender(ctx)
	}
}

func (m *Material) PostRender(ctx Context) {
	for _, param := range m.params {
		param.PostRender(ctx)
	}
}

func (m *Material) Add(para Parameter) {
	m.params = append(m.params, para)
}

func (m *Material) Sort() {
	sort.Sort(m.params)
}

func (m *Material) Lookup(name string) Parameter {

	return nil
}
