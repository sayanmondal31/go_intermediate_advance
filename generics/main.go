package main

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

func main() {
	// x, y := 1, 2
	// fmt.Println("Before swap: ", x, y)
	// x, y = swap(1, 2)
	// fmt.Println("After swap: ", x, y)

}
