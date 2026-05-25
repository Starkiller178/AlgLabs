package structures

import "errors"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Stack struct {
	items []*Node
}

func NewStack() *Stack {
	return &Stack{
		items: make([]*Node, 0, 10),
	}
}

func (s *Stack) GetItems() *[]*Node {
	return &s.items
}

func (s *Stack) Push(x *Node) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() (*Node, error) {
	if len(s.items) == 0 {
		return nil, errors.New("Err stack is empty")
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

func (s *Stack) Peek() (*Node, error) {
	if len(s.items) == 0 {
		return nil, errors.New("Err stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}
