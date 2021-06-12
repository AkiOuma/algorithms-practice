package linear

import (
	"errors"
	"fmt"
)

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

// 创建链表
func BuildLinkedList(data []int) LinkedListNode {
	var head LinkedListNode
	current := &head
	for _, v := range data {
		node := LinkedListNode{Data: v}
		current.Next = &node
		current = current.Next
	}
	return head
}

// 打印链表
func PrintLinkedList(head LinkedListNode) {
	current := head.Next
	fmt.Print("[")
	for current != nil {
		fmt.Print(current.Data)
		current = current.Next
		if current != nil {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

// 获取链表长度
func (l *LinkedListNode) Length() (size int) {
	current := l.Next
	for current != nil {
		size++
		current = current.Next
	}
	return
}

// 按下标寻找链表中的数据的值
func (l *LinkedListNode) FindKth(index int) (data int, err error) {
	current := l.Next
	i := 1
	for current != nil {
		if i == index {
			data = current.Data
			break
		}
		current = current.Next
		if current != nil {
			i++
		}
	}
	if i < index {
		err = errors.New("error: out of range")
	}
	return
}

// 按值寻找链表中第一个符合的节点的下标
func (l *LinkedListNode) FindValue(data int) (index int, err error) {
	current := l.Next
	i := 1
	for current != nil {
		if current.Data == data {
			index = i
			break
		}
		current = current.Next
		i++
		if current == nil {
			err = errors.New("error: data not exist")
		}
	}
	return
}

// 在特定位置插入新的节点
func (l *LinkedListNode) InsertNode(data int, index int) (err error) {
	current := l
	i := 1
	for current != nil {
		if i == index {
			node := LinkedListNode{Data: data}
			node.Next = current.Next
			current.Next = &node
			break
		}
		current = current.Next
		i++
		if current == nil {
			err = errors.New("error: new index out of range")
		}
	}
	return
}

// 删除一个特定位置的节点
func (l *LinkedListNode) DeleteNode(index int) (err error) {
	current := l
	i := 1
	for current.Next != nil {
		if index == i {
			if current.Next != nil {
				current.Next = current.Next.Next
			} else {
				current.Next = nil
			}
			break
		}
		current = current.Next
		i++
		if current.Next == nil {
			err = errors.New("error: index out of range")
		}
	}
	return
}
