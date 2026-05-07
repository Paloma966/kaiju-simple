package engine

import "kaiju-simple/matrix"

// Entity is the core game object. Every object in the game world is an Entity.
// Entities form a tree hierarchy (parent/children) through which transforms propagate.
//
// In the real kaiju engine, Entity also has:
//   - Named data slots for attaching arbitrary data
//   - Lifecycle events (OnActivate, OnDeactivate, OnDestroy)
//   - Integration with the collision/physics systems
type Entity struct {
	Name     string
	Transform matrix.Transform

	parent   *Entity
	children []*Entity

	// Enabled controls whether this entity participates in Update and Render.
	Enabled bool
}

// NewEntity creates a new entity with the given name at the origin.
func NewEntity(name string) *Entity {
	return &Entity{
		Name:     name,
		Transform: matrix.NewTransform(matrix.Vec3{}),
		Enabled:  true,
	}
}

// ------- Parent/Child Hierarchy -------
//
// The transform hierarchy works like this:
//   - An entity's Transform is in LOCAL space (relative to parent)
//   - WorldTransform = parent's WorldTransform * local Transform
//   - Changing a parent's transform affects all descendants

// SetParent removes this entity from its current parent and attaches it to newParent.
// Pass nil to detach from parent (make it a root entity).
//
// IMPORTANT: You must:
//  1. Remove this entity from old parent's children slice
//  2. Add this entity to new parent's children slice
//  3. Set the parent field
func (e *Entity) SetParent(newParent *Entity) {
	// TODO: YOUR CODE HERE
	// 1. If e already has a parent, remove e from the old parent's children list
	// 2. Set e.parent = newParent
	// 3. If newParent is not nil, add e to newParent's children list
	//    (avoid duplicates!)

}

// AddChild adds child as a child of this entity.
// This is a convenience method: it calls child.SetParent(e).
func (e *Entity) AddChild(child *Entity) {
	// TODO: YOUR CODE HERE

}

// RemoveChild removes child from this entity's children.
// This is a convenience method: it calls child.SetParent(nil).
func (e *Entity) RemoveChild(child *Entity) {
	// TODO: YOUR CODE HERE

}

// Parent returns this entity's parent (nil for root entities).
func (e *Entity) Parent() *Entity {
	return e.parent
}

// Children returns a copy of this entity's children list.
// Returning a copy prevents external modification of the internal slice.
func (e *Entity) Children() []*Entity {
	// TODO: YOUR CODE HERE
	// Return a copy so callers can't modify the internal slice
	return nil
}

// ChildCount returns the number of direct children.
func (e *Entity) ChildCount() int {
	return len(e.children)
}

// ------- World Transform -------

// WorldTransform computes the entity's transform in world space by walking up
// the parent chain and multiplying matrices.
//
//   worldMatrix = parent.WorldTransform() * local.Matrix()
//
// For a root entity (no parent), world transform = local transform matrix.
func (e *Entity) WorldTransform() matrix.Mat4 {
	// TODO: YOUR CODE HERE
	// 1. Get local matrix: e.Transform.Matrix()
	// 2. If has parent, multiply parent.WorldTransform() * local
	// 3. If no parent, return local
	return matrix.Identity()
}

// WorldPosition returns the entity's position in world space.
// Hint: extract the translation column from the world transform matrix.
func (e *Entity) WorldPosition() matrix.Vec3 {
	// TODO: YOUR CODE HERE
	// The translation is stored in the last column of a Mat4:
	// indices [12], [13], [14]
	return matrix.Vec3{}
}
