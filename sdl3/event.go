package sdl3

import (
	"github.com/Zyko0/go-sdl3/sdl"
)

type Event struct {
	E sdl.Event
}

type KeyboardEvent struct {
	E *sdl.KeyboardEvent
}

type MouseMotionEvent struct {
	E *sdl.MouseMotionEvent
}

func (e *Event) Type() EventType {
	return EventType(e.E.Type)
}

func (e *Event) KeyboardEvent() *KeyboardEvent {
	return &KeyboardEvent{
		E: e.E.KeyboardEvent(),
	}
}

func (e *Event) MouseMotionEvent() *MouseMotionEvent {
	return &MouseMotionEvent{
		E: e.E.MouseMotionEvent(),
	}
}

func (e *KeyboardEvent) KeyCode() KeyCode {
	return KeyCode(e.E.Key)
}

// Common keycode constants re-exported under the sdl3 package using
// the KeyCode type defined in consts.go. Keep this small and focused
// on commonly used keys; add more as needed.
const (
	K_UNKNOWN   KeyCode = KeyCode(sdl.K_UNKNOWN)
	K_RETURN    KeyCode = KeyCode(sdl.K_RETURN)
	K_ESCAPE    KeyCode = KeyCode(sdl.K_ESCAPE)
	K_BACKSPACE KeyCode = KeyCode(sdl.K_BACKSPACE)
	K_TAB       KeyCode = KeyCode(sdl.K_TAB)
	K_SPACE     KeyCode = KeyCode(sdl.K_SPACE)

	K_LEFT  KeyCode = KeyCode(sdl.K_LEFT)
	K_RIGHT KeyCode = KeyCode(sdl.K_RIGHT)
	K_UP    KeyCode = KeyCode(sdl.K_UP)
	K_DOWN  KeyCode = KeyCode(sdl.K_DOWN)

	K_LSHIFT KeyCode = KeyCode(sdl.K_LSHIFT)
	K_RSHIFT KeyCode = KeyCode(sdl.K_RSHIFT)
	K_LCTRL  KeyCode = KeyCode(sdl.K_LCTRL)
	K_RCTRL  KeyCode = KeyCode(sdl.K_RCTRL)
	K_LALT   KeyCode = KeyCode(sdl.K_LALT)
	K_RALT   KeyCode = KeyCode(sdl.K_RALT)

	K_F1  KeyCode = KeyCode(sdl.K_F1)
	K_F2  KeyCode = KeyCode(sdl.K_F2)
	K_F3  KeyCode = KeyCode(sdl.K_F3)
	K_F4  KeyCode = KeyCode(sdl.K_F4)
	K_F5  KeyCode = KeyCode(sdl.K_F5)
	K_F6  KeyCode = KeyCode(sdl.K_F6)
	K_F7  KeyCode = KeyCode(sdl.K_F7)
	K_F8  KeyCode = KeyCode(sdl.K_F8)
	K_F9  KeyCode = KeyCode(sdl.K_F9)
	K_F10 KeyCode = KeyCode(sdl.K_F10)
	K_F11 KeyCode = KeyCode(sdl.K_F11)
	K_F12 KeyCode = KeyCode(sdl.K_F12)
)

