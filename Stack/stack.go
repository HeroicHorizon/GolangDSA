package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

// Pop removes and returns the top item of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]    // Get the last item
	s.items = s.items[:len(s.items)-1] // Remove the last item
	return item, nil
}

// Top returns the top item without removing it
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Print print the stack elements
func (s Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

// Size returns the current size of the stack
func (s Stack) Size() int {
	return len(s.items)
}
