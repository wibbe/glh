package gfx

const (
	ATTRIB_POSITION = iota
	ATTRIB_NORMAL
	ATTRIB_COLOR
	ATTRIB_UV0
	ATTRIB_UV1
	ATTRIB_UV2
	ATTRIB_UV3
	ATTRIB_COUNT
)

type Geometry struct {
	buffers [ATTRIB_COUNT]Buffer
}

func (g *Geometry) AddBuffer(attrib int, buffer Buffer) {

}