// Additional commonly used keycodes
const (
	K_0 KeyCode = KeyCode(sdl.K_0)
	K_1 KeyCode = KeyCode(sdl.K_1)
	K_2 KeyCode = KeyCode(sdl.K_2)
	K_3 KeyCode = KeyCode(sdl.K_3)
	K_4 KeyCode = KeyCode(sdl.K_4)
	K_5 KeyCode = KeyCode(sdl.K_5)
	K_6 KeyCode = KeyCode(sdl.K_6)
	K_7 KeyCode = KeyCode(sdl.K_7)
	K_8 KeyCode = KeyCode(sdl.K_8)
	K_9 KeyCode = KeyCode(sdl.K_9)

	K_a KeyCode = KeyCode(sdl.K_A)
	K_b KeyCode = KeyCode(sdl.K_B)
	K_c KeyCode = KeyCode(sdl.K_C)
	K_d KeyCode = KeyCode(sdl.K_D)
	K_e KeyCode = KeyCode(sdl.K_E)
	K_f KeyCode = KeyCode(sdl.K_F)
	K_g KeyCode = KeyCode(sdl.K_G)
	K_h KeyCode = KeyCode(sdl.K_H)
	K_i KeyCode = KeyCode(sdl.K_I)
	K_j KeyCode = KeyCode(sdl.K_J)
	K_k KeyCode = KeyCode(sdl.K_K)
	K_l KeyCode = KeyCode(sdl.K_L)
	K_m KeyCode = KeyCode(sdl.K_M)
	K_n KeyCode = KeyCode(sdl.K_N)
	K_o KeyCode = KeyCode(sdl.K_O)
	K_p KeyCode = KeyCode(sdl.K_P)
	K_q KeyCode = KeyCode(sdl.K_Q)
	K_r KeyCode = KeyCode(sdl.K_R)
	K_s KeyCode = KeyCode(sdl.K_S)
	K_t KeyCode = KeyCode(sdl.K_T)
	K_u KeyCode = KeyCode(sdl.K_U)
	K_v KeyCode = KeyCode(sdl.K_V)
	K_w KeyCode = KeyCode(sdl.K_W)
	K_x KeyCode = KeyCode(sdl.K_X)
	K_y KeyCode = KeyCode(sdl.K_Y)
	K_z KeyCode = KeyCode(sdl.K_Z)

	K_MINUS        KeyCode = KeyCode(sdl.K_MINUS)
	K_EQUALS       KeyCode = KeyCode(sdl.K_EQUALS)
	K_LEFTBRACKET  KeyCode = KeyCode(sdl.K_LEFTBRACKET)
	K_RIGHTBRACKET KeyCode = KeyCode(sdl.K_RIGHTBRACKET)
	K_BACKSLASH    KeyCode = KeyCode(sdl.K_BACKSLASH)
	K_SEMICOLON    KeyCode = KeyCode(sdl.K_SEMICOLON)
	K_APOSTROPHE   KeyCode = KeyCode(sdl.K_APOSTROPHE)
	K_GRAVE        KeyCode = KeyCode(sdl.K_GRAVE)
	K_COMMA        KeyCode = KeyCode(sdl.K_COMMA)
	K_PERIOD       KeyCode = KeyCode(sdl.K_PERIOD)
	K_SLASH        KeyCode = KeyCode(sdl.K_SLASH)

	K_INSERT   KeyCode = KeyCode(sdl.K_INSERT)
	K_DELETE   KeyCode = KeyCode(sdl.K_DELETE)
	K_HOME     KeyCode = KeyCode(sdl.K_HOME)
	K_END      KeyCode = KeyCode(sdl.K_END)
	K_PAGEUP   KeyCode = KeyCode(sdl.K_PAGEUP)
	K_PAGEDOWN KeyCode = KeyCode(sdl.K_PAGEDOWN)

	K_CAPSLOCK   KeyCode = KeyCode(sdl.K_CAPSLOCK)
	K_NUMLOCK    KeyCode = KeyCode(sdl.K_NUMLOCKCLEAR)
	K_SCROLLLOCK KeyCode = KeyCode(sdl.K_SCROLLLOCK)

	K_PRINTSCREEN KeyCode = KeyCode(sdl.K_PRINTSCREEN)
	K_PAUSE       KeyCode = KeyCode(sdl.K_PAUSE)
	K_APPLICATION KeyCode = KeyCode(sdl.K_APPLICATION)

	K_KP_0 KeyCode = KeyCode(sdl.K_KP_0)
	K_KP_1 KeyCode = KeyCode(sdl.K_KP_1)
	K_KP_2 KeyCode = KeyCode(sdl.K_KP_2)
	K_KP_3 KeyCode = KeyCode(sdl.K_KP_3)
	K_KP_4 KeyCode = KeyCode(sdl.K_KP_4)
	K_KP_5 KeyCode = KeyCode(sdl.K_KP_5)
	K_KP_6 KeyCode = KeyCode(sdl.K_KP_6)
	K_KP_7 KeyCode = KeyCode(sdl.K_KP_7)
	K_KP_8 KeyCode = KeyCode(sdl.K_KP_8)
	K_KP_9 KeyCode = KeyCode(sdl.K_KP_9)

	K_KP_ENTER    KeyCode = KeyCode(sdl.K_KP_ENTER)
	K_KP_PLUS     KeyCode = KeyCode(sdl.K_KP_PLUS)
	K_KP_MINUS    KeyCode = KeyCode(sdl.K_KP_MINUS)
	K_KP_MULTIPLY KeyCode = KeyCode(sdl.K_KP_MULTIPLY)
	K_KP_DIVIDE   KeyCode = KeyCode(sdl.K_KP_DIVIDE)
	K_KP_PERIOD   KeyCode = KeyCode(sdl.K_KP_PERIOD)
)
