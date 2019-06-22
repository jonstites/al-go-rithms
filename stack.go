package main

import (
	"errors"
)

// Stack implements a standard stack data structure
type Stack struct {
	data []int
}

func newStack() Stack {
	var data [0]int
	return Stack {data: data[:]}
}

func (stack *Stack) size() int {
	return len(stack.data)
}

func (stack *Stack) isEmpty() bool {
	return stack.size() == 0
}

func (stack *Stack) peek() (int, error) {
	
	if stack.isEmpty() {
		return 0, errors.New("No items to peek")
	}

	item := stack.data[stack.size() - 1]
	return item, nil
}

func (stack *Stack) pop() (int, error) {	
	item, err := stack.peek()
	if err != nil {
		return 0, errors.New("No items to pop")
	}
	stack.data = stack.data[:stack.size() - 1]
	return item, nil
}

func (stack *Stack) push(i ...int) {
	stack.data = append(stack.data, i...)
}


