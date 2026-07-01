package model

import "math"

type Circle2DBuilder struct{}

func (c *Circle2DBuilder) Build() []float32 {
	segements := 100
	verticies := make([]float32, ((segements + 2) * 3))
	verticies[0] = 0
	verticies[1] = 0
	verticies[2] = 0

	radius := 0.5
	for i := 0; i <= segements; i++ {
		radian := 360.0 / float64(segements) * float64(i) * math.Pi / 180.0
		x := radius * math.Cos(radian)
		y := radius * math.Sin(radian)
		verticies[(i+1)*3] = float32(x)
		verticies[(i+1)*3+1] = float32(y)
		verticies[(i+1)*3+2] = float32(0.0)
	}

	return verticies
}
