package main

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main4() {
	l1 := &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Val: 9,
			},
			Val: 9,
		},
		Val: 9,
	}

	l2 := &ListNode{
		Next: &ListNode{
			Val: 9,
		},
		Val: 9,
	}
	res := addTwoNumbers(l1, l2)
	fmt.Printf("%+v\n", res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	tail := new(ListNode)
	var carry int
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		v := sum % 10
		carry = sum / 10

		if res == nil {
			res = &ListNode{
				Val: v,
			}
			tail = res
		} else {
			tail.Next = &ListNode{Val: v}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return res
}

func addTwoNumbers2(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
