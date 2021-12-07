package factory_method_pattern

import (
	"testing"
)

func TestNewCreateHuman(t *testing.T) {
	man, err := NewCreateHuman().CreateHuman("yellow")
	if err != nil {
		t.Error("yellow man, err:", err)
	}
	man.GetColor()
	man.Talk()
	man, err = NewCreateHuman().CreateHuman("black")
	if err != nil {
		t.Error("black man, err:", err)
	}
	man.GetColor()
	man.Talk()
	man, err = NewCreateHuman().CreateHuman("white")
	if err != nil {
		t.Error("white man, err:", err)
	}
	man.GetColor()
	man.Talk()

}
