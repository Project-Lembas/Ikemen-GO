package main

import (
	"fmt"
	"image"

	sdl "github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	*sdl.Window
	title      string
	fullscreen bool
	x, y, w, h int
}

func (s *System) newWindow(w, h int) (*Window, error) {
	var err error
	var window *sdl.Window
	//var monitor *sdl.Monitor

	// Initialize OpenGL
	chk(sdl.Init(sdl.INIT_EVERYTHING))

	//if monitor = glfw.GetPrimaryMonitor(); monitor == nil {
	//	return nil, fmt.Errorf("failed to obtain primary monitor")
	//}

	var mode, _ = window.GetDisplayMode()
	var x, y = (int(mode.W) - w) / 2, (int(mode.H) - h) / 2

	// "-windowed" overrides the configuration setting but does not change it
	_, forceWindowed := sys.cmdFlags["-windowed"]
	fullscreen := s.fullscreen && !forceWindowed

	//glfw.WindowHint(glfw.Resizable, glfw.False)
	//glfw.WindowHint(glfw.ContextVersionMajor, 2)
	//glfw.WindowHint(glfw.ContextVersionMinor, 1)

	// Create main window.
	// NOTE: Borderless fullscreen is in reality just a window without borders.
	if fullscreen && !s.borderless {
		window, err = sdl.CreateWindow(s.windowTitle, 0, 0, int32(w), int32(y), sdl.WINDOW_RESIZABLE|sdl.WINDOW_SHOWN)
		//window, err = glfw.CreateWindow(w, h, s.windowTitle, monitor, nil)
	} else {
		window, err = sdl.CreateWindow(s.windowTitle, 0, 0, int32(w), int32(y), sdl.WINDOW_RESIZABLE|sdl.WINDOW_SHOWN)
		//window, err = glfw.CreateWindow(w, h, s.windowTitle, nil, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %w", err)
	}

	// Set windows attributes
	if fullscreen {
		window.SetPosition(0, 0)
		if s.borderless {
			window.SetBordered(false)
			window.SetSize(mode.W, mode.H)
		}
		//TODO: hide cursor
		//window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	} else {
		window.SetSize(int32(w), int32(h))
		//TODO: unhide cursor
		//window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
		if s.windowCentered {
			window.SetPosition(int32(x), int32(y))
		}
	}

	//TODO: Figure out the callbacks here
	//window.MakeContextCurrent()
	//window.SetKeyCallback(keyCallback)
	//window.SetCharModsCallback(charCallback)

	//TODO: Does this even matter without gl?
	// V-Sync
	//if s.vRetrace >= 0 {
	//	glfw.SwapInterval(s.vRetrace)
	//}

	window.Show()
	ret := &Window{window, s.windowTitle, fullscreen, x, y, w, h}
	return ret, err
}

func (w *Window) SwapBuffers() {
	//w.Window.SwapBuffers()
}

func (w *Window) SetIcon(icon []image.Image) {
	//w.Window.SetIcon(icon)
}

func (w *Window) SetSwapInterval(interval int) {
	//glfw.SwapInterval(interval)
}

func (w *Window) GetSize() (int, int) {
	var wid, h = w.Window.GetSize()
	return int(wid), int(h)
}

func (w *Window) GetClipboardString() (string, error) {
	//return w.Window.GetClipboardString()
	return "", nil
}

func (w *Window) toggleFullscreen() {
	/*
		var mode = glfw.GetPrimaryMonitor().GetVideoMode()

		if w.fullscreen {
			w.SetAttrib(glfw.Decorated, 1)
			w.SetMonitor(&glfw.Monitor{}, w.x, w.y, w.w, w.h, mode.RefreshRate)
			w.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
		} else {
			w.SetAttrib(glfw.Decorated, 0)
			if sys.borderless {
				w.SetSize(mode.Width, mode.Height)
				w.SetMonitor(&glfw.Monitor{}, 0, 0, mode.Width, mode.Height, mode.RefreshRate)
			} else {
				w.x, w.y = w.GetPos()
				w.SetMonitor(glfw.GetPrimaryMonitor(), w.x, w.y, w.w, w.h, mode.RefreshRate)
			}
			w.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
		}
		if sys.vRetrace != -1 {
			glfw.SwapInterval(sys.vRetrace)
		}
		w.fullscreen = !w.fullscreen
	*/
}

func (w *Window) pollEvents() {
	//glfw.PollEvents()
}

func (w *Window) shouldClose() bool {
	//return w.Window.ShouldClose()
	return false
}

func (w *Window) Close() {
	sdl.Quit()
}

/*
func keyCallback(_ *glfw.Window, key Key, _ int, action glfw.Action, mk ModifierKey) {

	switch action {
	case glfw.Release:
		OnKeyReleased(key, mk)
	case glfw.Press:
		OnKeyPressed(key, mk)
	}

}

func charCallback(_ *glfw.Window, char rune, mk ModifierKey) {
	OnTextEntered(string(char))
}
*/
