// Package treesort implements a binary tree sort algorithm.
package treesort

import (
	"cmp"
	"fmt"
)

// Tree represents a binary search tree node that holds comparable values.
type Tree[T cmp.Ordered] struct {
	value       T        // The value stored in this node
	left, right *Tree[T] // Left and right child nodes
}

// NewTree creates and returns a new Tree node with the given value.
// T must be an ordered type that can be compared.
func NewTree[T cmp.Ordered](value T) *Tree[T] {
	return &Tree[T]{value: value}
}

// PrintTree recursively prints the binary tree structure in a rotated fashion.
// It displays the right subtree first, then the current node value indented by its level,
// followed by the left subtree. This helps visualize the tree hierarchy.
// The function is generic over ordered types T.
// Parameters:
//   - t: pointer to the root node of the tree to print.
//   - level: current depth level, used for indentation.
func PrintTree[T cmp.Ordered](t *Tree[T], level int) {
	if t != nil {
		PrintTree(t.right, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("   ")
		}
		fmt.Println(t.value)
		PrintTree(t.left, level+1)
	}
}

// Sort sorts the provided slice in ascending order using a binary tree sort algorithm.
// The sorting is done in-place, modifying the original slice.
// Time complexity: O(n log n) average case, O(n^2) worst case.
// Space complexity: O(n) to store the binary tree.
func Sort[T cmp.Ordered](values []T) {
	var root *Tree[T]
	for _, v := range values {
		root = Add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues performs an in-order traversal of the tree and appends values
// to the provided slice in sorted order. It returns the resulting slice.
// The in-order traversal ensures values are processed in ascending order.
func appendValues[T cmp.Ordered](values []T, t *Tree[T]) []T {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add inserts a new value into the binary search tree.
// If the tree is nil, creates a new tree with the given value.
// Otherwise, recursively inserts the value in the appropriate subtree
// to maintain binary search tree properties.
// Returns the root of the modified tree.
func Add[T cmp.Ordered](t *Tree[T], value T) *Tree[T] {
	if t == nil {
		// Equivalent to return &Tree{value: value}.
		t = new(Tree[T])
		t.value = value
		return t
	}

	if value < t.value {
		t.left = Add(t.left, value)
	} else {
		t.right = Add(t.right, value)
	}
	return t
}
