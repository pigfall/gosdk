package sdl3

import(
	"github.com/Zyko0/go-sdl3/sdl"
)
	
type Event struct{
	E sdl.Event
}

func (e *Event) Type() EventType{
	return EventType(e.E.Type)
}
