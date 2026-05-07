package engine

// UpdateCallback is a function that runs each frame. It receives the current frame's
// delta time (seconds since last frame).
type UpdateCallback func(deltaTime float64)

// Updater manages per-frame update callbacks. This is how game logic runs each frame.
//
// The key design pattern in kaiju's Updater is DOUBLE BUFFERING:
//   - Callbacks can be added or removed DURING an update cycle without
//     invalidating the iteration
//   - New additions go into a "back buffer" and are merged after the current update
//   - Removals are deferred to after the current update
//
// Without double buffering, adding/removing during iteration would cause bugs
// (slice modification while ranging).
type Updater struct {
	callbacks    []UpdateCallback
	backAdd    []UpdateCallback // deferred additions - merged after Update()
	backRemove []UpdateCallback // deferred removals - applied after Update()
	updating     bool           // true while Update() is running
}

// NewUpdater creates an empty Updater.
func NewUpdater() *Updater {
	return &Updater{}
}

// Add registers a callback to run each frame.
// Safe to call during Update() - the callback is queued for next frame.
func (u *Updater) Add(cb UpdateCallback) {
	// TODO: YOUR CODE HERE
	// If currently updating, add to backAdd (deferred).
	// Otherwise, add to callbacks directly.

}

// Remove unregisters a callback.
// Safe to call during Update() - removal is queued for after the current frame.
func (u *Updater) Remove(cb UpdateCallback) {
	// TODO: YOUR CODE HERE
	// If currently updating, add to backRemove (deferred).
	// Otherwise, remove from callbacks directly.
	// Note: you need to compare function pointers (hint below)

	// HINT: In Go, you can compare function pointers using:
	//   reflect.ValueOf(cb).Pointer() == reflect.ValueOf(existing).Pointer()
	// You'll need to import "reflect".
	// Alternatively, use an ID-based system.
}

// Update runs all callbacks with the given deltaTime.
//
// The update process:
//  1. Set updating = true
//  2. Run all callbacks (each receives deltaTime)
//  3. Set updating = false
//  4. Merge backAdd into callbacks
//  5. Remove backRemove from callbacks
//  6. Clear backAdd and backRemove
func (u *Updater) Update(deltaTime float64) {
	// TODO: YOUR CODE HERE
	// 1. Mark updating
	// 2. Execute all callbacks
	// 3. Unmark updating
	// 4. Apply deferred additions
	// 5. Apply deferred removals
	// 6. Clear back buffers
}

// Count returns the number of active callbacks.
func (u *Updater) Count() int {
	return len(u.callbacks)
}
