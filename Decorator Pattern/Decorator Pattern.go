// 装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

// 这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。
package main

import "fmt"

type Message string

func (msg Message) Print() {
	fmt.Println(msg)
}

type DetailMessage struct {
	Message
	Detail string
}

func (msg DetailMessage) PrintDetail() {
	fmt.Println(msg.Message + ": " + Message(msg.Detail))
}

func main() {
	msg := Message("a message")
	detailMsg := DetailMessage{msg, "some detail data"}

	detailMsg.Print()
	detailMsg.PrintDetail()
}
