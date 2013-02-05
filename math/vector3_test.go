package math

import (
    "testing"
)

func checkEq(t *testing.T, v1 , v2 Vector3) {
    if v1.X != v2.X || v1.Y != v2.Y || v1.Z != v2.Z {
        t.Fatalf("Vectors not equal %s != %s", v1, v2)
    }
}

func TestConstruction(test *testing.T) {
    vec := Vector3{1, 2, 3}

    if vec.X != 1 { test.Fatal("X != 1") }
    if vec.Y != 2 { test.Fatal("Y != 2") }
    if vec.Z != 3 { test.Fatal("Z != 3") }
}

func TestAdd(test *testing.T) {
    a := Vector3{1, 2, 3}
    b := Vector3{3, 2, 1}
    result := a.Add(b)

    checkEq(test, result, Vector3{4, 4, 4})
    checkEq(test, a, Vector3{1, 2, 3})
    checkEq(test, b, Vector3{3, 2, 1})
}

func TestSub(test *testing.T) {
    a := Vector3{4, 3, 2}
    b := Vector3{1, 1, 1}
    result := a.Sub(b)

    checkEq(test, result, Vector3{3, 2, 1})
    checkEq(test, a, Vector3{4, 3, 2})
    checkEq(test, b, Vector3{1, 1, 1})
}

func TestMul(test *testing.T) {
    a := Vector3{3, 2, 1}
    result := a.Mul(2)

    checkEq(test, result, Vector3{6, 4, 2})
    checkEq(test, a, Vector3{3, 2, 1})
}

func TestDot(test *testing.T) {
    a := Vector3{1, 0, 0}

    if a.Dot(Vector3{1, 0, 0}) != 1 { test.Fatal("Dot (1, 0, 0) . (1, 0, 0) != 1")}
    if a.Dot(Vector3{0, 1, 0}) != 0 { test.Fatal("Dot (1, 0, 0) . (0, 1, 0) != 0")}
    if a.Dot(Vector3{-1, 0, 0}) != -1 { test.Fatal("Dot (1, 0, 0) . (-1, 0, 0) != -1")}
}