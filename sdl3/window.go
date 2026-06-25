package sdl3

import (
	"github.com/Zyko0/go-sdl3/sdl"
)

type Window struct {
	w *sdl.Window
}

func (w *Window) CreateOpenGLContext(
	majorVersion,
	minorVersion int32,
	profile OpenGLProfile,
) (*OpenGLContext, error) {
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, majorVersion)
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, minorVersion)
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, int32(profile))

	glCtx, err := sdl.GL_CreateContext(w.w)
	if err != nil {
		return nil, err
	}

	return &OpenGLContext{glCtx}, nil
}

func (w *Window) Destroy() {
	w.w.Destroy()
}

func (w *Window) Swap() error {
	return sdl.GL_SwapWindow(w.w)
}
