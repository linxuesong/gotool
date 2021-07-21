package queue

import "testing"

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.PrintQueue()
	q.Enqueue(2)
	q.PrintQueue()
	q.Enqueue(3)
	q.PrintQueue()
	q.Enqueue(4)
	q.PrintQueue()
	q.Dequeue()
	q.PrintQueue()
	q.Dequeue()
	q.PrintQueue()
	q.Dequeue()
	q.PrintQueue()
	q.Dequeue()
	q.PrintQueue()
	q.Dequeue()
	q.PrintQueue()
}
