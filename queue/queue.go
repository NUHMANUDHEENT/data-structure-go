package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	element []int
}

func (q *Queue) Enqueue(value int) {
	q.element = append(q.element, value)
}
func (q *Queue) Dequeue() (int, error) {
	if len(q.element) == 0 {
		return 0, errors.New("queue is empty")
	}
	first := q.element[0]
	q.element = q.element[1:]
	return first, nil
}
func main() {
	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	dequeued, _ := queue.Dequeue()
	fmt.Println("Dequeued element:", dequeued)
	fmt.Println("elements",queue.element)
}
