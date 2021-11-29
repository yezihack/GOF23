package factory_method_pattern

import (
	"errors"
	"fmt"
)

// 定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类
/*
工厂方法实现过程：
1. 首先定义一个抽象产品 product，该产品负责定义产品的共性。
2. 再定义一个工厂，creator， 创建具体产品由具体工厂来完成。

*/

// 以下模仿女娲造人的过程
// 该过程涉及三个对象：女娲、八卦炉、三种不同肤色的人
// 1. 八卦炉类似于一个工厂，负责制造生产产品（即人类）
// 2. 三种不同肤色的人，他们都是同一个接口下的不同实现类，都是人嘛，只是肤
// 色、语言不同，对于八卦炉来说都是它生产出的产品

// 定义人的接口，具有人的形为
// 下面白，黑，黄种人实现此接口
// 这就是抽象产品 product，负责定义产品的共性。
type Human interface {
	GetColor() // 不同肤色
	Talk()     // 不同语言
}

// 白种人
type WhiteHuman struct {
}

// 黑种人
type BlackHuman struct {
}

// 黄种人
type YellowHuman struct {
}

// start 以下是不同人种的具体实现形为
func (h *WhiteHuman) GetColor() {
	fmt.Println("白人是白色皮肤")
}
func (h *WhiteHuman) Talk() {
	fmt.Println("白人说英语")
}

func (h *BlackHuman) GetColor() {
	fmt.Println("黑人是黑色皮肤")
}
func (h *BlackHuman) Talk() {
	fmt.Println("黑人说也说英语")
}

func (h *YellowHuman) GetColor() {
	fmt.Println("黄种人是黄色皮肤")
}
func (h *YellowHuman) Talk() {
	fmt.Println("黄种人说汉语")
}

// end 以上是不同人种的具体实现形为

// 定义一个八卦炉的接口，用于生产人。
// 定义一个工厂，creator，
type HumanFactoryer interface {
	CreateHuman(man string) (Human, error)
}

// 创建一个造人工厂
type HumanFactory struct {
}

func (h *HumanFactory) CreateHuman(man string) (Human, error) {
	switch man {
	case "white":
		return new(WhiteHuman), nil
	case "black":
		return new(BlackHuman), nil
	case "yellow":
		return new(YellowHuman), nil
	default:
		return nil, errors.New("not support type")
	}
}

// 人种有了，八卦炉也有了，剩下的工作就是女娲采集黄土，然后命令八卦炉开始生产

func NewCreateHuman() HumanFactoryer {
	return &HumanFactory{}
}
