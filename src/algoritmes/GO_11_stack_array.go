package main

import (
	"fmt"
)

type IntNode struct {
	Value int
	Next  *IntNode
}

func (n IntNode) GetNext() *IntNode {
	return n.Next
}
func NewIntNode(value int) *IntNode {
	return &IntNode{value, nil}
}

// ----------------------------------
func NewIntList() *IntList {
	return &IntList{0, nil}
}

type IntList struct {
	size int
	Head *IntNode
}

func (l IntList) Size() int {
	return l.size
}
func (l IntList) Get(index int) (*IntNode, error) {
	if index < 0 || index > l.Size() {
		return nil, fmt.Errorf("Неверный индекс списка")
	}
	node := l.Head
	for i := 1; i <= index; i++ {
		node = node.Next
	}
	return node, nil
}
func (l *IntList) Set(el int, index int) error {
	if index < 0 || index > l.Size() {
		return fmt.Errorf("Неверный индекс списка")
	}
	node := l.Head
	for i := 1; i <= index; i++ {
		node = node.Next
	}
	node.Value = el
	return nil
}

func (l *IntList) Add(el int) {
	newNode := NewIntNode(el)
	if l.Head == nil {
		l.Head = newNode
	} else {
		node := l.Head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = newNode
	}
	l.size++
}
func (l *IntList) Insert(el int, index int) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("Неверный индекс списка")
	}
	newNode := NewIntNode(el)
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
	} else {
		node, err := l.Get(index - 1)
		if err != nil {
			return err
		}
		newNode.Next = node.Next
		node.Next = newNode
	}
	l.size++
	return nil
}

func (l *IntList) Remove(index int) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("Неверный индекс списка")
	}
	node, err := l.Get(index - 1)
	if err != nil {
		return err
	}
	node.Next = node.GetNext().GetNext()
	l.size--
	return nil
}
func (l IntList) Print() {
	node := l.Head
	if node != nil {
		for node != nil {
			fmt.Printf("%d\t", node.Value)
			node = node.Next
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Список пуст!")
	}
}
func main() {
	list := NewIntList()
	list.Print()
	list.Add(1)
	list.Add(2)
	list.Insert(0, 0)
	list.Insert(0, 0)
	fmt.Println(list.Size())
	list.Print()
	list.Remove(1)
	list.Remove(1)
	list.Print()
	fmt.Println(list.Size())
}
