// 组合模式（Composite Pattern），又叫部分整体模式，是用于把一组相似的对象当作一个单一的对象。组合模式依据树形结构来组合对象，用来表示部分以及整体层次。这种类型的设计模式属于结构型模式，它创建了对象组的树形结构。

// 这种模式创建了一个包含自己对象组的类。该类提供了修改相同对象组的方式。

// 我们通过下面的实例来演示组合模式的用法。实例演示了一个组织中员工的层次结构。
package main

import "fmt"

type Node interface {
	Print(basePath string)
}

type Foder struct {
	name     string
	Children []Node
}

func (f Foder) Print(basePath string) {
	basePath += "/" + f.name
	fmt.Println(basePath)
	for _, v := range f.Children {
		v.Print(basePath)
	}
}

type File string

func (f File) Print(basePath string) {
	fmt.Println(basePath + "/" + string(f))
}
func main() {
	f1 := File("f1")
	f2 := File("f2")
	f3 := File("f3")

	fd1 := Foder{"fd1", []Node{f1, f2}}
	fd2 := Foder{"fd2", []Node{fd1, f3}}
	fd2.Print("")
}
