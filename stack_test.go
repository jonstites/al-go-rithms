package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	stack := newStack()
	if stack.isEmpty() == false {
		t.Error("Expected empty stack")
	}

	stack.push(0)
	if stack.isEmpty() == true {
		t.Error("Expected non-empty stack")
	}
}

func TestSize(t *testing.T) {
	stack := newStack()
	if stack.size() != 0 {
		t.Error("Expected 0 size stack")
	}

	stack.push(0, 2, 0)
	if stack.size() != 3 {
		t.Error("Expected 3 size stack")
	}

	stack.pop()
	if stack.size() != 2 {
		t.Error("Expected 2 size stack")
	}
}

func TestPushPop(t *testing.T) {
	stack := newStack()
	stack.push(1, 2, 3, "abc")

	var expected interface{}
	expected = "abc"
	actual, _ := stack.pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	expected = 3
	actual, _ = stack.pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	expected = 2
	actual, _ = stack.pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	expected = 1
	actual, _ = stack.pop()
	if actual != expected {
		t.Error(expected, actual)
	}

	_, err := stack.pop()
	if err == nil {
		t.Error("Expected error from pop of empty stack")
	}
}
