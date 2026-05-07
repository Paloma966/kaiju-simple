package rendering

import "kaiju-simple/matrix"

// Drawing represents something that can be rendered — it pairs a Mesh with a Shader
// and a transform (model matrix). The renderer iterates over Drawings each frame.
//
// In the real kaiju, Drawing is much more complex:
//   - It holds a DrawInstance which contains per-instance data for the GPU
//   - It links to Textures/Materials through the material system
//   - It manages descriptor sets (Vulkan) for shader uniforms
//   - It has culling info (bounding sphere) for frustum culling
type Drawing struct {
	Mesh      Mesh
	Shader    Shader
	Transform *matrix.Transform // pointer so it can be shared with the Entity's transform
	Visible   bool
}

// NewDrawing creates a Drawing.
func NewDrawing(mesh Mesh, shader Shader, transform *matrix.Transform) *Drawing {
	return &Drawing{
		Mesh:      mesh,
		Shader:    shader,
		Transform: transform,
		Visible:   true,
	}
}

// ModelMatrix returns the model matrix for this drawing.
// This is used by the renderer to position the object in world space.
func (d *Drawing) ModelMatrix() matrix.Mat4 {
	// TODO: YOUR CODE HERE
	// Return d.Transform.Matrix()
	return matrix.Identity()
}
