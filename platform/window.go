package platform

import "kaiju-simple/rendering"

// Window is the platform window abstraction.
// In the real kaiju, this is implemented per-platform:
//   - X11/XCB on Linux (platform/windowing/window_x11.go)
//   - Win32 on Windows
//   - Cocoa on macOS
//   - Android Native Activity on Android
//
// For the simplified version, YOU will create a basic implementation.
type Window interface {
	// ShouldClose returns true when the user has requested to close the window
	// (e.g., clicked the X button, pressed Alt+F4).
	ShouldClose() bool

	// Renderer returns the rendering device for this window.
	// In the real engine, this creates a Vulkan surface on the window.
	Renderer() rendering.Device

	// Width and Height return the current window dimensions in pixels.
	Width() int
	Height() int

	// Destroy cleans up window resources.
	Destroy()
}

// NewWindow creates a new window with the given title and dimensions.
//
// YOU implement this. For the skeleton, you have a few options:
//   A) Return a "null window" that always reports ShouldClose() = false (infinite loop)
//   B) Use a Go library like ebiten, pixel, or SDL2 bindings
//   C) Create a simple terminal-based "window"
func NewWindow(title string, width, height int) Window {
	// TODO: YOUR CODE HERE
	// Return a window implementation.
	// For now, you can return nil and implement this later.
	return nil
}
