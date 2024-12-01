package main

import (
	"errors"
	"fmt"
)

type stack struct {
	data []int
}

func (s *stack) Push(value int) {
	s.data = append(s.data, value)
}
func (s *stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack is empty")
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last, nil
}
func (s *stack) Pull() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}
func main() {
	s := &stack{[]int{}}
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println((s.Pull()))
	fmt.Println((s.Pop()))
	s.print()

}
func (s *stack) print() {
	fmt.Println(s.data)
}
