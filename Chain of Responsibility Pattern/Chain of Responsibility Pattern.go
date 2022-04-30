// 顾名思义，责任链模式（Chain of Responsibility Pattern）为请求创建了一个接收者对象的链。这种模式给予请求的类型，对请求的发送者和接收者进行解耦。这种类型的设计模式属于行为型模式。
// 在这种模式中，通常每个接收者都包含对另一个接收者的引用。如果一个对象不能处理该请求，那么它会把相同的请求传给下一个接收者，依此类推。
package main

import (
	"fmt"
)

type iPersion interface {
	exec(*persion996)
	setNext(iPersion)
}

type persion996 struct {
	name             string
	workDone         bool
	bossPuaDone      bool
	wifeGetMondyDone bool
}

type you struct {
	next iPersion
}
type boss struct {
	next iPersion
}
type yourWife struct {
	next iPersion
}

func (you *you) setNext(ip iPersion) {
	you.next = ip
}

func (you *you) exec(p996 *persion996) {
	if p996.name == "" {
		panic("no name")
	}
	p996.workDone = true
	fmt.Println("you worked")
	you.next.exec(p996)
}

func (boss *boss) setNext(ip iPersion) {
	boss.next = ip
}

func (boss *boss) exec(p996 *persion996) {
	if !p996.workDone {
		panic("workDone")
	}
	p996.bossPuaDone = true
	fmt.Println("boss puaed you")
	boss.next.exec(p996)
}

func (yourWife *yourWife) setNext(ip iPersion) {
	yourWife.next = ip
}

func (yourWife *yourWife) exec(p996 *persion996) {
	if !p996.bossPuaDone {
		panic("bossPuaDone")
	}
	p996.wifeGetMondyDone = true
	fmt.Println("your wife got money")
}

func main() {
	you := &you{}

	boss := &boss{}
	you.setNext(boss)

	yourWife := &yourWife{}
	boss.setNext(yourWife)

	persion996 := &persion996{name: "da_gong_ren"}
	you.exec(persion996)
}
