// 享元模式（Flyweight Pattern）主要用于减少创建对象的数量，以减少内存占用和提高性能。这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式。
package main

var colorMp = make(map[string]string)

func newColor(color string) string {
	colorMp[color] = color
	return colorMp[color]
}
func getColor(color string) string {
	if c, ok := colorMp[color]; ok {
		return c
	}
	return newColor(color)
}
