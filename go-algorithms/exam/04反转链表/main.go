package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	res := reverseList(head)
	fmt.Print(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	if cur != nil {
		tmp := ListNode{
			Val:  cur.Next.Val,
			Next: cur.Next.Next,
		}
		cur.Next = pre
		pre = cur
		cur = &tmp
	}
	return pre
}

func setNode(head, res *ListNode) {
	if head == nil {
		res = &ListNode{
			Val: head.Val,
		}
	}

	setNode(head.Next, res.Next)
}
