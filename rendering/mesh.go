package rendering

// VertexAttribute describes a single field in a vertex (position, normal, UV, etc.).
type VertexAttribute struct {
	Name     string // e.g., "position", "normal", "uv"
	Format   VertexFormat
	Location int // shader input location (layout(location = X) in GLSL)
}

// VertexFormat describes the data type of a vertex attribute.
type VertexFormat int

const (
	VertexFormatVec2 VertexFormat = iota // 2 floats
	VertexFormatVec3                     // 3 floats
	VertexFormatVec4                     // 4 floats
)

// Size returns the number of float32 components in this format.
func (f VertexFormat) Size() int {
	switch f {
	case VertexFormatVec2:
		return 2
	case VertexFormatVec3:
		return 3
	case VertexFormatVec4:
		return 4
	default:
		return 0
	}
}

// VertexLayout describes how vertex data is laid out in memory.
// It defines the attributes and the stride (bytes between consecutive vertices).
type VertexLayout struct {
	Attributes []VertexAttribute
}

// Stride returns the total number of float32 components per vertex.
//
// Example: a vertex with position (vec3) + normal (vec3) + uv (vec2) has stride = 8
func (l VertexLayout) Stride() int {
	// TODO: YOUR CODE HERE
	// Sum the Size() of each attribute
	return 0
}

// Mesh represents a GPU-resident mesh (vertex buffer + index buffer).
// In the real engine, this is a Vulkan buffer.
type Mesh interface {
	// VertexCount returns the number of vertices.
	VertexCount() int

	// IndexCount returns the number of indices (0 if not indexed).
	IndexCount() int

	// IsIndexed returns true if the mesh uses an index buffer.
	IsIndexed() bool

	// Destroy releases the GPU buffers.
	Destroy()
}
