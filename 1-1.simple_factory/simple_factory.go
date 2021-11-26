package simple_factory

import "fmt"

type TransportTool int
const (
	TruckTool TransportTool = iota
	ShipTool
)

type Transport interface {
	Deliver()
}

// Truck 货运卡车运输工具
type Truck struct {

}

func (t *Truck) Deliver() {
	fmt.Println("I'm truck, I can deliver")
}

// Ship 轮船运输工具
type Ship struct {

}
func (s *Ship) Deliver() {
	fmt.Println("I'm ship, I can deliver")
}

// NewTransport 工厂类, 根据不同的类型,创建不同的实例
func NewTransport(tool TransportTool) Transport {
	switch tool {
	case TruckTool:
		return &Truck{}
	case ShipTool:
		return &Ship{}
	}
	return nil
}
