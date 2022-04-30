// 命令模式（Command Pattern）是一种数据驱动的设计模式，它属于行为型模式。请求以命令的形式包裹在对象中，并传给调用对象。调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象，该对象执行命令。式
package main

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &tv{}
	onCommand := &onCommand{
		device: tv,
	}
	offCommand := &offCommand{
		device: tv,
	}
	onButton := &button{
		command: onCommand,
	}
	offButton := &button{
		command: offCommand,
	}

	onButton.press()
	offButton.press()
}
