package main

import "fmt"

func main() {
	xishi := NewSearcher("西施")
	xishi.Show()
}

// 定义漂亮姑娘接口
type IPettyGirl interface {
	IGoodGirl
	IGreatTemperament
}
type IGoodGirl interface {
	GoodLooking()
	NiceFigure()
}
type IGreatTemperament interface {
	GreatTemperament()
}
// 定义星探接口
type ISearcher interface {
	Show()
}
/*************************************/
// 实现漂亮姑娘接口
type PettyGirl struct {
	Name string
}
func (p *PettyGirl) GoodLooking() {
	fmt.Printf("%s 脸蛋好看\n", p.Name)
}
func (p *PettyGirl) NiceFigure() {
	fmt.Printf("%s 好身材\n", p.Name)
}
func (p *PettyGirl) GreatTemperament() {
	fmt.Printf("%s 好气质\n", p.Name)
}
/*************************************/
type Searcher struct {
	p PettyGirl
}

func (s Searcher) Show() {
	s.p.GoodLooking()
	s.p.NiceFigure()
	s.p.GreatTemperament()
}

func NewSearcher(name string) ISearcher {
	return &Searcher{
		p: PettyGirl{
			Name: name,
		},
	}
}
type Search2 struct {
	p IGreatTemperament
}
func (s Search2) Show() {
	s.p.GreatTemperament()
}
