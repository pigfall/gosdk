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
	pmath "github.com/pigfall/gosdk/math"
	"github.com/pigfall/gosdk/sdl3"
)

const vertexShaderSource = `#version 410 core
layout (location=0) in vec3 in_pos;
layout (location=1) in vec2 in_tex_coord;
out vec2 v_tex_coord;
uniform mat4 u_mvp;

void main(){
	gl_Position = u_mvp * vec4(in_pos,1.0);
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

	// setup viewport and clear color
	gl.GLViewport(0, 0, 800, 600)
	gl.GLClearColor(0.2, 0.3, 0.3, 1.0)

	imgFile, err := os.Open("examples/gl/assets/anya.jpeg")
	failOnError(err)
	defer imgFile.Close()

	jpeg, err := jpeg.Decode(imgFile)
	failOnError(err)
	rgbaImg := pimage.ToRGBA(jpeg)

	_, _, shaderProgram, err := gl.GLCompileProgram(vertexShaderSource, fragmentShaderSource)
	failOnError(err)
	gl.GLUseProgram(shaderProgram)

    // Cube vertices: position (x,y,z) + texcoord (u,v)
    positions := []float32{
        // Front face
        -50, 50, 50, 0.0, 1.0,
        -50, -50, 50, 0.0, 0.0,
        50, -50, 50, 1.0, 0.0,
        50, 50, 50, 1.0, 1.0,
        // Back face
        -50, 50, -50, 1.0, 1.0,
        -50, -50, -50, 1.0, 0.0,
        50, -50, -50, 0.0, 0.0,
        50, 50, -50, 0.0, 1.0,
    }

    // Indices for the cube (12 triangles)
    indices := []uint32{
        // front
        0, 1, 2, 0, 2, 3,
        // right
        3, 2, 6, 3, 6, 7,
        // back
        7, 6, 5, 7, 5, 4,
        // left
        4, 5, 1, 4, 1, 0,
        // top
        4, 0, 3, 4, 3, 7,
        // bottom
        1, 5, 6, 1, 6, 2,
    }

    // convert positions to bytes
    positionsBytes := make([]byte, len(positions)*4)
    for i, v := range positions {
        binary.LittleEndian.PutUint32(positionsBytes[i*4:(i+1)*4], math.Float32bits(v))
    }

    // convert indices to bytes (uint32)
    indicesBytes := make([]byte, len(indices)*4)
    for i, v := range indices {
        binary.LittleEndian.PutUint32(indicesBytes[i*4:(i+1)*4], v)
    }

    vao, err := gl.GLGenVertexArray()
    failOnError(err)
    gl.GLBindVertexArray(vao)

    // vertex buffer
    vbo, err := gl.GLGenBuffer()
    failOnError(err)
    gl.GLBindBuffer(gl.GLBindBufferTarget_ArrayBuffer, vbo)
    gl.GLBufferData(
        gl.GLBindBufferTarget_ArrayBuffer,
        len(positionsBytes),
        positionsBytes,
        gl.GLBufferDataUsage_StaticDraw,
    )
    // position
    gl.GLVertexAttribPointer(0, 3, gl.GLPrimitiveType_Float32, 5*4, 0)
    gl.GLEnableVertexAttribArray(0)
    // texcoord
    gl.GLVertexAttribPointer(1, 2, gl.GLPrimitiveType_Float32, 5*4, 3*4)
    gl.GLEnableVertexAttribArray(1)

    // element/index buffer
    ebo, err := gl.GLGenBuffer()
    failOnError(err)
    gl.GLBindBuffer(gl.GLBindBufferTarget_ElementArrayBuffer, ebo)
    gl.GLBufferData(
        gl.GLBindBufferTarget_ElementArrayBuffer,
        len(indicesBytes),
        indicesBytes,
        gl.GLBufferDataUsage_StaticDraw,
    )

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

	cameraUp := pmath.Vec3{0,1,0}
	cameraFront := pmath.Vec3{0,0,-1}
	cameraPos := pmath.Vec3{0,0,600}
	lookAt := pmath.Vector3Add(&cameraPos,&cameraFront)
	lookAtMatrix := pmath.Matrix4LookAt(
			cameraPos,
			lookAt,
			cameraUp,
	)
	projectionMatrix := pmath.Matrix4Perspective(
			45 * math.Pi/180,
			1.0,
			0.025,
			2048,
	)
	mvp := pmath.Matrix4Mul(&projectionMatrix,&lookAtMatrix)
    mvpUniformaLoc,err := gl.GLGetUniformLocation(shaderProgram,"u_mvp")
    failOnError(err)
    gl.GLUniformMatrix4fv(mvpUniformaLoc,&mvp)

    // Enable depth testing so the cube renders correctly in 3D
    gl.GLEnable(gl.GLCapability_DEPTH_TEST)

	running := true
	var ev sdl3.Event
	var cameraPosChanged = true
	const walkSpeed = 3
	for running {
		for sdl3.PollEvent(&ev) {
			switch ev.Type() {
			case sdl3.EventQuit:
				running = false
			case sdl3.EventKeyDown:
				keyEvent := ev.KeyboardEvent()
				keyCode := keyEvent.KeyCode()
				if keyCode == sdl3.K_ESCAPE{
					running = false
				}else if keyCode == sdl3.K_w{
					v := pmath.Vector3Multiple(&cameraFront,walkSpeed)
					cameraPos = pmath.Vector3Add(&cameraPos,&v)
					cameraPosChanged = true
				}else if keyCode == sdl3.K_a{
					v := pmath.Vector3Cross(&cameraFront,&cameraUp).Normalized()
					v = pmath.Vector3Multiple(&v,walkSpeed)
					cameraPos = pmath.Vector3Sub(&cameraPos,&v)
					cameraPosChanged = true
				}else if keyCode == sdl3.K_d{
					v := pmath.Vector3Cross(&cameraFront,&cameraUp).Normalized()
					v = pmath.Vector3Multiple(&v,walkSpeed)
					cameraPos = pmath.Vector3Add(&cameraPos,&v)
					cameraPosChanged = true

				}else if keyCode == sdl3.K_s{
					v := pmath.Vector3Multiple(&cameraFront,walkSpeed)
					cameraPos = pmath.Vector3Sub(&cameraPos,&v)
					cameraPosChanged = true
				}
			default:
			}
		}

        if cameraPosChanged{
            lookAt := pmath.Vector3Add(&cameraPos,&cameraFront)
            lookAtMatrix = pmath.Matrix4LookAt(
                    cameraPos,
                    lookAt,
                    cameraUp,
            )

            mvp := pmath.Matrix4Mul(&projectionMatrix,&lookAtMatrix)
            mvpUniformaLoc,err := gl.GLGetUniformLocation(shaderProgram,"u_mvp")
            failOnError(err)
            gl.GLUniformMatrix4fv(mvpUniformaLoc,&mvp)
        }

        // clear color and depth
        gl.GLClear(gl.GLClearMask_ColorBuffer | gl.GLClearMask_DepthBuffer)

        // draw elements from the EBO
        // 36 indices for the cube
        gl.GLDrawElements(gl.GLDrawArraysMode_Triangles, 36, 0)

		failOnError(win.Swap())
		cameraPosChanged = false
	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
