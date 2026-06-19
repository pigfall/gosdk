package math

import (
	"math"
)

type Matrix4 struct {
	Values []float32
}

func Matrix4Identity() Matrix4 {
	return Matrix4{
		Values: []float32{
			1.0, 0.0, 0.0, 0.0,
			0.0, 1.0, 0.0, 0.0,
			0.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 1.0,
		},
	}
}

func Matrix4Translate(v Vec3) Matrix4 {
	return Matrix4{
		Values: []float32{
			1.0, 0.0, 0.0, 0.0,
			0.0, 1.0, 0.0, 0.0,
			0.0, 0.0, 1.0, 0.0,
			v.X, v.Y, v.Z, 1.0,
		},
	}
}

func Matrix4Perspective(fovRad, aspect, zNear, zFar float32) Matrix4 {
	yScale := 1.0 / float32(math.Tan(float64(fovRad*0.5)))
	xScale := yScale / aspect

	return Matrix4{
		Values: []float32{
			xScale, 0.0, 0.0, 0.0,
			0.0, yScale, 0.0, 0.0,
			0.0, 0.0, zFar / (zNear - zFar), -1.0,
			0.0, 0.0, zNear * zFar / (zNear - zFar), 0.0,
		},
	}
}

func Matrix4Ortho(left, right, bottom, top, zNear, zFar float32) Matrix4 {
	return Matrix4{
		Values: []float32{
			2.0 / (right - left),
			0.0,
			0.0,
			0.0,
			0.0,
			2.0 / (top - bottom),
			0.0,
			0.0,
			0.0,
			0.0,
			-2.0 / (zFar - zNear),
			0.0,
			-(left + right) / (right - left),
			-(top + bottom) / (top - bottom),
			-(zFar + zNear) / (zFar - zNear),
			1.0,
		},
	}
}

func Matrix4LookAt(eyePos, at, up Vec3) Matrix4 {
	front := Vector3Sub(&at, &eyePos).Normalized()
	s := Vector3Cross(&front, &up).Normalized()
	//let u = s.cross(&f).normalized()?;
	u := Vector3Cross(&s, &front).Normalized()
	return Matrix4{
		Values: []float32{
			s.X, u.X, -front.X, 0.0,
			s.Y, u.Y, -front.Y, 0.0,
			s.Z, u.Z, -front.Z, 0.0,
			//-s.dot(&eye), -u.dot(&eye), f.dot(&eye), 1.0,
			-Vector3Dot(&s, &eyePos), -Vector3Dot(&u, &eyePos), Vector3Dot(&front, &eyePos), 1.0,
		},
	}
}

func Matrix4Mul(a, b *Matrix4) Matrix4 {
	return Matrix4{
		Values: []float32{
			a.Values[0]*b.Values[0] + a.Values[4]*b.Values[1] + a.Values[8]*b.Values[2] + a.Values[12]*b.Values[3],
			a.Values[1]*b.Values[0] + a.Values[5]*b.Values[1] + a.Values[9]*b.Values[2] + a.Values[13]*b.Values[3],
			a.Values[2]*b.Values[0] + a.Values[6]*b.Values[1] + a.Values[10]*b.Values[2] + a.Values[14]*b.Values[3],
			a.Values[3]*b.Values[0] + a.Values[7]*b.Values[1] + a.Values[11]*b.Values[2] + a.Values[15]*b.Values[3],
			a.Values[0]*b.Values[4] + a.Values[4]*b.Values[5] + a.Values[8]*b.Values[6] + a.Values[12]*b.Values[7],
			a.Values[1]*b.Values[4] + a.Values[5]*b.Values[5] + a.Values[9]*b.Values[6] + a.Values[13]*b.Values[7],
			a.Values[2]*b.Values[4] + a.Values[6]*b.Values[5] + a.Values[10]*b.Values[6] + a.Values[14]*b.Values[7],
			a.Values[3]*b.Values[4] + a.Values[7]*b.Values[5] + a.Values[11]*b.Values[6] + a.Values[15]*b.Values[7],
			a.Values[0]*b.Values[8] + a.Values[4]*b.Values[9] + a.Values[8]*b.Values[10] + a.Values[12]*b.Values[11],
			a.Values[1]*b.Values[8] + a.Values[5]*b.Values[9] + a.Values[9]*b.Values[10] + a.Values[13]*b.Values[11],
			a.Values[2]*b.Values[8] + a.Values[6]*b.Values[9] + a.Values[10]*b.Values[10] + a.Values[14]*b.Values[11],
			a.Values[3]*b.Values[8] + a.Values[7]*b.Values[9] + a.Values[11]*b.Values[10] + a.Values[15]*b.Values[11],
			a.Values[0]*b.Values[12] + a.Values[4]*b.Values[13] + a.Values[8]*b.Values[14] + a.Values[12]*b.Values[15],
			a.Values[1]*b.Values[12] + a.Values[5]*b.Values[13] + a.Values[9]*b.Values[14] + a.Values[13]*b.Values[15],
			a.Values[2]*b.Values[12] + a.Values[6]*b.Values[13] + a.Values[10]*b.Values[14] + a.Values[14]*b.Values[15],
			a.Values[3]*b.Values[12] + a.Values[7]*b.Values[13] + a.Values[11]*b.Values[14] + a.Values[15]*b.Values[15],
		},
	}
}

func Matrix4DecomposeTranslation(transform *Matrix4) Vec3 {
	return Vec3{X: transform.Values[12], Y: transform.Values[13], Z: transform.Values[14]}
}
