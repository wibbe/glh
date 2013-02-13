package math

import (
	"testing"
)

func checkColorEq(t *testing.T, c1, c2 Color) {
	if c1.R != c2.R || c1.G != c2.G || c1.B != c2.B || c1.A != c2.A {
		t.Fatalf("Colors not equal %s != %s", c1, c2)
	}
}

func TestColorConstruction(t *testing.T) {
	color := Color{1, 1, 1, 1}

	if color.R != 1 {
		t.Fatal("R != 1")
	}
	if color.G != 1 {
		t.Fatal("G != 1")
	}
	if color.B != 1 {
		t.Fatal("B != 1")
	}
}

func TestColorDecodeHex1(t *testing.T) {
	color, err := DecodeHexColor("ffffff")
	if err != nil {
		t.Fatal(err.Error())
	}

	checkColorEq(t, color, Color{1, 1, 1, 1})
}

func TestColorDecodeHex2(t *testing.T) {
	color, err := DecodeHexColor("ff0000")
	if err != nil {
		t.Fatal(err.Error())
	}

	checkColorEq(t, color, Color{1, 0, 0, 1})
}

func TestColorDecodeHex3(t *testing.T) {
	color, err := DecodeHexColor("00ff00")
	if err != nil {
		t.Fatal(err.Error())
	}

	checkColorEq(t, color, Color{0, 1, 0, 1})
}

func TestColorDecodeHex4(t *testing.T) {
	color, err := DecodeHexColor("0000ff")
	if err != nil {
		t.Fatal(err.Error())
	}

	checkColorEq(t, color, Color{0, 0, 1, 1})
}

func TestColorDecodeHex5(t *testing.T) {
	color, err := DecodeHexColor("ff000000")
	if err != nil {
		t.Fatal(err.Error())
	}

	checkColorEq(t, color, Color{1, 0, 0, 0})
}
