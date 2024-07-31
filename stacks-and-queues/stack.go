package main

import "fmt"

type Stacks struct{
	items []int
}

func (s *Stacks) push(n int) {
	s.items = append(s.items, n)
}

func (s *Stacks) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]

	return toRemove
}

func main(){
	myStack := Stacks{}

	myStack.push(44)
	myStack.push(55)
	myStack.push(66)

	fmt.Println(myStack)

	myStack.pop()

	fmt.Println(myStack)
}