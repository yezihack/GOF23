package simple_factory

import "fmt"

// 定义专属类型
type Tool int

const (
	TruckTool Tool = iota
	ShipTool
)

// 定义交通抽象形为接口
type Transport interface {
	Deliver()
}

// 定义两种交通工具， 如下：货运卡车，轮船

// Truck 货运卡车运输工具
type Truck struct {
}

// Ship 轮船运输工具
type Ship struct {
}

// 实现抽象接口
func (t *Truck) Deliver() {
	fmt.Println("I'm truck, I can deliver")
}

// 实现抽象接口
func (s *Ship) Deliver() {
	fmt.Println("I'm ship, I can deliver")
}

// NewTransport 工厂类, 根据不同的类型,创建不同的实例
func NewTransport(tool Tool) Transport {
	switch tool {
	case TruckTool:
		return &Truck{}
	case ShipTool:
		return &Ship{}
	}
	return nil
}
