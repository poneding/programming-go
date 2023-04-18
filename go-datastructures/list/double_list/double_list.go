package main

// import (
// 	"fmt"
// 	"sync"
// )

// // 双端列表，双端队列
// type DoubleList struct {
// 	head *ListNode  // 指向链表头部
// 	tail *ListNode  // 指向链表尾部
// 	len  int        // 列表长度
// 	lock sync.Mutex // 为了进行并发安全pop操作
// }

// // 列表节点
// type ListNode struct {
// 	pre   *ListNode // 前驱节点
// 	next  *ListNode // 后驱节点
// 	value string    // 值
// }

// // 获取节点值
// func (node *ListNode) GetValue() string {
// 	return node.value
// }

// // 获取节点前驱节点
// func (node *ListNode) GetPre() *ListNode {
// 	return node.pre
// }

// // 获取节点后驱节点
// func (node *ListNode) GetNext() *ListNode {
// 	return node.next
// }

// // 是否存在后驱节点
// func (node *ListNode) HashNext() bool {
// 	return node.pre != nil
// }

// // 是否存在前驱节点
// func (node *ListNode) HashPre() bool {
// 	return node.next != nil
// }

// // 是否为空节点
// func (node *ListNode) IsNil() bool {
// 	return node == nil
// }

// // 返回列表长度
// func (list *DoubleList) Len() int {
// 	return list.len
// }

// // 添加节点到链表头部的第N个元素之前，N=0表示新节点成为新的头部
// func (list *DoubleList) AddNodeFromHead(n int, v string) {
// 	// 加并发锁
// 	list.lock.Lock()
// 	defer list.lock.Unlock()
// 	// 索引超过列表长度，一定找不到，panic
// 	if n > list.len {
// 		panic("index out")
// 	}

// 	targetNode := list.head
// 	newNode := &ListNode{value: v}

// 	if list.len == 0 {
// 		list.head = newNode
// 		list.tail = newNode
// 	} else {
// 		// 分三种情况：
// 		if n == 0 {
// 			// 1. 插入head
// 			newNode.next = targetNode
// 			targetNode.pre = newNode
// 			list.head = newNode
// 		} else if n == list.len {
// 			// 3. 插入tail
// 			newNode.pre = list.tail
// 			list.tail.next = newNode
// 			list.tail = newNode
// 		} else {
// 			// 2. 介于head和tail之间
// 			for i := 1; i <= n; i++ {
// 				targetNode = targetNode.next
// 			}
// 			targetNode.pre.next = newNode
// 			newNode.next = targetNode
// 			newNode.pre = targetNode.pre
// 			targetNode.pre = newNode
// 		}
// 	}

// 	list.len++
// }

// // 添加节点到链表尾部的第N个元素之后，N=0表示新节点成为新的尾部
// func (list *DoubleList) AddNodeFromTail(n int, v string) {
// 	// 加并发锁
// 	list.lock.Lock()
// 	defer list.lock.Unlock()
// 	// 索引超过列表长度，一定找不到，panic
// 	if n > list.len {
// 		panic("index out")
// 	}

// 	targetNode := list.tail
// 	newNode := &ListNode{value: v}

// 	if list.len == 0 {
// 		list.head = newNode
// 		list.tail = newNode
// 	} else {
// 		// 分三种情况：
// 		if n == 0 {
// 			// 1. 插入tail
// 			newNode.pre = targetNode
// 			targetNode.next = newNode
// 			list.tail = newNode
// 		} else if n == list.len {
// 			// 3. 插入head
// 			newNode.pre = list.tail
// 			list.tail.next = newNode
// 			list.tail = newNode
// 		} else {
// 			// 2. 介于tail和head之间
// 			for i := 1; i <= n; i++ {
// 				targetNode = targetNode.pre
// 			}
// 			targetNode.next.pre = newNode
// 			newNode.pre = targetNode
// 			newNode.next = targetNode.next
// 			targetNode.next = newNode
// 		}
// 	}

// 	list.len++
// }

// // 返回列表链表头结点
// func (list *DoubleList) First() *ListNode {
// 	return list.head
// }

// // 返回列表链表尾结点
// func (list *DoubleList) Last() *ListNode {
// 	return list.tail
// }

// // 从头部开始往后找，获取第N+1个位置的节点，索引从0开始。
// func (list *DoubleList) IndexFromHead(n int) *ListNode {
// 	// 索引超过或等于列表长度，一定找不到，返回空指针
// 	if n >= list.len {
// 		return nil
// 	}
// 	// 获取头部节点
// 	node := list.head
// 	// 往后遍历拿到第 N+1 个位置的元素
// 	for i := 1; i <= n; i++ {
// 		node = node.next
// 	}
// 	return node
// }

