package sdl3

import (
	"github.com/pigfall/gosdk/gl"
)

type runOption struct {
	windowTitle  string
	windowWidth  int
	windowHeight int

	openglMajorVersion int
	openglMinorVersion int
	openglProfile      OpenGLProfile
}

type RunOption func(o *runOption)

func Run(
	init func(),
	onUpdate func(),
	draw func(),
	options ...RunOption,
) {
	unload, err := LoadEmbeddedSDL()
	failOnError(err)
	defer unload()

	failOnError(Init())
	defer Quit()

	option := &runOption{
		windowTitle:  "demo",
		windowWidth:  600,
		windowHeight: 600,

		openglMajorVersion: 4,
		openglMinorVersion: 1,
		openglProfile:      OpenGLCoreProfile,
	}
	for _, opt := range options {
		opt(option)
	}

	window, err := CreateWindowWithOpenGL(option.windowTitle, option.windowWidth, option.windowHeight, 0)
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
	var ev Event
	for running {
		for PollEvent(&ev) {
			switch ev.Type() {
			case EventQuit:
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
