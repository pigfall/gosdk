package sdl3

import (
	"github.com/Zyko0/go-sdl3/sdl"
)

type OpenGLContext struct {
	ctx *sdl.GLContextState
}

func (c *OpenGLContext) Destroy() {
	sdl.GL_DestroyContext(c.ctx)
}
