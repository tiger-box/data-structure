package main

import (
	"fmt"
)

type intQueue struct {
	queue []int
}

func (q *intQueue) Enqueue(value int) {
	(*q).queue = append(q.queue, value)
}

func (q *intQueue) Dequeue() int {
	element := q.queue[0]
	q.queue = q.queue[1:]

	return element
}

func (q *intQueue) Peek() int {
	element := q.queue[0]

	return element
}

func (q *intQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func main() {
	slice := intQueue{}
	slice.Enqueue(1)
	slice.Enqueue(2)
	slice.Enqueue(3)
	slice.Enqueue(4)

	fmt.Println(slice.Dequeue())
	fmt.Println(slice.queue)
	fmt.Println(slice.Peek())
	fmt.Println(slice.IsEmpty())
}
