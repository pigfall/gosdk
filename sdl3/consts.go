package sdl3

import(
	"github.com/Zyko0/go-sdl3/sdl"
)

type WindowOption sdl.WindowFlags
type OpenGLProfile int32
type EventType uint32

const(
	WindowOptionOpenGL WindowOption = WindowOption(sdl.WINDOW_OPENGL)
)

const(
		OpenGLCoreProfile OpenGLProfile = sdl.GL_CONTEXT_PROFILE_CORE
		OpenGLCompatibilityProfile OpenGLProfile = sdl.GL_CONTEXT_PROFILE_COMPATIBILITY
		OpenGLESProfile OpenGLProfile = sdl.GL_CONTEXT_PROFILE_ES
)

const (
		EventQuit EventType = EventType(sdl.EVENT_QUIT)
)


