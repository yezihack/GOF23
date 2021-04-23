package main

import "fmt"

func main() {
	s := NewStore()
	s.Add(NewNovelBook("论语",  "孔子", 2500))
	s.Add(NewNovelBook("苏菲的世界", "乔斯坦·贾德", 2900))
	s.Add(NewNovelBook("孙子兵法",  "孙武", 1800))
	s.Add(NewOffNovelBook("孙子兵法",  "孙武", 1800))
	s.Show()
}

type Store struct {
	list []IBook
}

func NewStore() *Store {
	return &Store{
		list: []IBook{},
	}
}
func (s *Store) Add(book IBook) {
	s.list = append(s.list,book)
}
func (s *Store) Show() {
	for _, item := range s.list {
		fmt.Printf("书名：%s, 作者：%s, 价格：%d\n", item.GetName(), item.GetAuthor(), item.GetPrice()/100)
	}
}

type IBook interface {
	GetName() string
	GetPrice() int
	GetAuthor() string
}
type NovelBook struct {
	name string
	price int
	author string
}

func NewNovelBook(name, author string, price int) *NovelBook {
	return &NovelBook{
		name: name,
		author: author,
		price: price,
	}
}


func (n *NovelBook) GetName() string {
	return n.name
}

func (n *NovelBook) GetPrice() int {
	return n.price
}

func (n *NovelBook) GetAuthor() string {
	return n.author
}
type OffNovelBook struct {
	name string
	price int
	author string
}
func NewOffNovelBook(name, author string, price int) IBook {
	return &OffNovelBook{
			name:   name,
			price:  price,
			author: author,
	}
}
func (n *OffNovelBook) GetName() string {
	return n.name
}

func (n *OffNovelBook) GetPrice() int {
	if n.price > 1000 {
		n.price *= 90/100
	} else {
		n.price *= 80/100
	}
	return n.price
}

func (n *OffNovelBook) GetAuthor() string {
	return n.author
}
