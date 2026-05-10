package matrix

import (
	"fmt"
	"math"
)

// Mat4 是 4x4 列主序矩阵，以 [16]float32 扁平数组存储。
//
// 列主序布局（索引）：
//
//	[0]  [4]  [8]  [12]
//	[1]  [5]  [9]  [13]
//	[2]  [6]  [10] [14]
//	[3]  [7]  [11] [15]
//
// 这是 OpenGL/Vulkan 的标准布局。每列代表一个轴：
// 第 0 列 = X 轴，第 1 列 = Y 轴，第 2 列 = Z 轴，第 3 列 = 平移。
type Mat4 [16]float32

// ------- 构造函数 -------

// Identity 返回 4x4 单位矩阵（对角线为 1）。
func Identity() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// NewMat4 用 16 个值（列主序）创建 Mat4。
func NewMat4(
	m00, m01, m02, m03,
	m10, m11, m12, m13,
	m20, m21, m22, m23,
	m30, m31, m32, m33 float32,
) Mat4 {
	return Mat4{
		m00, m01, m02, m03,
		m10, m11, m12, m13,
		m20, m21, m22, m23,
		m30, m31, m32, m33,
	}
}

// ------- 矩阵乘法 -------

// Multiply 返回矩阵乘积 m * other。
//
// 对于列主序矩阵，结果的 (row i, col j) 元素为：
//
//	result[i + j*4] = Σ(k=0..3) m[i + k*4] * other[k + j*4]
//
// 重要：矩阵乘法不可交换。一般情况下 m * other ≠ other * m。
func (m Mat4) Multiply(other Mat4) Mat4 {
	var result Mat4
	for i := range 4 {
		for j := range 4 {
			var sum float32
			for k := range 4 {
				sum += m[i+k*4] * other[k+j*4]
			}
			result[i+j*4] = sum
		}
	}
	return result
}

// ------- 变换矩阵 -------

// Translate 返回按 (x, y, z) 平移的矩阵。
//
// 平移量放在最后一列 [12], [13], [14]。
//
//	1  0  0  x
//	0  1  0  y
//	0  0  1  z
//	0  0  0  1
func Translate(x, y, z float32) Mat4 {
	// TODO: YOUR CODE HERE
	return Mat4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}

// Scale 返回缩放矩阵。
//
//	┌          ┐
//	│ x  0  0 0│
//	│ 0  y  0 0│
//	│ 0  0  z 0│
//	│ 0  0  0 1│
//	└          ┘
func Scale(x, y, z float32) Mat4 {
	// TODO: YOUR CODE HERE
	return Mat4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

// RotateX 返回绕 X 轴旋转 angle（弧度）的旋转矩阵。
//
//	┌                       ┐
//	│ 1   0      0     0    │
//	│ 0  cosA  -sinA   0    │
//	│ 0  sinA   cosA   0    │
//	│ 0   0      0     1    │
//	└                       ┘
//
// 提示：使用 math.Sin 和 math.Cos
func RotateX(angle float32) Mat4 {
	// TODO: YOUR CODE HERE
	return Mat4{
		1, 0, 0, 0,
		0, float32(math.Cos(float64(angle))), -float32(math.Sin(float64(angle))), 0,
		0, float32(math.Sin(float64(angle))), float32(math.Cos(float64(angle))), 0,
		0, 0, 0, 1,
	}
}

// RotateY 返回绕 Y 轴旋转 angle（弧度）的旋转矩阵。
//
//	┌                       ┐
//	│  cosA  0   sinA   0   │
//	│   0    1    0     0   │
//	│ -sinA  0   cosA   0   │
//	│   0    0    0     1   │
//	└                       ┘
func RotateY(angle float32) Mat4 {
	return Mat4{
		float32(math.Cos(float64(angle))), 0, float32(math.Sin(float64(angle))), 0,
		0, 1, 0, 0,
		-float32(math.Sin(float64(angle))), 0, float32(math.Cos(float64(angle))), 0,
		0, 0, 0, 1,
	}
}

// RotateZ 返回绕 Z 轴旋转 angle（弧度）的旋转矩阵。
//
//	┌                       ┐
//	│  cosA  -sinA  0   0   │
//	│  sinA   cosA  0   0   │
//	│   0      0    1   0   │
//	│   0      0    0   1   │
//	└                       ┘
func RotateZ(angle float32) Mat4 {
	return Mat4{
		float32(math.Cos(float64(angle))), -float32(math.Sin(float64(angle))), 0, 0,
		float32(math.Sin(float64(angle))), float32(math.Cos(float64(angle))), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// ------- 向量变换 -------

// MultiplyVec3 用该矩阵变换 Vec3，将其视为点（w=1）。
// 用于位置和平移变换。
//
//	result.X = m[0]*v.X + m[4]*v.Y + m[8]*v.Z  + m[12]
//	result.Y = m[1]*v.X + m[5]*v.Y + m[9]*v.Z  + m[13]
//	result.Z = m[2]*v.X + m[6]*v.Y + m[10]*v.Z + m[14]
func (m Mat4) MultiplyVec3(v Vec3) Vec3 {
	// TODO: YOUR CODE HERE
	return Vec3{}
}

// MultiplyVec3Direction 用该矩阵变换 Vec3，将其视为方向（w=0）。
// 平移量被忽略（不加最后一列）。用于法线和方向变换。
//
//	result.X = m[0]*v.X + m[4]*v.Y + m[8]*v.Z
//	result.Y = m[1]*v.X + m[5]*v.Y + m[9]*v.Z
//	result.Z = m[2]*v.X + m[6]*v.Y + m[10]*v.Z
func (m Mat4) MultiplyVec3Direction(v Vec3) Vec3 {
	// TODO: YOUR CODE HERE
	return Vec3{}
}

// ------- 工具方法 -------

// Transpose 返回该矩阵的转置（行变列、列变行）。
func (m Mat4) Transpose() Mat4 {
	// TODO: YOUR CODE HERE
	return Identity()
}

func (m Mat4) String() string {
	return fmt.Sprintf(
		"[%f %f %f %f]\n[%f %f %f %f]\n[%f %f %f %f]\n[%f %f %f %f]",
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	)
}
