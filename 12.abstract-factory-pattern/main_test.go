package abstractfactorypattern

import (
	"errors"
	"fmt"
	"testing"
)

func NewFactory(brand string) (ISportFactory, error) {
	switch brand {
	case "nike":
		return &NikeFactory{}, nil
	case "adidas":
		return &AdidasFactory{}, nil
	default:
		return nil, errors.New("Wrong brand type passed")
	}
}

func TestAbstractFactory(t *testing.T) {
	adidasFactory, err := NewFactory("nike")
	if err != nil {
		t.Error("adidas err:", err)
	}
	adidasShoe := adidasFactory.MakeShoe()
	adidasShoe.SetSize(120)
	adidasShoe.SetMaterial("布料")
	fmt.Printf("adidas 鞋子, size: %d, material:%s \n",
		adidasShoe.GetSize(), adidasShoe.GetMaterial())
}
