package factory_method_pattern

import (
	"fmt"
	"testing"
)

func TestNewSaveSystemCreator(t *testing.T) {
	c := NewSaveSystemCreator().Create("local")
	c.PutData("s001", "zhangsan")
	data := c.GetData("s001")
	val, ok := data.(string)
	if !ok {
		t.Error("assert is failed")
	}
	fmt.Println("val:", val)

	c = NewSaveSystemCreator().Create("remote")
	c.PutData("s001", "lisi")
	data = c.GetData("s001")
	val, ok = data.(string)
	if !ok {
		t.Error("assert is failed")
	}
	fmt.Println("val:", val)
}
