// Package build provides build-time constants, similar to kaiju's src/build/.
// In the real engine, these are set via Go build tags (-tags "editor,debug").
package build

// GameTitle identifies the type of build (game, editor, generator).
type GameTitle int

const (
	GameTitleRawGame GameTitle = iota
	GameTitleEditor
)

// These would be controlled by build tags in the real engine.
// For the simplified version, we hardcode them.
const (
	Title    = GameTitleRawGame
	IsEditor = false
	IsDebug  = true
)
