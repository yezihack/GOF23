package abstractfactorypattern

// 定义抽象工厂
// 一个抽象工厂可以定义多个不同的抽象产品。
// 即一个工厂可以生成不同产品，如阿迪工厂可以生成鞋子和裤子。

type ISportFactory interface {
	MakeShoe() IShoeProduct
	MakeTrousers() ITrousersProduct
}
