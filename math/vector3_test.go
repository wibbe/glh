package math

import (
	"math"
	"testing"
)

func checkVector3Eq(t *testing.T, v1, v2 Vector3) {
	if v1.X != v2.X || v1.Y != v2.Y || v1.Z != v2.Z {
		t.Fatalf("Vectors not equal %s != %s", v1, v2)
	}
}

func checkFloatEq(v1, v2 float32) bool {
	return math.Abs(float64(v1)-float64(v2)) < 0.000001
}

func TestVector3Construction(test *testing.T) {
	vec := Vector3{1, 2, 3}

	if vec.X != 1 {
		test.Fatal("X != 1")
	}
	if vec.Y != 2 {
		test.Fatal("Y != 2")
	}
	if vec.Z != 3 {
		test.Fatal("Z != 3")
	}
}

func TestVector3Add(test *testing.T) {
	a := Vector3{1, 2, 3}
	b := Vector3{3, 2, 1}
	result := a.Add(b)

	checkVector3Eq(test, result, Vector3{4, 4, 4})
	checkVector3Eq(test, a, Vector3{1, 2, 3})
	checkVector3Eq(test, b, Vector3{3, 2, 1})
}

func TestVector3Sub(test *testing.T) {
	a := Vector3{4, 3, 2}
	b := Vector3{1, 1, 1}
	result := a.Sub(b)

	checkVector3Eq(test, result, Vector3{3, 2, 1})
	checkVector3Eq(test, a, Vector3{4, 3, 2})
	checkVector3Eq(test, b, Vector3{1, 1, 1})
}

func TestVector3Mul(test *testing.T) {
	a := Vector3{3, 2, 1}
	result := a.Mul(2)

	checkVector3Eq(test, result, Vector3{6, 4, 2})
	checkVector3Eq(test, a, Vector3{3, 2, 1})
}

func TestVector3Dot(test *testing.T) {
	a := Vector3{1, 0, 0}

	if v := a.Dot(Vector3{1, 0, 0}); !checkFloatEq(v, 1.0) {
		test.Fatalf("Dot (1, 0, 0) . (1, 0, 0) != 1 is %f", v)
	}
	if v := a.Dot(Vector3{0, 1, 0}); !checkFloatEq(v, 0) {
		test.Fatalf("Dot (1, 0, 0) . (0, 1, 0) != 0 is %f", v)
	}
	if v := a.Dot(Vector3{-1, 0, 0}); !checkFloatEq(v, -1) {
		test.Fatalf("Dot (1, 0, 0) . (-1, 0, 0) != -1 is %f", v)
	}
}
