package queue

import "testing"

func TestSliceQueue(t *testing.T) {
	var queue SliceQueue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	value, err := queue.Dnqueue()
	if err != nil || value != 10 {
		t.Error("error: Slice Queue [method: Dequeue] did not pass test")
	}
	value, err = queue.Head()
	if err != nil || value != 20 {
		t.Error("error: Slice Queue [method: Head] did not pass test")
	}
	queue.Dnqueue()
	queue.Dnqueue()
	isEmpty := queue.IsEmpty()
	if isEmpty != true {
		t.Error("error: Slice Queue [method: IsEmpty] did not pass test")
	}
}
