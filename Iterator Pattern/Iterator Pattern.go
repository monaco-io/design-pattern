// 迭代器模式（Iterator Pattern）是 Java 和 .Net 编程环境中非常常用的设计模式。这种模式用于顺序访问集合对象的元素，不需要知道集合对象的底层表示。
// 迭代器模式属于行为型模式。
package main

import "fmt"

type iterator interface {
	hasNext() bool
	getNext() *string
}

type nodeIteractor struct {
	nodes []*string
	index int
}

func (ni *nodeIteractor) hasNext() bool {
	return ni.index < len(ni.nodes)
}
func (ni *nodeIteractor) getNext() *string {
	s := ni.nodes[ni.index]
	ni.index++
	return s
}

func createIteractor(nodes []*string) nodeIteractor {
	return nodeIteractor{
		nodes: nodes,
	}
}

func main() {
	var (
		n1 = "n1"
		n2 = "n2"
		n3 = "n3"
	)
	iter := createIteractor([]*string{&n1, &n2, &n3})
	for iter.hasNext() {
		fmt.Println(*iter.getNext())
	}
}
