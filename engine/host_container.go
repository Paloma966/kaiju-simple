package engine

import (
	"fmt"
	"kaiju-simple/platform"
)

// Container wraps a Host and manages the main game loop.
//
// In the real kaiju, Container is in engine/host_container/ and handles
// window creation, renderer initialization, audio setup, and the frame loop.
type Container struct {
	Host *Host
	Game GameInterface
}

// NewContainer creates a Container for the given game.
func NewContainer(game GameInterface) *Container {
	return &Container{
		Host: NewHost(),
		Game: game,
	}
}

// Run starts the engine and runs the main game loop. This function blocks
// until the window is closed or an error occurs.
//
// The startup sequence (matches the real kaiju exactly):
//  1. Load the content database    → Game.ContentDatabase()
//  2. Initialize engine subsystems → Host.Initialize()
//  3. Launch the game              → Game.Launch(host)
//  4. Enter the main loop          → Update + Render each frame
//  5. Shutdown                     → Host.Shutdown()
func (c *Container) Run() error {
	// ---- Step 1: Load asset database ----
	db, err := c.Game.ContentDatabase()
	if err != nil {
		return fmt.Errorf("failed to load content database: %w", err)
	}
	c.Host.Database = db

	// ---- Step 2: Initialize engine (window, renderer) ----
	if err := c.Host.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize host: %w", err)
	}

	// ---- Step 3: Launch the game ----
	c.Game.Launch(c.Host)

	// ---- Step 4: Main loop ----
	c.runLoop()

	// ---- Step 5: Shutdown ----
	c.Host.Shutdown()
	return nil
}

// runLoop runs the main game loop: Update → Render, frame after frame.
//
// The frame timing system:
//   - chrono.Ticker provides high-resolution time
//   - DeltaTime = time since last frame (capped to avoid spiral of death)
//   - Target frame rate is typically 60 FPS (but we don't enforce a cap)
//
// YOU fill in the loop body. The structure:
//
//	ticker := platform.NewTicker()
//	for !c.Host.ShouldClose() {
//	    ticker.Tick()                    // advance timer
//	    c.Host.DeltaTime = ticker.Delta()
//	    c.Host.TotalTime = ticker.Total()
//
//	    c.Host.Update()                  // game logic
//	    c.Host.Render()                  // draw
//
//	    c.Host.FrameCount++
//	}
func (c *Container) runLoop() {
	// TODO: YOUR CODE HERE
	// 1. Create a platform.Ticker for frame timing
	// 2. Loop until window should close
	// 3. Each iteration: tick → update host timing → update logic → render → count frame

	_ = platform.NewTicker // remove this when you implement runLoop()
	fmt.Println("Game loop would run here - fill in runLoop()")
}
