package abstractfactorypattern

/*
1. 抽象产品族，即多个抽象产品
2. 具体产品族，即多个具体产品
3. 抽象工厂族，即多个抽象工厂
4. 具体工厂族，即多个具体工厂
*/

// 1. 抽象产品族-01，即多个抽象产品
type IShoeProduct interface {
	SetSize(size int)            // 设置不同尺寸
	SetMaterial(meterial string) // 设置不同材质
	GetSize() int                // 获取尺寸
	GetMaterial() string         // 获取材质
}

// 1. 抽象产品族-02，即多个抽象产品
type ITrousersProduct interface {
	SetSize(size int)      // 设置不同尺寸
	SetColor(color string) // 设置颜色
	GetSize() int          // 获取尺寸
	GetColor() string      // 获取颜色
}
