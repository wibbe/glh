package color

import (
	"encoding/hex"
)

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func FromHex(value string) (Color, error) {
	if bytes, err := hex.DecodeString(value); err == nil {
		color := Color{0, 0, 0, 1}

		if len(bytes) >= 3 {
			color.R = float32(bytes[0]) / 255.0
			color.G = float32(bytes[1]) / 255.0
			color.B = float32(bytes[2]) / 255.0
		}

		if len(bytes) == 4 {
			color.A = float32(bytes[3]) / 255.0
		}

		return color, nil
	} else {
		return Color{}, err
	}

	return Color{}, nil
}
