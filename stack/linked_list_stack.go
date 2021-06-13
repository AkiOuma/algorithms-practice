package stack

import (
	"errors"
	"fmt"
)

type LinkStackNode struct {
	Data interface{}
	Last *LinkStackNode
}

type LinkStack struct {
	Top *LinkStackNode
}

func (l *LinkStack) Push(data interface{}) {
	node := LinkStackNode{Data: data}
	if l.Top != nil {
		node.Last = l.Top
	}
	l.Top = &node
}

func (l *LinkStack) Pop() (item interface{}, err error) {
	if l.Top == nil {
		err = errors.New("error: stack is empty")
		return
	}
	item = l.Top.Data
	l.Top = l.Top.Last
	return
}

func (l *LinkStack) PrintStack() {
	arr := make([]interface{}, 0)
	current := l.Top
	for current != nil {
		arr = append(arr, current.Data)
		current = current.Last
	}
	fmt.Println(ReverseSlice(arr))
}

func ReverseSlice(source []interface{}) []interface{} {
	size := len(source)
	target := make([]interface{}, 0, size)
	i := size - 1
	for i >= 0 {
		target = append(target, source[i])
		i--
	}
	return target
}

func (l *LinkStack) IsEmpty() (result bool) {
	result = false
	if l.Top == nil {
		result = true
	}
	return
}
