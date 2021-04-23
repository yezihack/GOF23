package main

import (
	"fmt"
	"sync"
	"time"
)

// 控制反转 ioc

// 定义一个发送接口
type Sender interface {
	Send() // 发送函数
}

// 定义一个发送 app 结构体，实现控制反转操作
type SendIoc struct {
	list chan Sender // 使用 chan 的发送接口
}

var (
	// 定义一个全局的 Ioc 对象，保证全局唯一
	__sendApp     *SendIoc
	__sendAppOnce sync.Once
)

func NewSendIoc() *SendIoc {
	__sendAppOnce.Do(func() {
		__sendApp = &SendIoc{
			list: make(chan Sender, 10),
		}
		go __sendApp.run() // 进行调度操作，也就是在框架内实现控制反转
	})
	return __sendApp
}

// 注册接口
func (s *SendIoc) Register(ss Sender) {
	s.list <- ss
}

// 实现框架自动调度，即控制反转
func (s *SendIoc) run() {
	for {
		select {
		case f, ok := <-s.list:
			if ok {
				f.Send()
			}
		}
	}
}

// 实现 Sender 接口
type TxtSend struct {
	ID int
}

func (t *TxtSend) Send() {
	fmt.Printf("%d:使用 txt 发送消息\n", t.ID)
}

// 实现 Sender 接口
type SmsSend struct {
	ID int
}

func (s *SmsSend) Send() {
	fmt.Printf("%d:使用 sms 发送消息\n", s.ID)
}
func main() {
	// 向框架里注册，框架自动完成控制反转操作。
	for i := 0; i < 100; i++ {
		NewSendIoc().Register(&TxtSend{ID: i})
		NewSendIoc().Register(&SmsSend{ID: i})
	}
	time.Sleep(time.Millisecond * 1500)
}
