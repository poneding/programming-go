package main

import (
	"errors"
	"fmt"
)

func main() {
	l := NewLink()
	l.AddNode(0, 1)
	l.AddNode(0, 2)
	l.AddNode(0, 4)
	l.AddNode(1, 3)
	fmt.Println(l)
	fmt.Println(l.GetNode(0))
	fmt.Println(l.GetNode(1))
	fmt.Println(l.GetNode(2))
	fmt.Println(l.GetNode(3))
	l.RemoveNode(3)
	fmt.Println(l)
	l.RemoveNode(1)
	fmt.Println(l)
}

type Link struct {
	Head, Tail *Node
	Size       int
}

type Node struct {
	Value      interface{}
	Prev, Next *Node
}

func NewLink() *Link {
	return new(Link)
}

func (l *Link) GetNode(index int) (*Node, error) {
	var res *Node
	if index >= l.Size {
		return res, errors.New("out of index")
	}
	next := l.Head
	for i := 0; i < index; i++ {
		next = next.Next
	}
	return next, nil
}

func (l *Link) AddNode(index int, v interface{}) error {
	if index > l.Size {
		return errors.New("out of index")
	}

	n := &Node{
		Value: v,
	}
	if l.Size == 0 {
		l.Head = n
		l.Tail = n
	} else {
		if index == 0 {
			l.Head.Prev = n
			n.Next = l.Head
			l.Head = n
		} else {
			next := l.Head
			for i := 0; i < index-1; i++ {
				next = next.Next
			}
			next.Next.Prev = n
			n.Next = next.Next
			n.Prev = next
			next.Next = n
		}
	}

	l.Size++
	return nil
}

func (l *Link) RemoveNode(index int) error {
	if index >= l.Size {
		return errors.New("out of index")
	}
	if index == 0 {
		l.Head = l.Head.Next
	} else if index == l.Size-1 {
		l.Tail.Prev = nil
		l.Tail = l.Tail.Prev
	} else {
		next := l.Head
		for i := 0; i < index-1; i++ {
			next = next.Next
		}
		next.Next.Next.Prev = next
		next.Next = next.Next.Next
	}
	l.Size--
	return nil
}
