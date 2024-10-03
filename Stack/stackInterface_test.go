package main

import "testing"

func TestStackUsingInteface(t *testing.T) {
	stack := NewStack()

	// Test initial state
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
	if size := stack.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test pushing items
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()

	if size := stack.Size(); size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	// Test top
	top, err := stack.Top()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if top != 3 {
		t.Errorf("Expected top item 3, got %d", top)
	}

	// Test popping items
	item, err := stack.Pop()
	if err != nil || item != 3 {
		t.Errorf("Expected to pop 3, got %d (error: %v)", item, err)
	}

	item, err = stack.Pop()
	if err != nil || item != 2 {
		t.Errorf("Expected to pop 2, got %d (error: %v)", item, err)
	}

	item, err = stack.Pop()
	if err != nil || item != 1 {
		t.Errorf("Expected to pop 1, got %d (error: %v)", item, err)
	}

	// Test popping from an empty stack
	_, err = stack.Pop()
	if err == nil {
		t.Error("Expected error when popping from an empty stack, got none")
	}

	// Test top on an empty stack
	_, err = stack.Top()
	if err == nil {
		t.Error("Expected error when peeking on an empty stack, got none")
	}

	// Final size check
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all items")
	}
	if size := stack.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

}
