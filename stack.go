package main

import (
	"errors"
)

// Stack implements a standard stack data structure
// Unfortunately, due to Go's lack of generics, there cannot
// be a compile-time check that all items in the Stack are
// the same type.
type Stack struct {
	data []interface{}
}

// NewStack creates a new Stack
func NewStack() Stack {
	return Stack{}
}

// Size returns the number of elements currently in the Stack
func (stack *Stack) Size() int {
	return len(stack.data)
}

// IsEmpty returns true iff there are no elements in the Stack
func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

// Peek returns the last item without removing it
func (stack *Stack) Peek() (interface{}, error) {

	if stack.IsEmpty() {
		return 0, errors.New("No items to peek")
	}

	item := stack.data[stack.Size()-1]
	return item, nil
}

// Pop removes and returns the last item
func (stack *Stack) Pop() (interface{}, error) {
	item, err := stack.Peek()
	if err != nil {
		return 0, errors.New("No items to pop")
	}
	stack.data = stack.data[:stack.Size()-1]
	return item, nil
}

// Push adds a new item or items to the Stack
func (stack *Stack) Push(i ...interface{}) {
	// this may copy data in worst case, but amoritized time of Pop should be O(1)
	// since in Go, the length of the slice does not have to be equal to capacity
	stack.data = append(stack.data, i...)
}
