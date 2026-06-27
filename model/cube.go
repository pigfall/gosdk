package model

type CubeBuilder struct {
	Scalar           float32
	WithTextureCoord bool
}

// Build generates vertex data and indices for a cube suitable for OpenGL.
// Vertex layout depends on withTextureCoord:
// - when withTextureCoord == true: each vertex is [x,y,z,u,v]
// - otherwise: each vertex is [x,y,z]
// Scalar is treated as the half-extent of the cube (i.e. coordinates range [-Scalar, Scalar]).
// Returns vertices and element indices (36 indices, 12 triangles using 24 distinct vertices).
func (b *CubeBuilder) Build() ([]float32, []uint32) {
	s := b.Scalar
	if s == 0 {
		s = 1.0
	}

	// define the 24 positions (4 vertices per face) and texture coords when requested
	// order matches typical cube: front, back, right, left, top, bottom
	if b.WithTextureCoord {
		vertices := []float32{
			// Front face
			-s, s, s, 0.0, 1.0,
			-s, -s, s, 0.0, 0.0,
			s, -s, s, 1.0, 0.0,
			s, s, s, 1.0, 1.0,
			// Back face
			-s, s, -s, 1.0, 1.0,
			-s, -s, -s, 1.0, 0.0,
			s, -s, -s, 0.0, 0.0,
			s, s, -s, 0.0, 1.0,
			// Right face
			s, s, s, 0.0, 1.0,
			s, -s, s, 0.0, 0.0,
			s, -s, -s, 1.0, 0.0,
			s, s, -s, 1.0, 1.0,
			// Left face
			-s, s, -s, 0.0, 1.0,
			-s, -s, -s, 0.0, 0.0,
			-s, -s, s, 1.0, 0.0,
			-s, s, s, 1.0, 1.0,
			// Top face
			-s, s, -s, 0.0, 1.0,
			-s, s, s, 0.0, 0.0,
			s, s, s, 1.0, 0.0,
			s, s, -s, 1.0, 1.0,
			// Bottom face
			-s, -s, s, 0.0, 1.0,
			-s, -s, -s, 0.0, 0.0,
			s, -s, -s, 1.0, 0.0,
			s, -s, s, 1.0, 1.0,
		}

		indices := []uint32{
			// front
			0, 1, 2, 0, 2, 3,
			// back
			4, 5, 6, 4, 6, 7,
			// right
			8, 9, 10, 8, 10, 11,
			// left
			12, 13, 14, 12, 14, 15,
			// top
			16, 17, 18, 16, 18, 19,
			// bottom
			20, 21, 22, 20, 22, 23,
		}

		return vertices, indices
	}

	// without texture coords: only positions (3 floats per vertex)
	vertices := []float32{
		// Front face
		-s, s, s,
		-s, -s, s,
		s, -s, s,
		s, s, s,
		// Back face
		-s, s, -s,
		-s, -s, -s,
		s, -s, -s,
		s, s, -s,
		// Right face
		s, s, s,
		s, -s, s,
		s, -s, -s,
		s, s, -s,
		// Left face
		-s, s, -s,
		-s, -s, -s,
		-s, -s, s,
		-s, s, s,
		// Top face
		-s, s, -s,
		-s, s, s,
		s, s, s,
		s, s, -s,
		// Bottom face
		-s, -s, s,
		-s, -s, -s,
		s, -s, -s,
		s, -s, s,
	}

	indices := []uint32{
		0, 1, 2, 0, 2, 3,
		4, 5, 6, 4, 6, 7,
		8, 9, 10, 8, 10, 11,
		12, 13, 14, 12, 14, 15,
		16, 17, 18, 16, 18, 19,
		20, 21, 22, 20, 22, 23,
	}

	return vertices, indices
}
