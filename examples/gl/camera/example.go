package main

import (
	"encoding/binary"
	"image/jpeg"
	"math"
	"os"
	"runtime"

	examplegl "github.com/pigfall/gosdk/examples/gl"
	"github.com/pigfall/gosdk/gl"
	pimage "github.com/pigfall/gosdk/image"
	"github.com/pigfall/gosdk/sdl3"
)

const vertexShaderSource = `#version 410 core
layout (location=0) in vec3 in_pos;
layout (location=1) in vec2 in_tex_coord;
out vec2 v_tex_coord;

void main(){
	gl_Position = vec4(in_pos,1.0);
	v_tex_coord = in_tex_coord;
}
`

const fragmentShaderSource = `#version 410 core
in vec2 v_tex_coord;
uniform sampler2D u_texture;
out vec4 FragColor;

void main(){
	FragColor = texture(u_texture,v_tex_coord);
}
`

func init() {
	runtime.LockOSThread()
}

func main() {
	unload, err := sdl3.LoadEmbeddedSDL()
	failOnError(err)
	defer unload()

	failOnError(sdl3.Init())
	defer sdl3.Quit()

	win, err := sdl3.CreateWindowWithOpenGL("demo", 800, 600, 0)
	failOnError(err)
	defer win.Destroy()

	major, minor := examplegl.OpenGLVersion()
	glCtx, err := win.CreateOpenGLContext(major, minor, sdl3.OpenGLCoreProfile)
	failOnError(err)
	defer glCtx.Destroy()

	failOnError(gl.Init())

	imgFile, err := os.Open("examples/gl/assets/anya.jpeg")
	failOnError(err)
	defer imgFile.Close()

	jpeg, err := jpeg.Decode(imgFile)
	failOnError(err)
	rgbaImg := pimage.ToRGBA(jpeg)

	_, _, shaderProgram, err := gl.GLCompileProgram(vertexShaderSource, fragmentShaderSource)
	failOnError(err)
	gl.GLUseProgram(shaderProgram)

	positions := []float32{
		-0.5, 0.5, 0.0, 0.0, 1.0,
		-0.5, -0.5, 0.0, 0.0, 0.0,
		0.5, -0.5, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.0, 1.0, 1.0,
	}

	positionsBytes := make([]byte, len(positions)*4)
	for i, v := range positions {
		binary.LittleEndian.PutUint32(positionsBytes[i*4:(i+1)*4], math.Float32bits(v))
	}

	vao, err := gl.GLGenVertexArray()
	failOnError(err)
	gl.GLBindVertexArray(vao)

	vbo, err := gl.GLGenBuffer()
	failOnError(err)
	gl.GLBindBuffer(gl.GLBindBufferTarget_ArrayBuffer, vbo)
	gl.GLBufferData(
		gl.GLBindBufferTarget_ArrayBuffer,
		len(positionsBytes),
		positionsBytes,
		gl.GLBufferDataUsage_StaticDraw,
	)
	gl.GLVertexAttribPointer(0, 3, gl.GLPrimitiveType_Float32, 5*4, 0)
	gl.GLEnableVertexAttribArray(0)
	// vertex attribute pointer.
	gl.GLVertexAttribPointer(
		1,
		2,
		gl.GLPrimitiveType_Float32,
		5*4,
		3*4,
	)
	gl.GLEnableVertexAttribArray(1)

	texture, err := gl.GLGenTexture()
	failOnError(err)
	gl.GLBindTexture(
		gl.GLBindTextureTarget_Texture2D,
		texture,
	)
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Min_Filter, gl.GLTextureParamter_MinFilter_Linear.ToGLType())
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Mag_Filter, gl.GLTextureParamter_MagFilter_Linear.ToGLType())
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Wrap_S, gl.GLTextureParamterWrapS_Repeat.ToGLType())
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Wrap_T, gl.GLTextureParamterWrapT_Repeat.ToGLType())
	gl.GLTexImage2D(
		gl.GLBindTextureTarget_Texture2D,
		0,
		gl.GLInternalColorFormat_RGBA,
		int32(rgbaImg.Rect.Size().X),
		int32(rgbaImg.Rect.Size().Y),
		0,
		gl.GLSourceImageColorFormat_RGBA,
		gl.GLPrimitiveType_UnsignedByte,
		rgbaImg.Pix,
	)

	running := true
	var ev sdl3.Event
	for running {
		for sdl3.PollEvent(&ev) {
			switch ev.Type() {
			case sdl3.EventQuit:
				running = false
			case sdl3.EventKeyDown:
				keyEvent := ev.KeyboardEvent()
				if keyEvent.KeyCode() == sdl3.K_ESCAPE{
					running = false
				}
			default:
			}
		}

		gl.GLClear(gl.GLClearMask_ColorBuffer)
		gl.GLDrawArrays(gl.GLDrawArraysMode_Triangles_Fan, 0, 4)

		failOnError(win.Swap())
	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
