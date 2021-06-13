package queue

import "testing"

func TestLinkedListQueue(t *testing.T) {
	var queue LinkQueue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	value, err := queue.Dequeue()
	if err != nil || value.(int) != 10 {
		t.Error("error: Linked List Queue [method: Dequeue] did not pass test")
	}
	value, err = queue.QueueHead()
	if err != nil || value.(int) != 20 {
		t.Error("error: Linked List Queue [method: QueueHead] did not pass test")
	}
	queue.Dequeue()
	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Error("error: Linked List Queue [method: IsEmpty] did not pass test")
	}
}
