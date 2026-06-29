package main

import (
	"encoding/binary"
	"image/jpeg"
	"math"
	"os"
	"runtime"

	"github.com/pigfall/gosdk/gl"
	"github.com/pigfall/gosdk/image"
	pmath "github.com/pigfall/gosdk/math"
	"github.com/pigfall/gosdk/sdl3"
)

const vertexShaderSource = `#version 410 core
layout (location=0) in vec3 pos;
layout (location=1) in vec2 texCoord;
out vec2 v_tex_coord;
uniform mat4 u_mvp;

void main(){
	gl_Position = u_mvp * vec4(pos,1.0);
	v_tex_coord = texCoord;
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
	must(err)
	defer unload()

	must(sdl3.Init())
	defer sdl3.Quit()

	window, err := sdl3.CreateWindowWithOpenGL("demo", 800, 600, 0)
	must(err)
	defer window.Destroy()

	glCtx, err := window.CreateOpenGLContext(4, 1, sdl3.OpenGLCoreProfile)
	must(err)
	defer glCtx.Destroy()

	must(gl.Init())

	lookAt := pmath.Matrix4LookAt(
		pmath.Vec3{X: 0.0, Y: 0.0, Z: 600.0},
		pmath.Vec3{X: 0.0, Y: 0.0, Z: -1.0},
		pmath.Vec3{X: 0.0, Y: 1.0, Z: 0.0},
	)
	projection := pmath.Matrix4Perspective(
		45.0*math.Pi/180,
		1.0,
		0.025,
		2048.0,
	)
	mvp := pmath.Matrix4Mul(&projection, &lookAt)

	f, err := os.Open("examples/gl/assets/anya.jpeg")
	must(err)
	defer f.Close()
	img, err := jpeg.Decode(f)
	must(err)
	rgbaImg := image.ToRGBA(img)

	_, _, shaderProgram, err := gl.GLCompileProgram(vertexShaderSource, fragmentShaderSource)
	must(err)
	gl.GLUseProgram(shaderProgram)

	mvpLoc, err := gl.GLGetUniformLocation(shaderProgram, "u_mvp")
	must(err)

	vao, err := gl.GLGenVertexArray()
	must(err)
	gl.GLBindVertexArray(vao)

	gl.GLUniformMatrix4fv(mvpLoc, &mvp)

	positions := []float32{
		-50.0, 50.0, 0.0,
		-50.0, -50.0, 0.0,
		50.0, -50.0, 0.0,
		50.0, 50.0, 0.0,
	}

	textureCoords := []float32{
		0.0, 1.0,
		0.0, 0.0,
		1.0, 0.0,
		1.0, 1.0,
	}

	var positionsBytes []byte = make([]byte, len(positions)*4)
	for i, v := range positions {
		binary.LittleEndian.PutUint32(positionsBytes[i*4:(i+1)*4], math.Float32bits(v))
	}
	var textureCoordBytes []byte = make([]byte, len(textureCoords)*4)
	for i, v := range textureCoords {
		binary.LittleEndian.PutUint32(textureCoordBytes[i*4:(i+1)*4], math.Float32bits(v))
	}

	gl.GLActiveTexture(gl.GLActiveTextureTarget_Texture0)
	texture, err := gl.GLGenTexture()
	must(err)
	gl.GLBindTexture(gl.GLBindTextureTarget_Texture2D, texture)
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
	// Avoid requiring mipmaps for minification, otherwise sampling can return black.
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Min_Filter, gl.GLTextureParamter_MinFilter_Linear.ToGLType())
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Mag_Filter, gl.GLTextureParamter_MagFilter_Linear.ToGLType())
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Wrap_S, gl.GLTextureParamterWrapS_Repeat.ToGLType())
	gl.GLTexParameteri(gl.GLBindTextureTarget_Texture2D, gl.GLTextureParamterName_Texture_Wrap_T, gl.GLTextureParamterWrapT_Repeat.ToGLType())

	//uTextureLoc, err := gl.GLGetUniformLocation(shaderProgram, "u_texture")
	//must(err)
	//gl.GLUniform1i(uTextureLoc, 0)

	vbo, err := gl.GLGenBuffer()
	must(err)
	gl.GLBindBuffer(gl.GLBindBufferTarget_ArrayBuffer, vbo)
	gl.GLBufferData(
		gl.GLBindBufferTarget_ArrayBuffer,
		len(positionsBytes)+len(textureCoordBytes),
		nil,
		gl.GLBufferDataUsage_StaticDraw,
	)
	gl.GLBufferSubData(
		gl.GLBindBufferTarget_ArrayBuffer,
		0,
		positionsBytes,
	)
	gl.GLBufferSubData(
		gl.GLBindBufferTarget_ArrayBuffer,
		len(positionsBytes),
		textureCoordBytes,
	)
	gl.GLVertexAttribPointer(0, 3, gl.GLPrimitiveType_Float32, 3*4, 0)
	gl.GLEnableVertexAttribArray(0)
	gl.GLVertexAttribPointer(1, 2, gl.GLPrimitiveType_Float32, 2*4, int32(len(positionsBytes)))
	gl.GLEnableVertexAttribArray(1)

	var ev sdl3.Event
	running := true
	for running {
		for sdl3.PollEvent(&ev) {
			switch ev.Type() {
			case sdl3.EventQuit:
				running = false
			default:
			}
		}

		gl.GLClear(gl.GLClearMask_ColorBuffer)
		gl.GLDrawArrays(gl.GLDrawArraysMode_Triangles_Fan, 0, 4)

		window.Swap()
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
