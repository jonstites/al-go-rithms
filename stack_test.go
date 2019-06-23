package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	stack := NewStack()
	if stack.IsEmpty() == false {
		t.Error("Expected empty stack")
	}

	stack.Push(0)
	if stack.IsEmpty() == true {
		t.Error("Expected non-empty stack")
	}
}

func TestSize(t *testing.T) {
	stack := NewStack()
	if stack.Size() != 0 {
		t.Error("Expected 0 size stack")
	}

	stack.Push(0, 2, 0)
	if stack.Size() != 3 {
		t.Error("Expected 3 size stack")
	}

	stack.Pop()
	if stack.Size() != 2 {
		t.Error("Expected 2 size stack")
	}
}

func TestPushPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1, 2, 3, "abc")

	var expected interface{}
	expected = "abc"
	actual, _ := stack.Pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	expected = 3
	actual, _ = stack.Pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	expected = 2
	actual, _ = stack.Pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	expected = 1
	actual, _ = stack.Pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected error from pop of empty stack")
	}
}
