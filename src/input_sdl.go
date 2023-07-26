package main

import (
	sdl "github.com/veandco/go-sdl2/sdl"
)

type Input struct {
	joystick []sdl.Joystick
}

type Key = sdl.Keycode
type ModifierKey = int

const (
	KeyUnknown = sdl.K_UNKNOWN
	KeyEscape  = sdl.K_ESCAPE
	KeyEnter   = sdl.K_RETURN
	KeyInsert  = sdl.K_INSERT
	KeyF12     = sdl.K_F12
)

func init() {
}

func StringToKey(s string) sdl.Keycode {

	return sdl.GetKeyFromName(s)
}

func KeyToString(k sdl.Keycode) string {
	return sdl.GetKeyName(k)
}

func KeyToInt(k sdl.Keycode) int {
	return int(sdl.GetScancodeFromKey(k))
}

func IntToKey(scancode int) sdl.Keycode {
	return sdl.GetKeyFromScancode(sdl.Scancode(scancode))
}

func NewModifierKey(ctrl, alt, shift bool) (mod ModifierKey) {
	if ctrl {
		mod |= sdl.KMOD_CTRL
	}
	if alt {
		mod |= sdl.KMOD_ALT
	}
	if shift {
		mod |= sdl.KMOD_SHIFT
	}
	return
}

// This needs to be completely reworked. Not needed for Lembas

/*
var input = Input{
	joystick: []glfw.Joystick{glfw.Joystick1, glfw.Joystick2, glfw.Joystick3,
		glfw.Joystick4, glfw.Joystick5, glfw.Joystick6, glfw.Joystick7,
		glfw.Joystick8, glfw.Joystick9, glfw.Joystick10, glfw.Joystick11,
		glfw.Joystick12, glfw.Joystick13, glfw.Joystick14, glfw.Joystick15,
		glfw.Joystick16},
}
*/

var input = Input{}

func (input *Input) GetMaxJoystickCount() int {
	return len(input.joystick)
}

func (input *Input) IsJoystickPresent(joy int) bool {
	/*
		if joy < 0 || joy >= len(input.joystick) {
			return false
		}
		return input.joystick[joy].IsPresent()*/
	return false
}

func (input *Input) GetJoystickName(joy int) string {
	/*
		if joy < 0 || joy >= len(input.joystick) {
			return ""
		}
		return input.joystick[joy].GetGamepadName()
	*/
	return ""
}

func (input *Input) GetJoystickAxes(joy int) []float32 {
	/*
		if joy < 0 || joy >= len(input.joystick) {
			return []float32{}
		}
		return input.joystick[joy].GetAxes()
	*/
	return []float32{}
}

/*
func (input *Input) GetJoystickButtons(joy int) []glfw.Action {
	if joy < 0 || joy >= len(input.joystick) {
		return []glfw.Action{}
	}
	return input.joystick[joy].GetButtons()
}
*/
