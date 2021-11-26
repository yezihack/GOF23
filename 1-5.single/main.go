package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 1. go 自带单例模式

var single sync.Once

func ReadBook(fn func()) {
	single.Do(func() {
		fn()
	})
}

// 2. 使用原子操作实现单例 check-lock-check
var (
	done int32
	lock sync.Mutex
)

func CustomSingle(fn func()) {
	if atomic.LoadInt32(&done) == 0 {
		// 加锁操作
		lock.Lock()
		defer lock.Unlock()
		// 再次判断 done 是否为0，防止并发
		if done == 0 {
			atomic.AddInt32(&done, 1) // 对 done 进行赋值 1
			fn()                      // 执行单例
		}
	}
}

func main() {
	count := 1000
	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ReadBook(func() {
				fmt.Println("I am reading book")
			})
		}()
	}
	// I am reading book

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CustomSingle(func() {
				fmt.Println("I am look movie")
			})
		}()
	}
	// I am look movie
	wg.Wait()
}
