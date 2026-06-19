package math

import (
	"math"
)

func RadianFromDegree(degree float32) float32 {
	return degree * (math.Pi / 180)
}
