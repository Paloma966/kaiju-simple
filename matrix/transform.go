package matrix

// Transform 表示一个 3D 变换，包含位置、旋转和缩放。
// 这是引擎中 Entity 使用的核心空间类型。
//
// 在真实的 kaiju 引擎中，旋转用四元数存储以避免万向节死锁。
// 这里为简单起见使用欧拉角。
type Transform struct {
	Position Vec3
	Rotation Vec3 // Euler angles in radians (pitch=X, yaw=Y, roll=Z)
	Scale    Vec3
}

// NewTransform 创建一个在 position 处的 Transform，旋转为零，缩放为 1。
func NewTransform(position Vec3) Transform {
	return Transform{
		Position: position,
		Rotation: Vec3{},          // TODO: 默认旋转应该是什么？
		Scale:    NewVec3(1, 1, 1), // TODO: 默认缩放应该是什么？
	}
}

// Matrix 计算该变换对应的 4x4 模型矩阵。
//
// 标准顺序为：缩放 → 旋转 → 平移
//   model = T * Rz * Ry * Rx * S
//
// （注意：矩阵乘法从右向左，所以最先应用于顶点的操作在最右边的矩阵。）
func (t Transform) Matrix() Mat4 {
	// TODO: YOUR CODE HERE
	// 1. Build translation matrix from t.Position
	// 2. Build rotation matrices (RotateX, RotateY, RotateZ) from t.Rotation
	// 3. Build scale matrix from t.Scale
	// 4. Combine: Translate * RotateZ * RotateY * RotateX * Scale
	return Identity()
}

// Forward 返回前方向量（变换"面朝"的方向）。
// 在右手坐标系（kaiju 使用此坐标系）中：
//
//	Forward = 负 Z 轴经过该变换的旋转后的方向
//	默认前方向（无旋转）= Vec3{0, 0, -1}
func (t Transform) Forward() Vec3 {
	// TODO: YOUR CODE HERE
	// Hint: use Mat4.RotateZ, RotateY, RotateX and MultiplyVec3Direction
	return NewVec3(0, 0, -1)
}

// Right 返回右方向量。
//
//	默认右方向（无旋转）= Vec3{1, 0, 0}
func (t Transform) Right() Vec3 {
	// TODO: YOUR CODE HERE
	return NewVec3(1, 0, 0)
}

// Up 返回上方向量。
//
//	默认上方向（无旋转）= Vec3{0, 1, 0}
func (t Transform) Up() Vec3 {
	// TODO: YOUR CODE HERE
	return NewVec3(0, 1, 0)
}
