package sdl3

import (
	"github.com/pigfall/gosdk/gl"
	"time"
)

type runOption struct {
	initFunc func()           // called after gl.Init()
	onUpdate func(dt float32) // called after polled event.
	draw     func()           // called after gl clear

	windowTitle  string
	windowWidth  int
	windowHeight int

	openglMajorVersion int
	openglMinorVersion int
	openglProfile      OpenGLProfile
}

type RunOption func(o *runOption)

func Run(
	options ...RunOption,
) {
	unload, err := LoadEmbeddedSDL()
	failOnError(err)
	defer unload()

	failOnError(Init())
	defer Quit()

	option := &runOption{
		initFunc:     func() {},
		onUpdate:     func(dt float32) {},
		draw:         func() {},
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

	option.initFunc()

	running := true
	var ev Event
	prevTime := time.Now()
	for running {
		for PollEvent(&ev) {
			switch ev.Type() {
			case EventQuit:
				running = false
			default:
			}
		}

		now := time.Now()
		dt := now.Sub(prevTime).Seconds()
		prevTime = now
		option.onUpdate(float32(dt))

		gl.GLClear(gl.GLClearMask_ColorBuffer | gl.GLClearMask_DepthBuffer | gl.GLClearMask_StencilBuffer)
		option.draw()
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

func WithInitFunc(initF func()) func(o *runOption) {
	return func(o *runOption) {
		o.initFunc = initF
	}
}

func WithOnUpdate(onUpdate func(dt float32)) func(o *runOption) {
	return func(o *runOption) {
		o.onUpdate = onUpdate
	}
}

func WithDraw(draw func()) func(o *runOption) {
	return func(o *runOption) {
		o.draw = draw
	}
}
