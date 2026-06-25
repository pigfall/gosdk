package sdl3

import (
	"github.com/Zyko0/go-sdl3/sdl"
)

type WindowOption sdl.WindowFlags
type OpenGLProfile int32
type EventType uint32
type KeyCode sdl.Keycode

const (
	WindowOptionOpenGL WindowOption = WindowOption(sdl.WINDOW_OPENGL)
)

const (
	OpenGLCoreProfile          OpenGLProfile = sdl.GL_CONTEXT_PROFILE_CORE
	OpenGLCompatibilityProfile OpenGLProfile = sdl.GL_CONTEXT_PROFILE_COMPATIBILITY
	OpenGLESProfile            OpenGLProfile = sdl.GL_CONTEXT_PROFILE_ES
)

const (
	EventFirst                  EventType = EventType(sdl.EVENT_FIRST)
	EventQuit                   EventType = EventType(sdl.EVENT_QUIT)
	EventUser                   EventType = EventType(sdl.EVENT_USER)
	EventKeyDown                EventType = EventType(sdl.EVENT_KEY_DOWN)
	EventKeyUp                  EventType = EventType(sdl.EVENT_KEY_UP)
	EventTextEditing            EventType = EventType(sdl.EVENT_TEXT_EDITING)
	EventTextInput              EventType = EventType(sdl.EVENT_TEXT_INPUT)
	EventMouseMotion            EventType = EventType(sdl.EVENT_MOUSE_MOTION)
	EventMouseButtonDown        EventType = EventType(sdl.EVENT_MOUSE_BUTTON_DOWN)
	EventMouseButtonUp          EventType = EventType(sdl.EVENT_MOUSE_BUTTON_UP)
	EventMouseWheel             EventType = EventType(sdl.EVENT_MOUSE_WHEEL)
	EventJoystickAxisMotion     EventType = EventType(sdl.EVENT_JOYSTICK_AXIS_MOTION)
	EventJoystickBallMotion     EventType = EventType(sdl.EVENT_JOYSTICK_BALL_MOTION)
	EventJoystickHatMotion      EventType = EventType(sdl.EVENT_JOYSTICK_HAT_MOTION)
	EventJoystickButtonDown     EventType = EventType(sdl.EVENT_JOYSTICK_BUTTON_DOWN)
	EventJoystickButtonUp       EventType = EventType(sdl.EVENT_JOYSTICK_BUTTON_UP)
	EventJoystickAdded          EventType = EventType(sdl.EVENT_JOYSTICK_ADDED)
	EventJoystickRemoved        EventType = EventType(sdl.EVENT_JOYSTICK_REMOVED)
	EventJoystickBatteryUpdated EventType = EventType(sdl.EVENT_JOYSTICK_BATTERY_UPDATED)
	EventGamepadAxisMotion      EventType = EventType(sdl.EVENT_GAMEPAD_AXIS_MOTION)
	EventGamepadButtonDown      EventType = EventType(sdl.EVENT_GAMEPAD_BUTTON_DOWN)
	EventGamepadButtonUp        EventType = EventType(sdl.EVENT_GAMEPAD_BUTTON_UP)
	EventGamepadAdded           EventType = EventType(sdl.EVENT_GAMEPAD_ADDED)
	EventGamepadRemoved         EventType = EventType(sdl.EVENT_GAMEPAD_REMOVED)
	EventClipboardUpdate        EventType = EventType(sdl.EVENT_CLIPBOARD_UPDATE)
	EventSensorUpdate           EventType = EventType(sdl.EVENT_SENSOR_UPDATE)
	EventLast                   EventType = EventType(sdl.EVENT_LAST)
)
