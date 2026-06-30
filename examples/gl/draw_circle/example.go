package main

import (
	"encoding/binary"
	"math"
	"runtime"

	"github.com/pigfall/gosdk/bytes"
	examplegl "github.com/pigfall/gosdk/examples/gl"
	"github.com/pigfall/gosdk/gl"
)

func init() {
	runtime.LockOSThread()
}

const vertexShaderSource = `#version 410 core
layout (location=0) in vec3 in_pos;

void main(){
	gl_Position = vec4(in_pos,1.0);
}

`

const fragmentShaderSource = `#version 410 core
out vec4 FragColor;

void main(){
	FragColor = vec4(1.0,0.0,0.0,1.0);
}
`

func main() {
	verticies := []float32{
		0.0, 0.0, 0.0, //Center point.
	}
	segements := 100
	for i := 0; i <= segements; i++ {
		rad := 360.0 / float32(segements) * float32(i) * (math.Pi / 180)
		x := math.Cos(float64(rad))
		y := math.Sin(float64(rad))
		verticies = append(verticies, float32(x), float32(y), 0.0)
	}
	verticesBytes := bytes.Float32SliceToBytes(verticies, binary.LittleEndian)

	examplegl.Run(
		func() { // Init function
			_, _, shaderProgram, err := gl.GLCompileProgram(vertexShaderSource, fragmentShaderSource)
			failOnError(err)
			gl.GLUseProgram(shaderProgram)

			vao, err := gl.GLGenVertexArray()
			failOnError(err)
			gl.GLBindVertexArray(vao)

			vbo, err := gl.GLGenBuffer()
			failOnError(err)
			gl.GLBindBuffer(gl.GLBindBufferTarget_ArrayBuffer, vbo)

			gl.GLBufferData(
				gl.GLBindBufferTarget_ArrayBuffer,
				len(verticesBytes),
				verticesBytes,
				gl.GLBufferDataUsage_StaticDraw,
			)
			gl.GLVertexAttribPointer(
				0,
				3,
				gl.GLPrimitiveType_Float32,
				0, 0,
			)
			gl.GLEnableVertexAttribArray(0)
		},
		func() { // update function

		},
		func() { //draw function
			gl.GLDrawArrays(
				gl.GLDrawArraysMode_Triangles_Fan, 0, int32(len(verticies)/3),
			)
		},
	)
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
