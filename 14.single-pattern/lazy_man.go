package main

import (
	"fmt"
	"sync"
)

// 懒汉式之单例模式
// 相对于饿汉式模式的优势是支持延迟加载。
type LazyMan struct {
}
var (
	lman *LazyMan
	lmanLock sync.Mutex
)

func NewLazyMan() *LazyMan {
	if lman == nil {
		defer lmanLock.Unlock()
		lmanLock.Lock()
		lman = new(LazyMan)
	}
	return lman
}
func main() {
	lazyMan := NewLazyMan()
	lazyMan2 := NewLazyMan()
	fmt.Printf("%p\n%p\n", lazyMan, lazyMan2)
}
