package stack

import (
	"errors"
	"fmt"
)

type ArrStack struct {
	data []int
}

func (s *ArrStack) Pop() (popData int, err error) {
	if len(s.data) == 0 {
		err = errors.New("error: stack is empty")
		return
	}
	index := len(s.data) - 1
	popData = s.data[index]
	s.data = s.data[:index]
	return
}

func (s *ArrStack) Push(data int) {
	s.data = append(s.data, data)
}

func (s *ArrStack) PrintStack() {
	fmt.Println(s.data)
}
