// 桥接（Bridge）是用于把抽象化与实现化解耦，使得二者可以独立变化。这种类型的设计模式属于结构型模式，它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。

// 这种模式涉及到一个作为桥接的接口，使得实体类的功能独立于接口实现类。这两种类型的类可被结构化改变而互不影响。
package main

import "fmt"

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

type printer interface {
	printFile()
}

type epson struct {
}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()
	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}
