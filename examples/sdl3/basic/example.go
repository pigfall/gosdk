package main

import (
	"fmt"
	"runtime"

	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/pigfall/gosdk/sdl3"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	unload, err := sdl3.LoadEmbeddedSDL()
	must(err)
	defer unload()

	win, err := sdl3.CreateWindowWithOpenGL("demo", 800, 800, 0)
	must(err)
	defer win.Destroy()

	running := true
	var ev sdl3.Event
	for running {
		for sdl3.PollEvent(&ev) {
			switch ev.Type() {
			case sdl3.EventQuit:
				running = false
			case sdl3.EventKeyDown, sdl3.EventKeyUp:
				if handleKeyboardEvent(&ev) {
					running = false
				}
			case sdl3.EventMouseMotion:
				mouseEv := ev.E.MouseMotionEvent()
				mouseEv.
			default:
			}
		}
	}
}

func handleKeyboardEvent(ev *sdl3.Event) bool {
	keyEvent := ev.E.KeyboardEvent()
	if keyEvent == nil {
		return false
	}

	if ev.Type() != sdl3.EventKeyDown {
		return false
	}

	switch keyEvent.Key {
	case sdl.K_ESCAPE:
		fmt.Println("escape pressed, exiting")
		return true
	case sdl.K_LEFT:
		fmt.Println("left key pressed")
	case sdl.K_RIGHT:
		fmt.Println("right key pressed")
	}

	return false
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
