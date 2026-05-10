package matrix

import "fmt"

// Vec4 是 4D 向量，常用于 RGBA 颜色和齐次坐标。
type Vec4 struct {
	X, Y, Z, W float32
}

func NewVec4(x, y, z, w float32) Vec4 {
	return Vec4{X: x, Y: y, Z: z, W: w}
}

// ------- 基本操作 -------

func (v Vec4) Add(other Vec4) Vec4 {
	return Vec4{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
		W: v.W + other.W,
	}
}

func (v Vec4) Sub(other Vec4) Vec4 {
	// TODO: YOUR CODE HERE
	return Vec4{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
		W: v.W - other.W,
	}
}

func (v Vec4) Scale(s float32) Vec4 {
	// TODO: YOUR CODE HERE
	return Vec4{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
		W: v.W * s,
	}
}

func (v Vec4) Dot(other Vec4) float32 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

// XYZ 返回前三个分量组成的 Vec3。
// 用于齐次坐标 (x,y,z,w) 中需要提取 xyz 部分的场景。
func (v Vec4) XYZ() Vec3 {
	return Vec3{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}

func (v Vec4) String() string {
	return fmt.Sprintf("Vec4(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}
