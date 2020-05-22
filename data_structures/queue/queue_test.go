package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue(4)
	if !q.IsEmpty() || q.IsFull() || q.Size() != 0 {
		t.Fatalf("failed to new queue")
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)
	if q.IsEmpty() || q.IsFull() || q.Size() != 3 {
		t.Fatalf("failed to push element")
	}

	q.Push(4)
	if !q.IsFull() {
		t.Fatalf("failed to check queue IsFull")
	}

	if q.Front() != 1 {
		t.Fatalf("failed to get front element")
	}

	if q.Back() != 4 {
		t.Fatalf("failed to get back element")
	}

	q.Pop()
	if q.Front() != 2 || q.Back() != 4{
		t.Fatalf("failed to pop element")
	}

	q.Pop()
	q.Pop()
	if q.IsEmpty() || q.IsFull() {
		t.Fatalf("failed to check IsEmpty or IsFull")
	}

	q.Pop()
	if !q.IsEmpty() {
		t.Fatalf("failed to check IsEmpty")
	}
}