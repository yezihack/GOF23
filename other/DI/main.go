package main

import "fmt"

// DI 依赖注入:Dependency Injection

// 定义一个发送接口类
type ManagerSender interface {
	// 发送类
	Send(phone, message string)
}

// 定义短信发送结构体，实现 ManagerSender 接口
type SmsSender struct {
}

func (s *SmsSender) Send(phone, message string) {
	fmt.Printf("正在使用短信发送消息，手机号:%s,消息内容:%s\n", phone, message)
}

// 定义站内发送结构体，实现 ManagerSender 接口
type InboxSender struct {
}

func (s *InboxSender) Send(phone, message string) {
	fmt.Printf("正在使用站内发送消息，手机号:%s,消息内容:%s\n", phone, message)
}

// ----------------- 实现注入的代码 --------------------
// 定义一个发消息结构体
type Notification struct {
	ms ManagerSender
}

// 将接口赋值给 Notification
func NewNotification(m ManagerSender) *Notification {
	return &Notification{ms: m}
}

// 实现发消息，使用接口的对象 Send
func (n *Notification) SendMessage(phone, message string) {
	n.ms.Send(phone, message)
}

func main() {
	// 传入不同的结构体，实现不同的类型发送消息。
	NewNotification(&SmsSender{}).SendMessage("15112341234", "你回来没有？")
	NewNotification(&InboxSender{}).SendMessage("13198771100", "我回来了!")
}
