package stack

import "fmt"

type iStack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Show()
}

type Stack struct {
	list []interface{}
	size uint
}

func NewStack() iStack {
	var i iStack
	i = &Stack{}
	return i
}

func (s *Stack) Push(v interface{}) {
	s.list = append(s.list, v)
	s.size++
}

func (s *Stack) Pop() interface{} {
	v := s.list[len(s.list)-1]
	s.list = s.list[1:]
	s.size--
	return v
}

func (s *Stack) Peek() interface{} {
	return s.list[len(s.list)-1]
}

func (s *Stack) Show() {
	fmt.Printf("\nstack size: %v \n", s.size)
	fmt.Printf("len %v cap %v \n\n", len(s.list), cap(s.list))
	for i, v := range s.list {
		fmt.Printf("%v | %v \n", i, v)
	}
	fmt.Printf("\n")
}
