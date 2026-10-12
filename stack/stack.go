package main

import (
	"fmt"
)

type intStack struct {
	stack []int
}

func (s *intStack) Push(value int) {
	(*s).stack = append(s.stack, value)
}

func (s *intStack) Pop() int {
	element := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return element
}

func (s *intStack) Peek() int {
	element := s.stack[len(s.stack)-1]

	return element
}

func (s *intStack) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}

	return false
}

func main() {
	stack := intStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Pop())
	fmt.Println(stack.stack)
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())
}
