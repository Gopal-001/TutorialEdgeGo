package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is empty")
	}
	pp := q.Items[0]
	q.Items = q.Items[1:]
	return pp, nil
}

func (q *Queue) Push(f Flight) {
	q.Items = append(q.Items, f)
}

func (q *Queue) Peek() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is empty")
	}
	return q.Items[0], nil
}

func (q *Queue) IsEmpty() bool {
	if len(q.Items) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Go Queue Implementation")
	q := Queue{
		Items: []Flight{},
	}
	fmt.Println(q.IsEmpty())
	q.Push(Flight{Price: 5})
	q.Push(Flight{Price: 1000})
	q.Push(Flight{Price: 23})
	q.Push(Flight{Price: 7})
	fmt.Println(q.IsEmpty())
	pp, _ := q.Pop()
	fmt.Println("Pop: ", pp)
	vr, _ := q.Peek()
	fmt.Println("Peek: ", vr)
}
