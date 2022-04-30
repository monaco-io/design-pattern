// 在状态模式（State Pattern）中，类的行为是基于它的状态改变的。这种类型的设计模式属于行为型模式。

// 在状态模式中，我们创建表示各种状态的对象和一个行为随着状态对象改变而改变的 context 对象。
package main

import "fmt"

type State struct {
	status string
}

func (s *State) Status() string {
	return s.status
}

func (s *State) Set(sta string) {
	s.status = sta
}

func Start(s *State) {
	s.Set("Start")
}

func Stop(s *State) {
	s.Set("Stop")
}

func main() {
	var s State
	Start(&s)
	fmt.Println(s)
	Stop(&s)
	fmt.Println(s)
}