// // 从尾部开始往前找，获取第N+1个位置的节点，索引从0开始。
// func (list *DoubleList) IndexFromTail(n int) *ListNode {
// 	// 索引超过或等于列表长度，一定找不到，返回空指针
// 	if n >= list.len {
// 		return nil
// 	}
// 	// 获取尾部节点
// 	node := list.tail
// 	// 往前遍历拿到第 N+1 个位置的元素
// 	for i := 1; i <= n; i++ {
// 		node = node.pre
// 	}
// 	return node
// }

// // 从头部开始往后找，获取第N+1个位置的节点，并移除返回
// func (list *DoubleList) PopFromHead(n int) *ListNode {
// 	// 加并发锁
// 	list.lock.Lock()
// 	defer list.lock.Unlock()
// 	// 索引超过或等于列表长度，一定找不到，返回空指针
// 	if n >= list.len {
// 		return nil
// 	}
// 	// 获取头部
// 	node := list.head
// 	// 往后遍历拿到第 N+1 个位置的元素
// 	for i := 1; i <= n; i++ {
// 		node = node.next
// 	}
// 	// 移除的节点的前驱和后驱
// 	pre := node.pre
// 	next := node.next
// 	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
// 	if pre.IsNil() && next.IsNil() {
// 		list.head = nil
// 		list.tail = nil
// 	} else if pre.IsNil() {
// 		// 表示移除的是头部节点，那么下一个节点成为头节点
// 		list.head = next
// 		next.pre = nil
// 	} else if next.IsNil() {
// 		// 表示移除的是尾部节点，那么上一个节点成为尾节点
// 		list.tail = pre
// 		pre.next = nil
// 	} else {
// 		// 移除的是中间节点
// 		pre.next = next
// 		next.pre = pre
// 	}
// 	// 节点减一
// 	list.len = list.len - 1
// 	return node
// }

// // 从尾部开始往前找，获取第N+1个位置的节点，并移除返回
// func (list *DoubleList) PopFromTail(n int) *ListNode {
// 	// 加并发锁
// 	list.lock.Lock()
// 	defer list.lock.Unlock()
// 	// 索引超过或等于列表长度，一定找不到，返回空指针
// 	if n >= list.len {
// 		return nil
// 	}
// 	// 获取尾部
// 	node := list.tail
// 	// 往前遍历拿到第 N+1 个位置的元素
// 	for i := 1; i <= n; i++ {
// 		node = node.pre
// 	}
// 	// 移除的节点的前驱和后驱
// 	pre := node.pre
// 	next := node.next
// 	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
// 	if pre.IsNil() && next.IsNil() {
// 		list.head = nil
// 		list.tail = nil
// 	} else if pre.IsNil() {
// 		// 表示移除的是头部节点，那么下一个节点成为头节点
// 		list.head = next
// 		next.pre = nil
// 	} else if next.IsNil() {
// 		// 表示移除的是尾部节点，那么上一个节点成为尾节点
// 		list.tail = pre
// 		pre.next = nil
// 	} else {
// 		// 移除的是中间节点
// 		pre.next = next
// 		next.pre = pre
// 	}
// 	// 节点减一
// 	list.len = list.len - 1
// 	return node
// }
// func main() {
// 	list := new(DoubleList)
// 	// 在列表头部插入新元素
// 	list.AddNodeFromHead(0, "33")
// 	list.AddNodeFromHead(0, "11")
// 	list.AddNodeFromHead(1, "22")
// 	list.AddNodeFromHead(3, "44")
// 	// 在列表尾部插入新元素
// 	// 在列表尾部插入新元素
// 	list.AddNodeFromTail(0, "66")
// 	list.AddNodeFromTail(1, "55")
// 	// 正常遍历，比较慢
// 	for i := 0; i < list.Len(); i++ {
// 		// 从头部开始索引
// 		node := list.IndexFromHead(i)
// 		// 节点为空不可能，因为list.Len()使得索引不会越界
// 		if !node.IsNil() {
// 			fmt.Println(node.GetValue())
// 		}
// 	}
// 	fmt.Println("----------")
// 	// 正常遍历，特别快
// 	// 先取出第一个元素
// 	first := list.First()
// 	for !first.IsNil() {
// 		// 如果非空就一直遍历
// 		fmt.Println(first.GetValue())
// 		// 接着下一个节点
// 		first = first.GetNext()
// 	}
// 	fmt.Println("----------")
// 	// 元素一个个 POP 出来
// 	for {
// 		node := list.PopFromHead(0)
// 		if node.IsNil() {
// 			// 没有元素了，直接返回
// 			break
// 		}
// 		fmt.Println(node.GetValue())
// 	}
// 	fmt.Println("----------")
// 	fmt.Println("len", list.Len())
// }
