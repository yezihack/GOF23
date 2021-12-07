package single_pattern

import "fmt"

// 饿汉模式
type HungryMan struct {
}

var hMan = new(HungryMan)

func (h *HungryMan) Display() {
	fmt.Println("hello hungry man!")
}
