package stack

import "testing"

func TestRPN(t *testing.T) {
	var stack LinkStack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	value, err := stack.Pop()
	if err != nil || value.(int) != 30 {
		t.Error("error: Linked List Stack [method: Pop] did not pass the test")
	}
	stack.Pop()
	isEmpty := stack.IsEmpty()
	if isEmpty != false {
		t.Error("error: Linked List [method: IsEmpty] did not pass the test")
	}
	stack.Pop()
	isEmpty = stack.IsEmpty()
	if isEmpty != true {
		t.Error("error: Linked List [method: IsEmpty] did not pass the test")
	}
}
