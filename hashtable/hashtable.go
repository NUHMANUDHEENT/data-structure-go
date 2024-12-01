package main

import "fmt"

type Node struct {
	key   string
	value int
	next  *Node
}
type Hashtable struct {
	bucket []*Node
	size   int
}

func NewHashTable() *Hashtable {
	return &Hashtable{
		bucket: make([]*Node, 10),
		size:   10,
	}
}
func (h *Hashtable) hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % h.size

}
func (h *Hashtable) Insert(key string, value int) {
	index := h.hash(key)
	if h.bucket[index] == nil {
		h.bucket[index] = &Node{key: key, value: value}
		return
	}
	current := h.bucket[index]
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{key: key, value: value}

}
func (h *Hashtable) get(key string) int {
	index := h.hash(key)
	if h.bucket[index].key == key {
		return h.bucket[index].value
	}
	currenct := h.bucket[index]
	for currenct != nil {
		if currenct.key == key {
			return currenct.value
		}
		currenct = currenct.next
	}
	return 0
}
func main() {
	h := NewHashTable()
	h.Insert("apple", 10)
	h.Insert("orange", 11)
	fmt.Println(h.get("apple"))
}
