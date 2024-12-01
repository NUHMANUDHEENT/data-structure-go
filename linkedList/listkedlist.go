package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type linkedList struct {
	Head *Node
}

func (l *linkedList) Push(value int) {
	newNode := &Node{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func (l *linkedList) InsertAtBegining(value int) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}
func (l *linkedList) Delete(value int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
func (l *linkedList) Search(value int) bool {
	if l.Head == nil {
		return false
	}
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func main() {
	l := &linkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	// l.printList()
	// l.InsertAtBegining(0)
	l.NodeDelete(1)
	// fmt.Println("search value", l.Search(1))
	l.printList()

}

func (l *linkedList) printList() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *linkedList) insert(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func (l *linkedList) NodeInsert(value int) {
	newNode := &Node{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func (l *linkedList) NodeDelete(value int) {
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	prev := l.Head
	for current != nil {
		if current.Value == value {
			prev.Next = current.Next
		}
		prev = current
		current = current.Next
	}
}
