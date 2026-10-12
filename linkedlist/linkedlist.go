package main

import (
	"fmt"
)

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	size int
}

func NewList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) PushFront(value int) {
	head := &LinkedListNode{
		value: value,
		next:  l.head,
	}
	l.head = head
	l.size++
}

func (l *LinkedList) PushBack(value int) {
	current := l.head
	prev := current

	for current != nil {
		prev = current
		current = current.next

	}
	prev.next = &LinkedListNode{
		value: value,
		next:  nil,
	}

	l.size++
}
func (l *LinkedList) Insert(index, value int) error {
	if index > l.size || index < 0 {
		return fmt.Errorf("ошибка %d", index)
	}

	if index == 0 {
		head := &LinkedListNode{
			value: value,
			next:  l.head,
		}
		l.head = head
		l.size++
		return nil
	}

	current := l.head
	prev := current

	for i := 0; current != nil && i < index; i++ {
		prev = current
		current = current.next
	}

	node := &LinkedListNode{value: value, next: current}
	prev.next = node

	l.size++

	return nil
}
func (l *LinkedList) Delete(index int) (int, error) {

	if index > l.size || index < 0 {
		return 0, fmt.Errorf("index out of range %d", index)
	}

	current := l.head
	prev := current

	if index == 0 {
		value := l.head.value
		current = current.next
		l.head = current
		l.size--
		return value, nil
	}

	for i := 0; i < index; i++ {
		prev = current
		current = current.next

	}
	prev.next = current.next
	l.size--

	return current.value, nil
}
func (l *LinkedList) Find(value int) int {

	current := l.head
	for i := 0; i < l.size; i++ {
		if current.value == value {
			return i
		}
		current = current.next
	}
	return -1
}

func PrintList(l *LinkedList) {
	if l.head == nil {
		return
	}

	current := l.head
	i := 0
	for current != nil {
		fmt.Println(i, ": value ", current.value, " next", current.next)

		i++
		current = current.next
	}
	fmt.Println()
}

func main() {
	l := NewList()
	l.PushFront(10)
	l.PushFront(20)
	l.PushFront(45)
	l.PushBack(55)
	l.Insert(2, 7)
	PrintList(l)
	value, err := l.Delete(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
	PrintList(l)
	fmt.Println(l.Find(7))

}
