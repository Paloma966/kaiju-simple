package platform

import "time"

// Ticker provides high-resolution frame timing.
// It tracks delta time (time since last tick) and total elapsed time.
//
// In the real kaiju, platform/chrono uses platform-specific high-resolution timers.
// For simplicity, we use Go's time package.
type Ticker struct {
	lastTime time.Time
	startTime time.Time
	deltaTime float64 // seconds since last Tick()
	totalTime float64 // seconds since creation
}

// NewTicker creates a new Ticker initialized to the current time.
func NewTicker() *Ticker {
	now := time.Now()
	return &Ticker{
		lastTime:  now,
		startTime: now,
	}
}

// Tick advances the timer. Call this once per frame at the BEGINNING of the frame.
// Computes delta = time since last Tick().
func (t *Ticker) Tick() {
	// TODO: YOUR CODE HERE
	// 1. Get current time (time.Now())
	// 2. Compute delta = now - lastTime, in seconds (use time.Since)
	// 3. Cap delta at a maximum (e.g., 0.1 seconds) to avoid "spiral of death"
	//    when debugging or when the game freezes momentarily
	// 4. Update t.deltaTime, t.totalTime, t.lastTime

}

// Delta returns the delta time of the most recent Tick() in seconds.
func (t *Ticker) Delta() float64 {
	return t.deltaTime
}

// Total returns the total elapsed time since the Ticker was created, in seconds.
func (t *Ticker) Total() float64 {
	return t.totalTime
}
