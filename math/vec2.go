package math

import (
	stdmath "math"
)

type Vec2 struct {
	X float32
	Y float32
}

func Vec2Sub(a, b Vec2) Vec2 {
	return Vec2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func Vec2Add(a, b Vec2) Vec2 {
	return Vec2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (v1 Vec2) Normalize() Vec2 {
	l := 1.0 / v1.Len()
	return Vec2{v1.X * l, v1.Y * l}
}

func (v1 Vec2) Len() float32 {
	return float32(stdmath.Hypot(float64(v1.X), float64(v1.Y)))
}

func (v Vec2) Perpendicular() Vec2 {
	return Vec2{
		X: -v.Y,
		Y: v.X,
	}
}

func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
	}
}
