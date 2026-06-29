package examplegl

import (
	"github.com/pigfall/gosdk/gl"
	"github.com/pigfall/gosdk/sdl3"
)

type runOption struct {
	windowTitle  string
	windowWidth  int
	windowHeight int

	openglMajorVersion int
	openglMinorVersion int
	openglProfile      sdl3.OpenGLProfile
}

type RunOption func(o *runOption)

func Run(
	init func(),
	onUpdate func(),
	draw func(),
	options ...RunOption,
) {
	unload, err := sdl3.LoadEmbeddedSDL()
	failOnError(err)
	defer unload()

	failOnError(sdl3.Init())
	defer sdl3.Quit()

	option := &runOption{
		windowTitle:  "demo",
		windowWidth:  600,
		windowHeight: 600,

		openglMajorVersion: 4,
		openglMinorVersion: 1,
		openglProfile:      sdl3.OpenGLCoreProfile,
	}
	for _, opt := range options {
		opt(option)
	}

	window, err := sdl3.CreateWindowWithOpenGL(option.windowTitle, option.windowWidth, option.windowHeight, 0)
	failOnError(err)
	defer window.Destroy()

	glCtx, err := window.CreateOpenGLContext(
		int32(option.openglMajorVersion),
		int32(option.openglMinorVersion),
		option.openglProfile,
	)
	failOnError(err)
	defer glCtx.Destroy()

	failOnError(gl.Init())

	init()
	running := true
	var ev sdl3.Event
	for running {
		for sdl3.PollEvent(&ev) {
			switch ev.Type() {
			case sdl3.EventQuit:
				running = false
			default:
			}
		}

		onUpdate()

		gl.GLClear(gl.GLClearMask_ColorBuffer | gl.GLClearMask_DepthBuffer | gl.GLClearMask_StencilBuffer)
		draw()
		failOnError(window.Swap())
	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func WithWindowWidth(width int) func(o *runOption) {
	return func(o *runOption) {
		o.windowWidth = width
	}
}

func WithWindowHeight(height int) func(o *runOption) {
	return func(o *runOption) {
		o.windowHeight = height
	}
}
