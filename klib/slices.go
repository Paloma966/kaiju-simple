package klib

// The klib package is the "Kaiju Standard Library" — general-purpose utilities
// used throughout the engine. In the real kaiju, klib contains: slice ops,
// string ops, memory management, maps, sets, stacks, serialization, etc.

// ------- Slice Operations -------

// RemoveUnordered removes the element at index i from the slice by swapping
// it with the last element and truncating. This is O(1) but does NOT preserve order.
// Use this when element order doesn't matter (e.g., render lists).
func RemoveUnordered[T any](s []T, i int) []T {
	// TODO: YOUR CODE HERE
	// 1. Swap s[i] with s[len(s)-1]
	// 2. Return s[:len(s)-1]
	return s
}

// RemoveOrdered removes the element at index i while preserving order.
// This is O(n) because all elements after i must shift left.
func RemoveOrdered[T any](s []T, i int) []T {
	// TODO: YOUR CODE HERE
	// 1. Shift all elements after i one position left
	// 2. Return s[:len(s)-1]
	return s
}

// FindIndex returns the first index where pred returns true, or -1 if not found.
// This is a common functional pattern used throughout the engine.
func FindIndex[T any](s []T, pred func(T) bool) int {
	// TODO: YOUR CODE HERE
	// Loop through s, return index of first element matching pred
	return -1
}

// Filter returns a new slice containing only elements where pred returns true.
func Filter[T any](s []T, pred func(T) bool) []T {
	// TODO: YOUR CODE HERE
	return nil
}

// Contains returns true if the slice contains the given element.
// Uses == comparison, so T must be comparable.
func Contains[T comparable](s []T, elem T) bool {
	// TODO: YOUR CODE HERE
	return false
}
