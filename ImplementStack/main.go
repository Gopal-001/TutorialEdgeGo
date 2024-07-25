package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is empty")
	}
	ln := len(s.Items)
	poped := s.Items[ln-1]
	s.Items = s.Items[:ln-1]
	return poped, nil
}

func (s *Stack) Push(f Flight) {
	s.Items = append(s.Items, f)
}

func (s *Stack) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is empty")
	}
	ln := len(s.Items)
	poped := s.Items[ln-1]
	return poped, nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Go Stack Implementation")
	st := Stack{
		Items: []Flight{},
	}
	fmt.Println(st.IsEmpty())
	st.Push(Flight{Price: 5})
	vr, _ := st.Peek()
	fmt.Println(vr)
	st.Push(Flight{Price: 7})
	fmt.Println(st.IsEmpty())
	pp, _ := st.Pop()
	fmt.Println(pp)
}
