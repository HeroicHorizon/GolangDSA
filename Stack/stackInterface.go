package main

import (
	"errors"
	"fmt"
)

type StackImpl interface {
	Push(int)
	Pop() (int, error)
	IsEmpty() bool
	Print()
	Size() int
	Top() (int, error)
}

type StackStruct struct {
	items []int
}

func NewStack() StackImpl {
	return &StackStruct{
		items: []int{},
	}
}

func (stack *StackStruct) Push(data int) {
	stack.items = append(stack.items, data)
}

func (stack *StackStruct) Pop() (int, error) {
	if len(stack.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	item := stack.items[len(stack.items)-1]        // Get the last item
	stack.items = stack.items[:len(stack.items)-1] // Remove the last item
	return item, nil
}

func (stack StackStruct) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack StackStruct) Size() int {
	return len(stack.items)
}

func (stack StackStruct) Top() (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return stack.items[len(stack.items)-1], nil
}

func (stack StackStruct) Print() {
	for _, item := range stack.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
