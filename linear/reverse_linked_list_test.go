package linear

import (
	"testing"
)

func TestReverseLinkedListWithStack(t *testing.T) {
	var list1 LinkedList
	list1.Init(1, 2, 3, 4, 5, 6)
	err := list1.ReverseLinkedListWithStack(4)
	if index, _ := list1.FindKth(3); err != nil || index != 1 {
		t.Error("error: Linked List [method: ReverseLinkedListWithStack] did not pass test")
	}
	var list2 LinkedList
	list2.Init(1, 2, 3, 4, 5, 6)
	err = list2.ReverseLinkedListWithStack(7)
	if err == nil {
		t.Error("error: Linked List [method: ReverseLinkedListWithStack] did not pass test")
	}
}

func TestReverseLinkedList(t *testing.T) {
	var list1 LinkedList
	list1.Init(1, 2, 3, 4, 5, 6)
	err := list1.ReverseLinkedList(4)
	if index, _ := list1.FindKth(3); err != nil || index != 1 {
		t.Error("error: Linked List [method: ReverseLinkedList] did not pass test")
	}
	var list2 LinkedList
	list2.Init(1, 2, 3, 4, 5, 6)
	err = list2.ReverseLinkedList(7)
	if err == nil {
		t.Error("error: Linked List [method: ReverseLinkedList] did not pass test")
	}
}
