package rendering

// Shader represents a GPU shader program (vertex + fragment shader compiled and linked).
//
// In the real kaiju, shaders are:
//   - Written in GLSL
//   - Pre-compiled to SPIR-V during the build process (via generators/spirv/)
//   - Loaded and cached by the shader system
//   - Support live reloading in the editor (re-compile GLSL → SPIR-V on file change)
//   - Organized into pipeline layouts with descriptor set layouts
type Shader interface {
	// Bind activates this shader for subsequent draw calls.
	Bind()

	// Unbind deactivates this shader.
	Unbind()

	// Destroy releases GPU shader resources.
	Destroy()
}
