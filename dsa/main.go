// Package dsa implements a generic thread-safe linked list data structure
package dsa

import (
	"fmt"
	"sync"
)

// node represents a single node in the linked list
// It contains a value of generic type T and a pointer to the next node
type node[T any] struct {
	value T
	next  *node[T]
}

// LinkedList implements a generic singly-linked list that can be thread-safe
// T must be comparable to support operations like Remove
type LinkedList[T comparable] struct {
	head *node[T]      // Points to first node
	tail *node[T]      // Points to last node
	len  int           // Number of nodes in list
	mu   *sync.RWMutex // Mutex for thread safety, nil = not thread-safe
}

// New creates a new empty LinkedList
// threadSafe parameter determines if the list should use mutex locking
func New[T comparable](threadSafe bool) *LinkedList[T] {
	var mu *sync.RWMutex
	if threadSafe {
		mu = &sync.RWMutex{}
	}
	return &LinkedList[T]{mu: mu}
}

// lock acquires write lock if list is thread-safe
func (l *LinkedList[T]) lock() {
	if l.mu != nil {
		l.mu.Lock()
	}
}

// unlock releases write lock if list is thread-safe
func (l *LinkedList[T]) unlock() {
	if l.mu != nil {
		l.mu.Unlock()
	}
}

// rlock acquires read lock if list is thread-safe
func (l *LinkedList[T]) rlock() {
	if l.mu != nil {
		l.mu.RLock()
	}
}

// runlock releases read lock if list is thread-safe
func (l *LinkedList[T]) runlock() {
	if l.mu != nil {
		l.mu.RUnlock()
	}
}

func (l *LinkedList[T]) Print() {
	if l.len != 0 {
		for n := l.head; n != nil; n = n.next {
			fmt.Printf("%v -> ", n.value)
		}
		fmt.Println()
	} else {
		fmt.Printf("List is Empty")
	}
}

// Append adds a new value to the end of the list
func (l *LinkedList[T]) Append(value T) {
	l.lock()
	defer l.unlock()

	n := &node[T]{value: value}

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.len++
}

// Prepend adds a new value to the start of the list
func (l *LinkedList[T]) Prepend(value T) {
	l.lock()
	defer l.unlock()

	n := &node[T]{value: value, next: l.head}
	l.head = n

	if l.tail == nil {
		l.tail = n
	}
	l.len++
}

// Len returns the current length of the list
func (l *LinkedList[T]) Len() int {
	l.rlock()
	defer l.runlock()
	return l.len
}

// Remove removes the first occurrence of value from the list
// Returns true if value was found and removed, false otherwise
func (l *LinkedList[T]) Remove(value T) bool {
	l.lock()
	defer l.unlock()

	var prev *node[T]
	curr := l.head

	for curr != nil {
		if curr.value == value {
			if prev == nil {
				l.head = curr.next
			} else {
				prev.next = curr.next
			}
			if curr == l.tail {
				l.tail = prev
			}
			l.len--
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

// RemoveAt removes the node at the specified index
// Returns the removed value and true if successful, zero value and false if index invalid
func (l *LinkedList[T]) RemoveAt(index int) (T, bool) {
	l.lock()
	defer l.unlock()

	var zero T
	if index < 0 || index >= l.len {
		return zero, false
	}

	var prev *node[T]
	curr := l.head

	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}

	if prev == nil {
		l.head = curr.next
	} else {
		prev.next = curr.next
	}

	if curr == l.tail {
		l.tail = prev
	}

	l.len--
	return curr.value, true
}

// Reverse reverses the order of nodes in the list in-place
func (l *LinkedList[T]) Reverse() {
	l.lock()
	defer l.unlock()

	var prev *node[T]
	curr := l.head
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev
}

// Map creates a new list by applying the provided function to each element
// The function transforms values of type T to values of type R
func Map[T comparable, R comparable](
	l *LinkedList[T],
	fn func(T) R,
) *LinkedList[R] {
	result := New[R](false)

	l.rlock()
	defer l.runlock()

	for n := l.head; n != nil; n = n.next {
		result.Append(fn(n.value))
	}
	return result
}

// Filter creates a new list containing only elements that satisfy the predicate
func (l *LinkedList[T]) Filter(predicate func(T) bool) *LinkedList[T] {
	result := New[T](false)

	l.rlock()
	defer l.runlock()

	for n := l.head; n != nil; n = n.next {
		if predicate(n.value) {
			result.Append(n.value)
		}
	}
	return result
}

// Iterator provides sequential access to list elements
type Iterator[T any] struct {
	current *node[T]
}

// Iterator returns a new Iterator positioned at the start of the list
func (l *LinkedList[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{current: l.head}
}

// Next returns the next value in the iteration
// Returns the value and true if there was a next element, zero value and false if at end
func (it *Iterator[T]) Next() (T, bool) {
	if it.current == nil {
		var zero T
		return zero, false
	}
	val := it.current.value
	it.current = it.current.next
	return val, true
}

// Example usage of the LinkedList implementation
func main() {
	list := New[int](true)

	list.Append(1)
	list.Append(2)
	list.Prepend(0)

	list.Remove(2)
	list.Reverse()

	for it := list.Iterator(); ; {
		v, ok := it.Next()
		if !ok {
			break
		}
		println(v)
	}

	doubled := Map(list, func(v int) int { return v * 2 })
	println("Doubled length:", doubled.Len())
}
