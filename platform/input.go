package platform

// Key represents a keyboard key.
type Key int

const (
	KeyUnknown Key = iota
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeySpace
	KeyEnter
	KeyEscape
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyShift
	KeyCtrl
	KeyAlt
)

// MouseButton represents a mouse button.
type MouseButton int

const (
	MouseLeft MouseButton = iota
	MouseMiddle
	MouseRight
)

// InputState tracks the current state of keyboard and mouse.
// In the real kaiju, this is much richer: it tracks per-frame transitions
// (JustPressed, JustReleased vs IsDown), controller state, touch, stylus, etc.
type InputState struct {
	// KeysDown tracks which keys are currently held down.
	KeysDown map[Key]bool

	// MouseButtons tracks which mouse buttons are currently held down.
	MouseButtons map[MouseButton]bool

	// MouseX, MouseY are the current mouse cursor position.
	MouseX, MouseY int

	// MouseDeltaX, MouseDeltaY are the mouse movement since last frame.
	MouseDeltaX, MouseDeltaY int
}

// NewInputState creates a new InputState with empty maps.
// YOU fill in the implementation.
func NewInputState() *InputState {
	// TODO: YOUR CODE HERE
	return nil
}

// IsKeyDown returns true if the key is currently held.
func (s *InputState) IsKeyDown(key Key) bool {
	// TODO: YOUR CODE HERE
	return false
}

// IsMouseDown returns true if the mouse button is currently held.
func (s *InputState) IsMouseDown(btn MouseButton) bool {
	// TODO: YOUR CODE HERE
	return false
}
