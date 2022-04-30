// 备忘录模式（Memento Pattern）保存一个对象的某个状态，以便在适当的时候恢复对象。备忘录模式属于行为型模式。
package main

import "fmt"

type originator struct {
	state string
}

func (e *originator) createMemento() *memento {
	return &memento{state: e.state}
}

func (e *originator) restoreState(m *memento) {
	e.state = m.getSavedState()
}

func (e *originator) setState(state string) {
	e.state = state
}

func (e *originator) getState() string {
	return e.state
}

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMenento(index int) *memento {
	return c.mementoArray[index]
}

func main() {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	originator.restoreState(caretaker.getMenento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())
	originator.restoreState(caretaker.getMenento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
