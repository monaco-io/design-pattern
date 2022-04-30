// 中介者模式（Mediator Pattern）是用来降低多个对象和类之间的通信复杂性。这种模式提供了一个中介类，该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护。中介者模式属于行为型模式。

package main

import "fmt"

type ChatBox struct {
	msg []string
}

func (box *ChatBox) SendMsg(user User, msg string) {
	box.msg = append(box.msg, fmt.Sprintf("%s: %s", user.Name, msg))
}
func (box *ChatBox) ShowMsg() {
	for _, v := range box.msg {
		fmt.Println(v)
	}
}

type User struct {
	Name string
}

func main() {
	box := new(ChatBox)

	xiaoming := User{"xiaoming"}
	xiaohong := User{"xiaohong"}

	box.SendMsg(xiaoming, "hello, i am xiaoming")
	box.SendMsg(xiaohong, "hello, good moning")

	box.ShowMsg()
}
