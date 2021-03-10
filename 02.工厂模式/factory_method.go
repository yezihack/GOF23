package main

// 卡车Truck和 轮船Ship类都必须实现 运输Transport接口， 该接口声明了一个名为 deliver交付的方法
type Transport interface {
	Deliver() // 交付
}

// 每个类都将以不同的方式实现该方法： 卡车走陆路交付货物， 轮船走海路交付货物。
// 陆路运输RoadLogistics类中的工厂方法返回卡车对象，
// 海路运输SeaLogistics类则返回轮船对象。
type Truck struct {
}

func (t *Truck) Deliver() {

}

type Ship struct {
}

func (s *Ship) Deliver() {

}

// 只要产品类实现一个共同的接口， 你就可以将其对象传递给客户代码， 而无需提供额外数据。

//调用工厂方法的代码 （通常被称为客户端代码） 无需了解不同子类返回实际对象之间的差别。
//客户端将所有产品视为抽象的 运输 。 客户端知道所有运输对象都提供 交付方法， 但是并不关心其具体实现方式。
