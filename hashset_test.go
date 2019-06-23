package main

import (
	"testing"
)

func TestHashSetAddRemove(t *testing.T) {
	set := newHashSet()

	if !set.add("abc") {
		t.Error("Expected value not to be in set")
	}

	if set.add("abc") {
		t.Error("Expected value to be in set")
	}

	if !set.remove("abc") {
		t.Error("Expected value to be removed from set")
	}

	if set.remove("abc") {
		t.Error("Expected value not to be removed from set")
	}
}

func TestHashSetContains(t *testing.T) {
	set := newHashSet()

	if set.contains("abc") {
		t.Error("Expected set not to contain value")
	}

	set.add("abc")
	if !set.contains("abc") {
		t.Error("Expected set to contain value")
	}

	if set.contains("123") {
		t.Error("Expected set not to contain value")
	}

	if set.contains(123) {
		t.Error("Expected set not to contain value")
	}

	set.remove("abc")
	if set.contains("abc") {
		t.Error("Expected set not to contain value")
	}
}

func TestHashSetSize(t *testing.T) {
	set := newHashSet()
	
	if set.size() != 0 {
		t.Error("Expected empty set")
	}

	if !set.isEmpty() {
		t.Error("Expected empty set")
	}

	set.add("abc")
	if set.isEmpty() {
		t.Error("Expected non-empty set")
	}

	if set.size() != 1 {
		t.Error("Expected size of 1")
	}

	set.add(123)
	set.add(32)
	set.remove("abc")

	if set.isEmpty() {
		t.Error("Expected non-empty set")
	}

	if set.size() != 2 {
		t.Error("Expected size of 2")
	}

	set.clear()

	if set.size() != 0 {
		t.Error("Expected empty set")
	}

	if !set.isEmpty() {
		t.Error("Expected empty set")
	}

}
