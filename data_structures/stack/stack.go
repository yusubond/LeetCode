package stack

import (
	"sync"
)

//---------------------------------------------------------------
// Stack
// LIFO后入先出的数据结构，属于线性表的一种
//---------------------------------------------------------------

type Stack struct {
	stack []interface{}
	len   int
	lock  sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0
	s.lock = sync.Mutex{}
	return s
}

func (s *Stack) Push(e interface{}) {
	s.stack = append(s.stack, e)
	s.len++
	return
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() { return nil }
	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.len--
	return e
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len == 0
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() { return nil }
	return s.stack[len(s.stack)-1]
}