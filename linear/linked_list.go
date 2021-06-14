package linear

import (
	"errors"
	"fmt"
)

type LinkedListNode struct {
	Data interface{}
	Next *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

// 初始化链表
func (l *LinkedList) Init(data ...interface{}) {
	var current *LinkedListNode
	for _, v := range data {
		node := LinkedListNode{Data: v}
		if l.Head == nil {
			l.Head = &node
			current = &node
		} else {
			current.Next = &node
			current = current.Next
		}
	}
}

// 判断链表是否为空
func (l *LinkedList) IsEmpty() (result bool) {
	result = false
	if l.Head == nil {
		result = true
	}
	return
}

// 打印链表
func (l *LinkedList) Print() {
	fmt.Print("[")
	current := l.Head
	for current != nil {
		fmt.Print(current.Data)
		if current.Next != nil {
			fmt.Print(" ")
		}
		current = current.Next
	}
	fmt.Println("]")
}

// 返回链表长度
func (l *LinkedList) Size() (size int) {
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}
	return
}

// 按下标查找链表元素
func (l *LinkedList) FindKth(index int) (data interface{}, err error) {
	if index > l.Size()-1 {
		err = errors.New("error: index out of range")
		return
	}
	current := l.Head
	i := 0
	for current != nil {
		if index == i {
			data = current.Data
			break
		}
		current = current.Next
		i++
	}
	return
}

// 按值查找链表元素
func (l *LinkedList) FindValue(value interface{}) (index int, err error) {
	current := l.Head
	found := false
	for current != nil {
		if current.Data == value {
			found = true
			break
		}
		index++
		current = current.Next
	}
	if !found {
		err = errors.New("error: value not found")
	}
	return
}

// 在特定位置插入元素
func (l *LinkedList) Insert(index int, value interface{}) (err error) {
	if index > l.Size() {
		err = errors.New("error: index out of range")
		return
	}
	node := LinkedListNode{Data: value}
	i := 0
	current := l.Head
	// 当位置是0的时候需要特别处理
	if index == 0 {
		node.Next = l.Head
		l.Head = &node
		return
	}
	// 位置是非0时
	for current != nil {
		if i == index-1 {
			node.Next = current.Next
			current.Next = &node
			break
		}
		i++
		current = current.Next
	}
	return
}

// 删除特定位置的元素
func (l *LinkedList) Delete(index int) (err error) {
	if index > l.Size()-1 {
		err = errors.New("error: index out of range")
		return
	}
	current := l.Head
	i := 0
	// 位置为0时需要特殊处理
	if index == 0 {
		l.Head = l.Head.Next
		return
	}
	for current != nil {
		if i == index-1 {
			if current.Next == nil {
				current.Next = nil
			} else {
				current.Next = current.Next.Next
			}
			break
		}
		current = current.Next
		i++
	}
	return
}
