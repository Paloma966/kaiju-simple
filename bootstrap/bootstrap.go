package bootstrap

import (
	"fmt"
	"kaiju-simple/engine"
)

// Main is the engine entry point. It creates a Container for the game,
// runs the startup sequence, and enters the main loop.
//
// This mirrors kaiju's bootstrap.Main() function, which:
//  1. Creates a host_container.Container
//  2. Calls container.Run() to start the engine and enter the game loop
//  3. Handles any errors from Run()
//
// Unlike the real kaiju, we skip:
//   - Logging initialization
//   - Profiling/tracing setup
//   - CLI argument parsing (launch params)
//   - Crash handling
//   - Steam integration
func Main(game engine.GameInterface) {
	// TODO: YOUR CODE HERE
	// 1. Create a Container: engine.NewContainer(game)
	// 2. Call container.Run()
	// 3. If Run() returns an error, print it with fmt.Println
	// 4. If Run() returns nil, print "Goodbye!"

	fmt.Println("bootstrap.Main() - fill in this function")
}
