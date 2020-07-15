package containers

import "fmt"

type stackItem struct {
	data interface{}
	prev *stackItem
}

type Stack struct {
	top  *stackItem
	size uint
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(el interface{}) {
	newEl := &stackItem{
		data: el,
		prev: s.top,
	}
	s.top = newEl
	s.size++
}

func (s *Stack) GetSize() uint {
	return (*s).size
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.size > 0 {
		data := s.top.data
		s.top = s.top.prev
		s.size--
		return data, true
	}

	return nil, false
}

func (s *Stack) Peek() interface{} {
	if s.size > 0 {
		return s.top.data
	}

	return nil
}

func (s *Stack) PrintPop() {
	size := s.size
	for i := 0; uint(i) < size; i++ {
		fmt.Println(s.Pop())
	}
}
