package sdl3

import (
	"fmt"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
)

func LoadEmbeddedSDL()(unload func(),err error){
	defer func(){
		e := recover()
		if e != nil{
			err = fmt.Errorf("load sdl error: %v",e)
		}
	}()

	lib := binsdl.Load()
	unload = func(){
		lib.Unload()
	}

	return
}

func Init()error{
	return sdl.Init(sdl.INIT_VIDEO)
}

func Quit(){
	sdl.Quit()
}

func CreateWindowWithOpenGL(title string,width, height int,options WindowOption)(*Window,error){
	w,err :=  sdl.CreateWindow(title,width,height,sdl.WindowFlags(options) | sdl.WINDOW_OPENGL)
	if err != nil{
		return nil,err
	}
	return &Window{w},nil	
}

func PollEvent(ev *Event)bool {
	return sdl.PollEvent(&ev.E)
}
