package radian

import (
	"fmt"
	"math"
)

func Convert() {
	var input float64
	_, _ = fmt.Scanln(&input)
	radian := input * (math.Pi / 180.0)
	fmt.Printf("%.6f\n", radian)
}
