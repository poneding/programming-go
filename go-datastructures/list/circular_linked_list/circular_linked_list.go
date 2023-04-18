package main

type Ring struct {
	Head *Node
	Size int
}
type Node struct {
	Value      interface{}
	Prev, Next *Node
}

func NewRing() *Ring {
	return new(Ring)
}

func (r *Ring) AddNode(index int, v interface{}) {
	n := &Node{
		Value: v,
	}
	if r.Head == nil {
		r.Head = n
		r.Head.Prev = n
		r.Head.Next = n
	} else {
		n := &Node{
			Value: v,
		}
		next := r.Head
		for i := 0; i < index-1; i++ {
			next = next.Next
		}
		next.Prev.Next = n
		n.Prev = next.Prev
		next.Prev = n
		n.Next = next

		r.Head.Prev.Next = n
		n.Prev = r.Head.Prev
		n.Next = r.Head
		r.Head.Prev = n
	}
	r.Size++
}

func (r *Ring) RemoveNode(index int) {

}
