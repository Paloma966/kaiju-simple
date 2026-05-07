package matrix

import (
	"fmt"
	"math"
)

// Vec2 是具有 X 和 Y 分量的 2D 向量。
type Vec2 struct {
	X, Y float32
}

// NewVec2 创建一个新的 Vec2。
func NewVec2(x, y float32) Vec2 {
	return Vec2{X: x, Y: y}
}

// -------基本操作（填写各个方法体） -------

// Add 返回 v 和 other（按分量）的总和。
//
//	示例：Vec2{1,2}.Add(Vec2{3,4}) = Vec2{4,6}
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Sub 返回差值 v -other（按分量）。
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Scale 将 v 的每个分量乘以 s。
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// Dot 返回 v 和 other 的点积（内积）。
//
//	点 = v.X*其他.X + v.Y*其他.Y
func (v Vec2) Dot(other Vec2) float32 {
	return v.X*other.X + v.Y*other.Y
}

// 长度返回向量的大小（欧几里德范数）。
//
//	长度 = sqrt(X² + Y²)
//
// 提示：使用标准库中的 math.Sqrt。
func (v Vec2) Length() float32 {
	result := v.X*v.X + v.Y*v.Y
	return float32(math.Sqrt(float64(result)))
}

// LengthSquared 返回长度的平方（比较距离时避免开方）。
func (v Vec2) LengthSquared() float32 {
	result := v.X*v.X + v.Y*v.Y
	return result
}

// Normalize 返回与 v 方向相同的单位向量。
// 如果向量长度为零，则返回 Vec2{}。
func (v Vec2) Normalize() Vec2 {
	if v.LengthSquared() == 0 {
		return Vec2{}
	} else {
		l := v.Length()
		return Vec2{X: v.X / l, Y: v.Y / l}
	}

}

// Lerp 按因子 t 在 v 和目标之间执行线性插值。
//
//	lerp(a, b, t) = a + (b -a) *t
func (v Vec2) Lerp(target Vec2, t float32) Vec2 {
	return Vec2{
		X: v.X + (target.X-v.X)*t,
		Y: v.Y + (target.Y-v.Y)*t,
	}
}

// Negate 返回每个分量都被取反的向量。
func (v Vec2) Negate() Vec2 {
	// TODO：您的代码在这里
	return Vec2{
		X: -v.X,
		Y: -v.Y,
	}
}

// 字符串返回人类可读的表示形式。
func (v Vec2) String() string {
	return fmt.Sprintf("Vec2(%f, %f)", v.X, v.Y)
}
