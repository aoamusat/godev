package stackds

import "testing"

func TestStackPushPop(t *testing.T) {
	stack := Stack[int]{size: 3}

	// Test pushing elements onto the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test stack overflow
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for stack overflow, but did not occur")
		}
	}()
	stack.Push(4) // This should panic
}

func TestStackPopEmpty(t *testing.T) {
	stack := Stack[int]{size: 3}

	// Test popping from an empty stack
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for stack underflow, but did not occur")
		}
	}()
	stack.Pop() // This should panic
}

func TestStackPushPopSequence(t *testing.T) {
	stack := Stack[string]{size: 2}

	stack.Push("first")
	stack.Push("second")

	if stack.Pop() != "second" {
		t.Errorf("expected 'second', got different value")
	}

	if stack.Pop() != "first" {
		t.Errorf("expected 'first', got different value")
	}

	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty after popping all elements")
	}
}
