package main

import "fmt"

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Current stack :")
	stack.Print()
	fmt.Println("Current stack size:", stack.Size()) // Output: 3

	top, _ := stack.Top()
	fmt.Println("Top item:", top) // Output: 3

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println("Popped item:", item)
	}

	fmt.Println("Current stack size after popping all items:", stack.Size()) // Output: 0
}
