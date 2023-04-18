package main

type ListNode struct {
	prev, next *ListNode
	v          int
}

// 获取节点值
func (node *ListNode) Value() int {
	return node.v
}

// 获取节点前驱节点
func (node *ListNode) Prev() *ListNode {
	return node.prev
}

// 获取节点后驱节点
func (node *ListNode) Next() *ListNode {
	return node.next
}

// 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.prev != nil
}

// 是否存在前驱节点
func (node *ListNode) HasPrev() bool {
	return node.next != nil
}

// 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}
