package gfx

import (
	"github.com/wibbe/glh/math"
)

type Camera struct {
	Projection math.Matrix4
	View       math.Matrix4
}
