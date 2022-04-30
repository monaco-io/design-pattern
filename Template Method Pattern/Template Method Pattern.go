// 在模板模式（Template Pattern）中，一个抽象类公开定义了执行它的方法的方式/模板。它的子类可以按需要重写方法实现，但调用将以抽象类中定义的方式进行。这种类型的设计模式属于行为型模式。
package main

import "fmt"

type Game interface {
	Start()
	Stop()
}

type LOL struct{}

func (game LOL) Start() {
	fmt.Println("lol start")
}

func (game LOL) Stop() {
	fmt.Println("lol stop")
}

type DNF struct{}

func (game DNF) Start() {
	fmt.Println("DNF start")
}

func (game DNF) Stop() {
	fmt.Println("DNF stop")
}

func main() {
	var game Game

	game = new(DNF)
	game.Start()
	game.Stop()

	game = new(LOL)
	game.Start()
	game.Stop()
}
