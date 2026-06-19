package math

import (
	stdmath "math"
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v *Vec3) Len() float32 {
	return float32(stdmath.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vec3) Normalized() Vec3 {
	vectorLen := v.Len()
	invLen := 1.0 / vectorLen
	return Vec3{
		X: v.X * invLen,
		Y: v.Y * invLen,
		Z: v.Z * invLen,
	}
}

func Vector3Sub(a, b *Vec3) Vec3 {
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func Vector3Add(a, b *Vec3) Vec3 {
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Vector3Cross(a, b *Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func Vector3Dot(a, b *Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
