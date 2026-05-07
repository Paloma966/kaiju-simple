package engine

import "kaiju-simple/assets"

// GameInterface is the contract every game (or editor) must implement.
// This is the same pattern as kaiju's bootstrap.GameInterface.
//
// The engine calls these methods during startup:
//  1. ContentDatabase() - load assets
//  2. Launch() - initialize your game
//  3. (game loop runs Update + Render each frame)
type GameInterface interface {
	// Launch is called once after the engine initializes. Use the Host to
	// create entities, load assets, set up cameras, etc.
	Launch(host *Host)

	// ContentDatabase returns the asset database for this game.
	// Assets include textures, meshes, shaders, sounds, etc.
	ContentDatabase() (assets.Database, error)
}
