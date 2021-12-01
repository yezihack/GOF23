package abstractfactorypattern

/*
1. 抽象产品族，即多个抽象产品
2. 具体产品族，即多个具体产品
3. 抽象工厂族，即多个抽象工厂
4. 具体工厂族，即多个具体工厂
*/

// 具体产品，如鞋子
type Shoe struct {
	size     int    // 大小
	meterial string // 颜色
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}
func (s *Shoe) SetMaterial(meterial string) {
	s.meterial = meterial
}
func (s *Shoe) GetSize() int {
	return s.size
}
func (s *Shoe) GetMaterial() string {
	return s.meterial
}

// 具体品牌的鞋子
type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

// 具体产品02， 如裤子
type Trousers struct {
	size  int    // 大小
	color string // 颜色
}

func (s *Trousers) SetSize(size int) {
	s.size = size
}
func (s *Trousers) SetColor(color string) {
	s.color = color
}
func (s *Trousers) GetSize() int {
	return s.size
}
func (s *Trousers) GetColor() string {
	return s.color
}

// 具体品牌的鞋子
type AdidasTrousers struct {
	Trousers
}

type NikeTrousers struct {
	Trousers
}
