package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(elemet T) {
	s.elements = append(s.elements, elemet)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Print("Stack elements: ")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	// x, y := 1, 2
	// fmt.Println("Before swap: ", x, y)
	// x, y = swap(1, 2)
	// fmt.Println("After swap: ", x, y)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	intStack.printAll()

	intStack.pop()
	intStack.printAll()

	stringStack := Stack[string]{}

	stringStack.push("hello ")
	stringStack.push("world ")
	stringStack.push("Sayan ")

	stringStack.printAll()
	fmt.Println(stringStack.pop())
	stringStack.printAll()

}
