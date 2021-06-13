package queue

import (
	"errors"
	"fmt"
)

type LinkQueueNode struct {
	Data interface{}
	Next *LinkQueueNode
}

type LinkQueue struct {
	Head *LinkQueueNode
	Tail *LinkQueueNode
}

func (l *LinkQueue) Enqueue(data interface{}) {
	node := LinkQueueNode{Data: data}
	if l.IsEmpty() {
		l.Head = &node
		l.Tail = &node
		return
	}
	l.Tail.Next = &node
	l.Tail = &node
}

func (l *LinkQueue) Dequeue() (data interface{}, err error) {
	data, err = l.QueueHead()
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	return
}

func (l *LinkQueue) IsEmpty() (result bool) {
	result = false
	if l.Head == nil && l.Tail == nil {
		result = true
	}
	return
}

func (l *LinkQueue) QueueHead() (data interface{}, err error) {
	if l.IsEmpty() {
		err = errors.New("error: queue is empty")
	}
	data = l.Head.Data
	return
}

func (l *LinkQueue) Print() {
	current := l.Head
	fmt.Print("[")
	for current != nil {
		fmt.Print(current.Data)
		if current.Next != nil {
			fmt.Print(" ")
		}
		current = current.Next
	}
	fmt.Println("]")
}
