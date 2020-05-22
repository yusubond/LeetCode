package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	if !s.IsEmpty() || s.len != 0 || s.Len() != 0 {
		t.Fatalf("failed to new stack")
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.IsEmpty() || s.Len() != 3 {
		t.Fatalf("failed to push element")
	}
	s.Pop()
	if s.IsEmpty() || s.Len() != 2 {
		t.Fatalf("faield to pop element")
	}
	if e := s.Peek(); e != 2 {
		t.Fatalf("faield to peek element")
	}

	s.Pop()
	s.Pop()
	if !s.IsEmpty() {
		t.Fatalf("failed to check IsEmpty")
	}

	if s.Peek() != nil {
		t.Fatalf("failed to peek element")
	}
}