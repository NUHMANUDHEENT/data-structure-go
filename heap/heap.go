package main

import "fmt"

// MinHeap represents a min-heap data structure
type MinHeap struct {
	elements []int
}

func (h *MinHeap) Insert(value int) {
	h.elements = append(h.elements, value)
	h.HeapifyUp(len(h.elements) - 1)
}
func (h *MinHeap) HeapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.elements[i] >= h.elements[parent] {
			break
		}
		h.elements[i], h.elements[parent] = h.elements[parent], h.elements[i]
		i = parent
	}
}
func (h *MinHeap) Delete(value int) {
	if len(h.elements) == 0 {
		fmt.Println("heap is nill")
	}
	var index int
	for i, v := range h.elements {
		if value == v {
			index = i
			break
		}
	}
	h.elements[index] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.HeapifyDown(index)
}
func (h *MinHeap) HeapifyDown(i int) {
	lastIndex := len(h.elements) - 1
	for {
		leftChild := i*2 + 1
		rightChild := i*2 + 2
		smallest := i
		if leftChild <= lastIndex && h.elements[leftChild] < h.elements[smallest] {
			smallest = leftChild
		}
		if rightChild <= lastIndex && h.elements[rightChild] < h.elements[smallest] {
			smallest = rightChild
		}
		if smallest == i {
			break
		}
		h.elements[i], h.elements[smallest] = h.elements[smallest], h.elements[i]
		i = smallest
	}
}
func main() {
	h := &MinHeap{}
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)
	h.Insert(4)
	h.Insert(2)
	h.Delete(2)
	fmt.Println("Heap elements:", h.elements)

}
func (h *MinHeap) heapy(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.elements[parent] < h.elements[i] {
			break
		}
		h.elements[parent], h.elements[i] = h.elements[i], h.elements[parent]
		i = parent
	}

}
