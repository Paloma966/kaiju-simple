package engine

import (
	"kaiju-simple/assets"
	"kaiju-simple/platform"
	"kaiju-simple/rendering"
)

// Host is the central mediator of the engine — the "god object" that provides
// access to all engine subsystems. It is passed to every game via GameInterface.Launch().
//
// In the real kaiju, Host also implements context.Context (for cancellation)
// and manages: cameras, input, physics, collision, work groups, caches, events, etc.
type Host struct {
	Window   platform.Window
	Renderer rendering.Device
	Updater  *Updater
	Database assets.Database

	// Frame timing
	DeltaTime   float64 // seconds since last frame
	TotalTime   float64 // seconds since engine started
	FrameCount  uint64  // total frames rendered
}

// NewHost creates a Host. Systems are not yet initialized — call Initialize() for that.
func NewHost() *Host {
	return &Host{
		Updater: NewUpdater(),
	}
}

// Initialize sets up the engine subsystems.
//
// YOU fill in the implementation. The steps are:
//  1. Create a window (platform.NewWindow)
//  2. Get the renderer from the window
//  3. Call renderer.Initialize()
//  4. Store window and renderer in the Host
func (h *Host) Initialize() error {
	// TODO: YOUR CODE HERE
	// 1. h.Window = platform.NewWindow("Kaiju Simple", 1280, 720)
	// 2. h.Renderer = h.Window.Renderer()
	// 3. return h.Renderer.Initialize()
	return nil
}

// Update runs one frame of game logic. Called by the game loop in host_container.
func (h *Host) Update() {
	// TODO: YOUR CODE HERE
	// Run all registered updaters with the current delta time
}

// Render draws one frame. Called by the game loop after Update.
func (h *Host) Render() {
	// TODO: YOUR CODE HERE
	// Tell the renderer to begin a frame, render, then end the frame
	// h.Renderer.BeginFrame()
	// h.Renderer.EndFrame()
}

// Shutdown cleans up engine resources.
func (h *Host) Shutdown() {
	// TODO: YOUR CODE HERE
	// Destroy renderer and window
}

// ShouldClose returns true when the window has been requested to close.
func (h *Host) ShouldClose() bool {
	if h.Window == nil {
		return false
	}
	return h.Window.ShouldClose()
}
