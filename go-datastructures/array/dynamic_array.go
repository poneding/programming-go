package main

import "sync"

type Array struct {
	data     []int
	len, cap int
	lock     sync.Mutex //为了并发安全使用的锁
}

// 新建一个可变长数组
func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}
	// 把切片当数组用
	array := make([]int, cap, cap)
	// 元数据
	s.data = array
	s.cap = cap
	s.len = 0
	return s
}

// 增加一个元素
func (a *Array) Append(element int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()
	// 大小等于容量，表示没多余位置了
	if a.len == a.cap {
		// 没容量，数组要扩容，扩容到两倍
		newCap := 2 * a.len
		// 如果之前的容量为0，那么新容量为1
		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]int, newCap, newCap)
		// 把老数组的数据移动到新数组
		for k, v := range a.data {
			newArray[k] = v
		}
		// 替换数组
		a.data = newArray
		a.cap = newCap
	}
	// 把元素放在数组里
	a.data[a.len] = element
	// 真实长度+1
	a.len = a.len + 1
}

// 增加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// 获取某个下标的元素
func (a *Array) Get(index int) int {
	// 越界了
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.data[index]
}

// 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// 返回容量
func (a *Array) Cap() int {
	return a.cap
}
