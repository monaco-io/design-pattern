// 建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。
// 一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。
package main

import "fmt"

type iBuilder interface {
	buildWindow()
	buildDoor()
	buildRoom()
	getHouse() house
}

type house struct {
	window string
	door   string
	room   string
}

func Build(builder iBuilder) house {
	builder.buildWindow()
	builder.buildDoor()
	builder.buildRoom()
	return builder.getHouse()
}

// 普通房子建造
type normalBuilder struct {
	window string
	door   string
	room   string
}

func (b *normalBuilder) buildWindow() {
	b.window = "Normal Window"
}

func (b *normalBuilder) buildDoor() {
	b.door = "Normal Door"
}

func (b *normalBuilder) buildRoom() {
	b.room = "Normal Room"
}

func (b *normalBuilder) getHouse() house {
	return house{
		window: b.window,
		door:   b.door,
		room:   b.room,
	}
}

// 大房子建造
type bigBuilder struct {
	window string
	door   string
	room   string
}

func (b *bigBuilder) buildWindow() {
	b.window = "Big Window"
}

func (b *bigBuilder) buildDoor() {
	b.door = "Big Door"
}

func (b *bigBuilder) buildRoom() {
	b.room = "Big Room"
}

func (b *bigBuilder) getHouse() house {
	return house{
		window: b.window,
		door:   b.door,
		room:   b.room,
	}
}

func main() {
	normalhouse := Build(new(normalBuilder))
	fmt.Println(normalhouse)

	bigHouse := Build(new(bigBuilder))
	fmt.Println(bigHouse)
}
