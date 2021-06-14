package linear

import (
	"algorithm/stack"
	"errors"
)

// 逆转链表的指定个数的节点: 借助stack结构实现
func (l *LinkedList) ReverseLinkedListWithStack(n int) (err error) {
	if n > l.Size() {
		err = errors.New("error: out of range")
		return
	}
	// 将要逆序的节点的指针依原本的顺序放入栈中
	var stack stack.LinkStack
	current := l.Head
	for i := 0; i < n; i++ {
		stack.Push(current)
		current = current.Next
	}
	// 记录剩余的节点，并清空链表头
	rest := current
	l.Head = nil
	// 对储存了原先链表顺序的栈进行出栈，此时的顺序为逆序
	for !stack.IsEmpty() {
		// 若头节点为空，则指向第一个出栈的节点
		if l.Head == nil {
			temp, _ := stack.Pop()
			l.Head = temp.(*LinkedListNode)
			current = l.Head
			continue
		}
		temp, _ := stack.Pop()
		current.Next = temp.(*LinkedListNode)
		current = current.Next
	}
	// 链接上剩下无需逆序的节点
	current.Next = rest
	return
}

// 逆转链表的指定个数的节点，不借助任何事先实现的其它数据结构
func (l *LinkedList) ReverseLinkedList(n int) (err error) {
	if n > l.Size() {
		err = errors.New("error: out of range")
		return
	}
	// 使用两个个指针分别记住新的头部指针与剩余指针
	var newHead, rest *LinkedListNode
	current := l.Head
	for i := 0; i < n; i++ {
		// 初始化两个指针，并将原先头节点的的next指向空，此步不做逆序操作
		if i == 0 {
			newHead = current.Next
			rest = current.Next.Next
			current.Next = nil
			continue
		}
		// 开始逆序
		newHead.Next = current
		current = newHead
		newHead = rest
		// 防止空指针报错
		if rest != nil {
			rest = rest.Next
		}
	}
	// 连接剩下无须逆序的部分
	l.Head.Next = newHead
	l.Head = current
	return
}
