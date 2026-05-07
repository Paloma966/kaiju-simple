package matrix

// Color 表示一个 RGBA 颜色，各分量用 float32 存储，范围为 [0, 1]。
// 与着色器表示颜色的方式一致。
type Color struct {
	R, G, B, A float32
}

// 预定义颜色。
var (
	ColorWhite  = Color{1, 1, 1, 1}
	ColorBlack  = Color{0, 0, 0, 1}
	ColorRed    = Color{1, 0, 0, 1}
	ColorGreen  = Color{0, 1, 0, 1}
	ColorBlue   = Color{0, 0, 1, 1}
	ColorYellow = Color{1, 1, 0, 1}
	ColorCyan   = Color{0, 1, 1, 1}
	ColorMagenta = Color{1, 0, 1, 1}
)

// NewColor 创建一个 Color。所有值会被钳制到 [0, 1]。
func NewColor(r, g, b, a float32) Color {
	// TODO: YOUR CODE HERE - clamp each component to [0, 1]
	return Color{r, g, b, a}
}

// ToVec4 将颜色转换为 Vec4（对应关系：R=X, G=Y, B=Z, A=W）。
func (c Color) ToVec4() Vec4 {
	// TODO: YOUR CODE HERE
	return Vec4{}
}

// Scale 将 RGB 通道乘以 s（Alpha 不变）。结果钳制到 [0, 1]。
func (c Color) Scale(s float32) Color {
	// TODO: YOUR CODE HERE
	return c
}
