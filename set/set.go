package main

import (
	"fmt"
)

type intSet struct {
	set map[int]struct{}
}

func (s *intSet) add(value int) {
	s.set[value] = struct{}{}
}

func (s *intSet) Remove(value int) error {

	if _, ok := s.set[value]; !ok {
		return fmt.Errorf("No such value was found.")
	}

	newSet := make(map[int]struct{}, len(s.set)-1)

	for key := range s.set {
		if key != value {
			newSet[key] = struct{}{}
		}
	}

	s.set = newSet

	return nil
}

func Constructor() *intSet {

	return &intSet{
		set: make(map[int]struct{}),
	}
}

func (s *intSet) Contains(value int) bool {
	if _, ok := s.set[value]; !ok {
		return false
	}

	return true
}

func (s *intSet) size() int {
	return len(s.set)
}

func main() {

	set := Constructor()

	set.add(10)
	set.add(20)
	err := set.Remove(40)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(set)
	}

	err1 := set.Contains(100)
	if err != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(set)
	}

	fmt.Println(set.size())

}
