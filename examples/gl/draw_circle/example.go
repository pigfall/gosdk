package main

import (
	"encoding/binary"
	"runtime"

	"github.com/pigfall/gosdk/bytes"
	examplegl "github.com/pigfall/gosdk/examples/gl"
	"github.com/pigfall/gosdk/gl"
	"github.com/pigfall/gosdk/model"
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
	circleBuilder := model.Circle2DBuilder{}

	verticies := circleBuilder.Build()

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
