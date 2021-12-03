package builderpattern

import "fmt"

// 建造者模式,实现一个造汽车的程序

// 首先定义一个造汽车所需部件. 一辆汽车简化分为引擎,车身,底盘三种组成.

// 定义建造者模型接口
type ICarBuilder interface {
	SetEngine()     // 引擎
	SetBody()       // 车身
	SetFloor()      // 底盘
	Result() string // 生产出一台车
}

// 定义指挥者模型
type CarDirector struct {
	build ICarBuilder
}

func NewCarDirector(b ICarBuilder) *CarDirector {
	return &CarDirector{
		build: b,
	}
}
func (d *CarDirector) SetBuilder(b ICarBuilder) {
	d.build = b
}

func (d *CarDirector) BuildCar() string {
	d.build.SetBody()
	d.build.SetEngine()
	d.build.SetFloor()
	return d.build.Result()
}

// 定义具体建造者.01

type BMWCarBuilder struct {
	engine string
	body   string
	floor  string
}

func (b *BMWCarBuilder) SetEngine() {
	b.engine = "X引擎"
}
func (b *BMWCarBuilder) SetBody() {
	b.body = "蓝色外壳"
}
func (b *BMWCarBuilder) SetFloor() {
	b.floor = "合金底盘"
}
func (b *BMWCarBuilder) Result() string {
	return fmt.Sprintf("品牌:%s, 引擎:%s, 车身:%s, 底盘:%s\n", "宝马", b.engine, b.body, b.floor)
}

// 定义具体建造者.02
type BenZCarBuilder struct {
	engine string
	body   string
	floor  string
}

func (b *BenZCarBuilder) SetEngine() {
	b.engine = "Z引擎"
}
func (b *BenZCarBuilder) SetBody() {
	b.body = "黑色外壳"
}
func (b *BenZCarBuilder) SetFloor() {
	b.floor = "钛合金底盘"
}
func (b *BenZCarBuilder) Result() string {
	return fmt.Sprintf("品牌:%s, 引擎:%s, 车身:%s, 底盘:%s\n", "奔驰", b.engine, b.body, b.floor)
}

// 客户端
func CarClient() {
	director := NewCarDirector(&BenZCarBuilder{})
	fmt.Println(director.BuildCar())

	director.SetBuilder(&BMWCarBuilder{})
	fmt.Println(director.BuildCar())
}
