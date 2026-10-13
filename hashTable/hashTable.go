package main

import "fmt"

type node struct {
	key   string
	value int
	next  *node
}

type hashTable struct {
	buckets []*node
	size    int
}

func (ht *hashTable) hashFunc(key string) int {

	if len(ht.buckets) == 0 {
		fmt.Println("empty")
	}

	sum := 0

	for _, intString := range key {
		sum = int(intString) + sum
	}

	return sum % len(ht.buckets)
}

func (ht *hashTable) set(key string, value int) {
	index := ht.hashFunc(key)
	head := ht.buckets[index]
	current := head

	if head == nil {
		head = &node{
			key:   key,
			value: value}

		ht.size++
		ht.buckets[index] = head

		return
	}

	for current != nil {
		if current.key == key {
			current.value = value
			return
		}

		if current.next == nil {
			break
		}

		current = current.next
	}

	current.next = &node{key: key, value: value}
	ht.size++

}

func (ht *hashTable) Get(key string) (int, bool) {
	index := ht.hashFunc(key)
	head := ht.buckets[index]

	if head == nil {
		return 0, false
	}

	current := head

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}

	return 0, false
}

func (ht *hashTable) delete(key string) {
	index := ht.hashFunc(key)
	head := ht.buckets[index]

	current := head.next
	prev := head

	if head.key == key {
		ht.buckets[index] = head
		ht.size--
		return
	}

	for current != nil {
		if current.key == key {
			prev.next = current.next
			ht.size--
			return
		}

		prev = current
		current = current.next

	}
}

func Constructor() *hashTable {
	return &hashTable{
		buckets: make([]*node, 10),
		size:    0,
	}

}

func main() {

	ht := Constructor()

	ht.set("apple", 100)
	ht.set("banana", 200)
	ht.set("orange", 300)

	fmt.Println(ht)
	fmt.Println(ht.Get("banana"))

}
