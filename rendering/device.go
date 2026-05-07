package rendering

// Device abstracts the GPU/renderer. In the real kaiju, this is a Vulkan logical
// device with swap chain, command buffers, descriptor sets, pipeline layouts, etc.
//
// For the simplified version, this is an interface that YOU implement.
// A "null device" that does nothing is fine for learning the architecture.
type Device interface {
	// Initialize sets up the GPU resources.
	Initialize() error

	// BeginFrame prepares the device for a new frame.
	BeginFrame()

	// EndFrame submits the frame to the display.
	EndFrame()

	// Destroy releases all GPU resources.
	Destroy()

	// CreateMesh uploads vertex/index data to the GPU and returns a Mesh handle.
	//
	// Vertices are interleaved: [posX, posY, posZ, normalX, normalY, normalZ, u, v, ...]
	// The vertex layout defines how to interpret these floats.
	CreateMesh(vertices []float32, indices []uint32, layout VertexLayout) Mesh

	// CreateShader compiles/loads a shader program from source strings.
	// In the real kaiju, shaders are SPIR-V compiled from GLSL.
	CreateShader(vertexSource, fragmentSource string) Shader
}
