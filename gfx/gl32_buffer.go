package gfx

import (
	gl "github.com/wibbe/glh/gl32c"
	"unsafe"
)

type gl32Buffer struct {
	buffer gl.Uint
	target gl.Enum
	usage  gl.Enum
}

func newGL32Buffer(bufferType, usage int) *gl32Buffer {

}

func (b *gl32Buffer) Bind(ctx Context) {
	gl.BindBuffer(b.target, b.buffer)
}

func (b *gl32Buffer) Uploadf32(data []float32) {
	b.Bind(nil)
	var f float32
	gl.BufferData(b.target, gl.Sizeiptr(unsafe.Sizeof(f)*uintptr(len(data))), gl.Pointer(&data[0]), b.usage)
}

func (b *gl32Buffer) Uploadi16(data []uint16) {
	b.Bind(nil)
	var i uint16
	gl.BufferData(b.target, gl.Sizeiptr(unsafe.Sizeof(i)*uintptr(len(data))), gl.Pointer(&data[0]), b.usage)
}
