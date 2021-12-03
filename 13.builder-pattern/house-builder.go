package builderpattern

import (
	"fmt"
)

// Builder：抽象建造者

type IHouseBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	Result() HouseProduct
}

// 选择不同的具体建造者.
func GetBuilder(builderType string) IHouseBuilder {
	switch builderType {
	case "normal":
		return &NormalBuilder{}
	case "igloo":
		return &IglooBuilder{}
	default:
		return nil
	}
}

// 1. ConcreteBuilder：具体建造者 01
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (n *NormalBuilder) SetWindowType() {
	n.windowType = "Wooden Window"
}
func (n *NormalBuilder) SetDoorType() {
	n.doorType = "Wooden Door"
}
func (n *NormalBuilder) SetNumFloor() {
	n.floor = 2
}
func (n *NormalBuilder) Result() HouseProduct {
	return HouseProduct{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}

// 1. ConcreteBuilder：具体建造者 02
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (n *IglooBuilder) SetWindowType() {
	n.windowType = "igloo Window"
}
func (n *IglooBuilder) SetDoorType() {
	n.doorType = "igloo Door"
}
func (n *IglooBuilder) SetNumFloor() {
	n.floor = 1
}
func (n *IglooBuilder) Result() HouseProduct {
	return HouseProduct{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}

// 1. Product：产品角色
type HouseProduct struct {
	windowType string
	doorType   string
	floor      int
}

// 1. Director：指挥者
// 指挥者负责按一定的顺序组装
type Director struct {
	builder IHouseBuilder
}

func NewDirector(builder IHouseBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(b IHouseBuilder) {
	d.builder = b
}

// 开始建造一个房子,按一定的顺序建造而成.
func (d *Director) BuildHouse() HouseProduct {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloor()
	return d.builder.Result()
}

// 客户
// 客户只需要知道建造类型即可.无须了解内部细节.
func HouseClient() {
	// 获取普通房的建造师
	normalBuilder := GetBuilder("normal")
	// 获取冰房子的建造师
	IglooBuilder := GetBuilder("igloo")

	// 请一位指挥者可以指挥建造师
	// 即把普通房子的建造师交给指挥者调遣
	director := NewDirector(normalBuilder)

	fmt.Printf("normal window:%s\n", director.BuildHouse().windowType)
	fmt.Printf("normal door:%s\n", director.BuildHouse().doorType)
	fmt.Printf("normal floor:%d\n", director.BuildHouse().floor)

	director.SetBuilder(IglooBuilder)
	fmt.Printf("normal window:%s\n", director.BuildHouse().windowType)
	fmt.Printf("normal door:%s\n", director.BuildHouse().doorType)
	fmt.Printf("normal floor:%d\n", director.BuildHouse().floor)
}
