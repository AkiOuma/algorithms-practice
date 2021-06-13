package queue

import "errors"

type SliceQueue struct {
	Data []interface{}
}

func (q *SliceQueue) Enqueue(value interface{}) {
	q.Data = append(q.Data, value)
}

func (q *SliceQueue) Dnqueue() (value interface{}, err error) {
	value, err = q.Head()
	q.Data = q.Data[1:]
	return
}

func (q *SliceQueue) IsEmpty() (result bool) {
	result = false
	if len(q.Data) == 0 {
		result = true
	}
	return
}

func (q *SliceQueue) Head() (value interface{}, err error) {
	if q.IsEmpty() {
		err = errors.New("error: queue is empty")
		return
	}
	value = q.Data[0]
	return
}
