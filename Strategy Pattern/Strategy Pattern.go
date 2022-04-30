// 在策略模式（Strategy Pattern）中，一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式。

// 在策略模式中，我们创建表示各种策略的对象和一个行为随着策略对象改变而改变的 context 对象。策略对象改变 context 对象的执行算法。
package main

import "fmt"

type Strategy interface {
	Do(a, b int) int
}

type Add struct{}

func (Add) Do(a, b int) int {
	return a + b
}

type Sub struct{}

func (Sub) Do(a, b int) int {
	return a - b
}

func Do(strategy Strategy, a, b int) int {
	return strategy.Do(a, b)
}

func main() {
	var (
		add Add
		sub Sub
	)
	fmt.Println(Do(add, 1, 1))
	fmt.Println(Do(sub, 1, 1))
}
