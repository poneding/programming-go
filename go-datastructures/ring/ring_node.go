package main

import "fmt"

type RingNode struct {
	prev, next *RingNode
	Value      interface{}
}

func New() *RingNode {
	n := new(RingNode)
	n.prev = n
	n.next = n
	return n
}

// 创建N个节点的循环链表
func NewFrom(n int) *RingNode {
	if n <= 0 {
		return nil
	}
	r := new(RingNode)
	p := r
	for i := 1; i < n; i++ {
		p.next = &RingNode{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// 获取下一个节点
func (r *RingNode) Next() *RingNode {
	if r.next == nil {
		return New()
	}
	return r.next
}

// 获取上一个节点
func (r *RingNode) Prev() *RingNode {
	if r.prev == nil {
		return New()
	}
	return r.prev
}

// 获取第 n 个节点
func (r *RingNode) Move(n int) *RingNode {
	if r.next == nil {
		return New()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *RingNode) Link(s *RingNode) *RingNode {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 查看循环链表长度
func (r *RingNode) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

func linkNewTest() {
	// 第一个节点
	r := &RingNode{Value: -1}
	// 链接新的五个节点
	r.Link(&RingNode{Value: 1})
	r.Link(&RingNode{Value: 2})
	r.Link(&RingNode{Value: 3})
	r.Link(&RingNode{Value: 4})
	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.Next()
		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}
func main() {
	linkNewTest()
}
