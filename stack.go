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

func newStack() Stack {
	return Stack{}
}

func (stack *Stack) size() int {
	return len(stack.data)
}

func (stack *Stack) isEmpty() bool {
	return stack.size() == 0
}

func (stack *Stack) peek() (interface{}, error) {

	if stack.isEmpty() {
		return 0, errors.New("No items to peek")
	}

	item := stack.data[stack.size()-1]
	return item, nil
}

func (stack *Stack) pop() (interface{}, error) {
	item, err := stack.peek()
	if err != nil {
		return 0, errors.New("No items to pop")
	}
	stack.data = stack.data[:stack.size()-1]
	return item, nil
}

func (stack *Stack) push(i ...interface{}) {
	stack.data = append(stack.data, i...)
}
