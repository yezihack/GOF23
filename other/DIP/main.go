package main

import "fmt"


// 依赖倒置原则
// 1. 面向接口编程，而不是面向实现编程。
// 2. 高层不依赖低层。 两者都应该依赖抽象
// 3. 抽象不依赖细节。
// 4. 细节应该依赖抽象。

// 高层模块：定义司机开汽车的接口
type Driver interface {
	Drive(carer Carer)
}

// 低层模块：实现开车接口
type DriveMan struct {
}
func (d DriveMan) Drive(carer Carer) {
	carer.Run()
}
func NewDriveMan() Driver {
	return &DriveMan{}
}


// 高层模块： 定义汽车接口，适合任何品牌的车
type Carer interface {
	Run()
}

// 低层模块：宝马车
type BMW struct {

}
func (b BMW) Run() {
	fmt.Println("宝马车开车了")
}
// 低层模块：奔驰车
type Benz struct {

}
func (b Benz) Run() {
	fmt.Println("奔驰车开车了")
}

// 实现细节
func main() {
	zhangSan := NewDriveMan()
	zhangSan.Drive(&BMW{})
	zhangSan.Drive(&Benz{})
}