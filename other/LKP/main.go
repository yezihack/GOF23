package main

import "fmt"

func main() {
	t := &Teacher{}
	listGirl := make([]Girl, 0)
	for i := 0; i < 10; i++ {
		listGirl = append(listGirl, Girl{})
	}
	leader := NewGroupLeader(listGirl)
	t.Command(leader)
}

// 定义老师类
type Teacher struct {
}
// 老师类接触了体育委员 和 女生。这不符合迪米特原则。
// 改造为老师不接触女生。接触女生由体育委员接触
func (t Teacher) Command(gl *GroupLeader) {
	gl.CountGirl()
}
// 体育委员 接触 女生
type GroupLeader struct {
	girls []Girl
}
//func (g GroupLeader) CountGirl(listGirl []Girl) {
//	fmt.Println("girl count:", len(listGirl))
//}
func NewGroupLeader(listGirl []Girl) *GroupLeader{
	return &GroupLeader{girls: listGirl}
}
func (g GroupLeader) CountGirl() {
	fmt.Println("girl count:", len(g.girls))
}
type Girl struct {
}