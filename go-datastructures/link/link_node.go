package main

import "fmt"

type LinkNode struct {
	Value interface{}
	Next  *LinkNode
}

func main() {
	// 新的节点
	node := new(LinkNode)
	node.Value = 2
	// 新的节点
	node1 := new(LinkNode)
	node1.Value = 3
	node.Next = node1 // node1 链接到 node 节点上
	// 新的节点
	node2 := new(LinkNode)
	node2.Value = 4
	node1.Next = node2 // node2 链接到 node1 节点上
	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			fmt.Println(nowNode.Value)
			// 获取下一个节点
			nowNode = nowNode.Next
		}
		// 如果下一个节点为空，表示链表结束了
		break
	}
}
