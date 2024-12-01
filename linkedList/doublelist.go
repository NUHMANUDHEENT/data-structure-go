package main

import (
	"fmt"
)

type DoubleNode struct {
	Prev  *DoubleNode
	Next  *DoubleNode
	Value int
}
type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

func (dl *DoubleLinkedList) DoubleListInsert(value int) {
	newNode := &DoubleNode{Value: value, Next: nil, Prev: nil}
	if dl.Tail == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}
	dl.Tail.Next = newNode
	newNode.Prev = dl.Tail
	dl.Tail = newNode

}
func (dl *DoubleLinkedList) Delete(value int) {
	if dl.Head == nil {
		return
	}
	current := dl.Head
	for current != nil && current.Value != value {
		current = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
		current.Prev.Next = current.Next
	} else {
		dl.Tail = current.Prev
	}
}

func main() {
	dll := &DoubleLinkedList{}
	dll.DoubleListInsert(1)
	dll.DoubleListInsert(2)
	dll.DoubleListInsert(3)
	dll.DoubleListInsert(4)
	dll.Delete(4)
	dll.Traversal()
	// fmt.Println("cycle ", dll.middlevalue())

}
func (dl *DoubleLinkedList) Traversal() {
	for current := dl.Head; current != nil; current = current.Next {
		fmt.Println(current.Value)
	}
}
func (dl *DoubleLinkedList) hasCycle() bool {
	if dl.Head == nil || dl.Head.Next == nil {
		return false
	}
	slow := dl.Head
	fast := dl.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}
func (dl *DoubleLinkedList) middlevalue() int {
	if dl.Head == nil {
		return 0
	}
	slow, fast := dl.Head, dl.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Value
}
func (dl *DoubleLinkedList) NodeDoubleLinkedListInsert(value int) {
	newNode := &DoubleNode{Value: value, Next: nil, Prev: nil}
	if dl.Tail == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}
	dl.Tail.Next = newNode
	newNode.Prev = dl.Tail
	dl.Tail = newNode
}
func (dl *DoubleLinkedList) NodeDoubleDelete(value int) {
	if dl.Head.Value == value {
		dl.Head = dl.Head.Next
		dl.Head.Prev = nil
	}
	current := dl.Head
	for current.Next != nil {
		if current.Value == value && current.Next != nil {
			fmt.Println("gi")
			current.Next.Prev = current.Prev
			current.Prev.Next = current.Next
			return
		}
		current = current.Next
	}
	dl.Tail = current.Prev
}
