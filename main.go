package main

import (
	"fmt"
	"kaiju-simple/assets"
	"kaiju-simple/bootstrap"
	"kaiju-simple/engine"
	"kaiju-simple/matrix"
)

// -------- Your Game --------

// MyGame implements engine.GameInterface.
// This is YOUR game — fill in Launch() to create entities, set up the scene, etc.
type MyGame struct{}

// ContentDatabase returns an asset database.
// For now, we return nil — implement your own assets.Database to load real assets.
func (g *MyGame) ContentDatabase() (assets.Database, error) {
	// TODO: YOUR CODE HERE
	// Return a database that can load assets from disk
	// For now, nil is fine for the skeleton
	return nil, nil
}

// Launch is called after the engine initializes.
// This is where you create entities, set up the scene, load assets, etc.
//
// Example of what you might do here:
//   host.Updater.Add(func(dt float64) {
//       // rotate an entity each frame
//   })
func (g *MyGame) Launch(host *engine.Host) {
	// TODO: YOUR CODE HERE
	// 1. Create an entity
	// 2. Position it in the scene
	// 3. Register update callbacks for game logic
	// 4. Register update callbacks for the input system

	fmt.Println("Game launched! Fill in Launch() to create your scene.")
	fmt.Printf("Host initialized: window=%v, renderer=%v\n", host.Window, host.Renderer)

	// Example entity creation (uncomment when you've filled in matrix.NewVec3):
	// player := engine.NewEntity("Player")
	// player.Transform.Position = matrix.NewVec3(0, 0, -5)
	// fmt.Printf("Created entity: %s at %v\n", player.Name, player.Transform.Position)

	// Example update callback (uncomment when Updater is implemented):
	// host.Updater.Add(func(dt float64) {
	//     // game logic here - runs every frame
	// })
}

// -------- Entry Point --------

func main() {
	fmt.Println("Kaiju Simple Engine")
	fmt.Println("===================")

	// Show that our math library works (fill in matrix functions first!)
	v := matrix.NewVec2(3, 4)
	fmt.Printf("Vec2 test: %v, length=%v\n", v, v.Length())

	game := &MyGame{}
	bootstrap.Main(game)
}
