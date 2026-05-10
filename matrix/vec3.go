package matrix

import (
	"fmt"
	"math"
)

// Vec3 是 3D 向量，用于 3D 空间中的位置、方向和颜色。
type Vec3 struct {
	X, Y, Z float32
}

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

// ------- 基本操作 -------

func (v Vec3) Add(other Vec3) Vec3 {
	// Todo：您的代码在这里
	return Vec3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	// Todo：您的代码在这里
	return Vec3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3) Scale(s float32) Vec3 {
	// Todo：您的代码在这里
	return Vec3{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

func (v Vec3) Dot(other Vec3) float32 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross 返回叉积 v × other。
// 叉积产生一个垂直于两个输入向量的向量。
//
//	cross.X = v.Y*other.Z - v.Z*other.Y
//	cross.Y = v.Z*other.X - v.X*other.Z
//	cross.Z = v.X*other.Y - v.Y*other.X
func (v Vec3) Cross(other Vec3) Vec3 {

	return Vec3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vec3) Length() float32 {
	result := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	return float32(math.Sqrt(float64(result)))

}

func (v Vec3) LengthSquared() float32 {
	// Todo：您的代码在这里
	result := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	return result
}

func (v Vec3) Normalize() Vec3 {
	if v.LengthSquared() == 0 {
		return Vec3{}
	}
	l := v.Length()
	return Vec3{X: v.X / l, Y: v.Y / l, Z: v.Z / l}
}

func (v Vec3) Lerp(target Vec3, t float32) Vec3 {

	return Vec3{
		X: v.X + (target.X-v.X)*t,
		Y: v.Y + (target.Y-v.Y)*t,
		Z: v.Z + (target.Z-v.Z)*t,
	}
}

func (v Vec3) Negate() Vec3 {

	return Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vec3) String() string {
	return fmt.Sprintf("Vec3(%f, %f, %f)", v.X, v.Y, v.Z)
}
