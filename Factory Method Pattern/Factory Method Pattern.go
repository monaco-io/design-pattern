// 工厂模式（Factory Pattern）是 Java 中最常用的设计模式之一。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。
// 在工厂模式中，我们在创建对象时不会对客户端暴露创建逻辑，并且是通过使用一个共同的接口来指向新创建的对象。
package main

import "fmt"

type iStudent interface {
	getName() string
	getAge() int8
}

func getStudent(name string) iStudent {
	switch name {
	case "xiaoming":
		return newXiaoming()
	case "xiaohong":
		return newXiaohone()
	}
	return nil
}

type student struct {
	name string
	age  int8
}

func (s *student) getName() string {
	return s.name
}

func (s *student) getAge() int8 {
	return s.age
}

type xiaoming struct {
	student
}

func newXiaoming() iStudent {
	return &xiaoming{
		student: student{
			name: "xiaoming",
			age:  20,
		},
	}
}

type xiaohong struct {
	student
}

func newXiaohone() iStudent {
	return &xiaoming{
		student: student{
			name: "xiaohong",
			age:  21,
		},
	}
}

func main() {
	xiaoming := getStudent("xiaoming")
	xiaohong := getStudent("xiaohong")

	fmt.Println(xiaoming)
	fmt.Println(xiaohong)
}
