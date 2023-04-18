package main

import (
	"fmt"
	"sync"
)

type LinkNode struct {
	v    *TreeNode
	next *LinkNode
}

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

func (queue *LinkQueue) In(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		queue.root = &LinkNode{v: v}
	} else {
		n := queue.root
		for n.next != nil {
			n = n.next
		}
		n.next = &LinkNode{v: v}
	}
	queue.size++
}

func (queue *LinkQueue) Out() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty queue")
	}

	v := queue.root.v
	queue.root = queue.root.next
	queue.size--
	return v
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func (queue *LinkQueue) IsEmpty() bool {
	return queue.size == 0
}

type TreeNode struct {
	Data        string
	Left, Right *TreeNode
}

// 三种遍历：专业用词 前序/中序/后序 preorder/inorder/postorder
// 先序遍历：根》左》右
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Data)
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// 中序遍历：左》中》右
func InOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	InOrder(tree.Left)
	fmt.Println(tree.Data)
	InOrder(tree.Right)
}

// 后序遍历：左》右》中
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Println(tree.Data)
}

// 广度遍历：LayerOrder
func LayerOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	queue := new(LinkQueue)
	queue.In(tree)
	for queue.size > 0 {
		e := queue.Out()
		fmt.Println(e.Data)

		if e.Left != nil {
			queue.In(e.Left)
		}

		if e.Right != nil {
			queue.In(e.Right)
		}
	}
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}
	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	InOrder(t)
	fmt.Println("\n后序排序")
	PostOrder(t)
	fmt.Println("\n广度遍历")
	LayerOrder(t)
}
