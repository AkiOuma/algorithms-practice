package main

import (
	"algorithm/stack"
	"fmt"
)

func main() {
	var stack1 stack.LinkStack
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)
	stack1.PrintStack()
	stack1.Pop()
	stack1.Pop()
	stack1.Pop()
	_, err := stack1.Pop()
	if err != nil {
		fmt.Println(err)
	}
	stack1.PrintStack()
}
