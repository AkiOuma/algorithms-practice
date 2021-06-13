package stack

import (
	"errors"
	"fmt"
)

type ArrStack struct {
	data []interface{}
}

// 出栈
func (s *ArrStack) Pop() (popData interface{}, err error) {
	if len(s.data) == 0 {
		err = errors.New("error: stack is empty")
		return
	}
	index := len(s.data) - 1
	popData = s.data[index]
	s.data = s.data[:index]
	return
}

// 入栈
func (s *ArrStack) Push(data interface{}) {
	s.data = append(s.data, data)
}

func (s *ArrStack) PrintStack() {
	fmt.Println(s.data)
}

// 判断栈是否为空
func (s *ArrStack) IsEmpty() (result bool) {
	result = false
	if len(s.data) == 0 {
		result = true
	}
	return
}

// 获取栈顶信息
func (s *ArrStack) Top() (data interface{}, err error) {
	if s.IsEmpty() {
		err = errors.New("error: stack is empty now")
		return
	}
	size := len(s.data)
	data = s.data[size-1]
	return
}
