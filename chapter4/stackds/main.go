package main

type Stack[T any] struct {
	elements []any
	size     int
}

// Push adds an element to the top of the stack.
// It panics if the stack has reached its maximum size capacity.
func (s *Stack[T]) Push(elem T) {
	if len(s.elements) == s.size {
		panic("stack overflow")
	}
	s.elements = append(s.elements, elem)
}

// Pop removes and returns the top element from the stack.
// It panics if the stack is empty (stack underflow).
func (s *Stack[T]) Pop() T {
	// Check if the stack is empty
	if len(s.elements) == 0 {
		panic("stack underflow")
	}

	// Get the top element
	top := s.elements[len(s.elements)-1]

	// Remove the top element from the stack
	s.elements = s.elements[:len(s.elements)-1]

	// Return the top element
	return top.(T)
}

// IsEmpty returns true if the stack contains no elements, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {

}
