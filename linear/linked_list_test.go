package linear

import "testing"

func TestLinkedList(t *testing.T) {
	var list LinkedList
	list.Init(1, 2, 3, 4, 5, 6)
	// size
	if list.Size() != 6 {
		t.Error("error: Linked List [method: Size] did not pass test")
	}
	// FindKth
	if value, err := list.FindKth(3); value.(int) != 4 || err != nil {
		t.Error("error: Linked List [method: FindKth] did not pass test")
	}
	// FindValue
	if index, err := list.FindValue(3); index != 2 || err != nil {
		t.Error("error: Linked List [method: FindValue] did not pass test")
	}
	// Insert
	err := list.Insert(3, 99)
	if err != nil {
		t.Error("error: Linked List [method: Insert] did not pass test")
	}
	if value, _ := list.FindKth(3); value != 99 {
		t.Error("error: Linked List [method: Insert] did not pass test")
	}
	// Delete
	err = list.Delete(3)
	if err != nil {
		t.Error("error: Linked List [method: Delete] did not pass test")
	}
	if value, _ := list.FindKth(3); value != 4 {
		t.Error("error: Linked List [method: Insert] did not pass test")
	}
}
