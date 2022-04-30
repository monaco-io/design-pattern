// 在访问者模式（Visitor Pattern）中，我们使用了一个访问者类，它改变了元素类的执行算法。通过这种方式，元素的执行算法可以随着访问者改变而改变。这种类型的设计模式属于行为型模式。根据模式，元素对象已接受访问者对象，这样访问者对象就可以处理元素对象上的操作。
package main

import "fmt"

type iComputerPart interface {
	Accept(v iComputerPartVisitor)
}
type iComputerPartVisitor interface {
	Visit(p iComputerPart)
}
type KeyBoardVisitor struct{}
type KeyBoard struct{}

func (part KeyBoard) Accept(v iComputerPartVisitor) {
	fmt.Println("keyboard accepd")
}

func (v KeyBoardVisitor) Visit(part iComputerPart) {
	fmt.Println("visit keyboard")
}

func Accept(v iComputerPartVisitor, p iComputerPart) {
	p.Accept(v)
}
func Visit(v iComputerPartVisitor, p iComputerPart) {
	v.Visit(p)
}

func main() {
	kbd := KeyBoard{}
	kbdV := KeyBoardVisitor{}

	Accept(kbdV, kbd)
}
