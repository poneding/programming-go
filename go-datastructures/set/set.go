package main

import "sync"

type Set struct {
	// 为什么使用空结构体？
	// 因为：空结构体的内存地址都一样，并且不占用内存空间。
	data map[string]struct{}
	len  int
	lock sync.RWMutex
}

func NewSet(cap int) *Set {
	set := &Set{
		data: make(map[string]struct{}, cap),
	}
	return set
}

func (set *Set) Add(item string) {
	set.lock.Lock()
	defer set.lock.Unlock()
	set.data[item] = struct{}{}
	set.len = len(set.data)
}

func (set *Set) Remove(item string) {
	set.lock.Lock()
	defer set.lock.Unlock()
	if set.len == 0 {
		return
	}

	delete(set.data, item)
	set.len = len(set.data)
}

func (set *Set) Has(item string) bool {
	set.lock.RLock()
	defer set.lock.RUnlock()

	_, ok := set.data[item]
	return ok
}

func (set *Set) Size() int {
	return set.len
}

func (set *Set) IsEmpty() bool {
	return set.len == 0
}

func (set *Set) Clear() {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.data = make(map[string]struct{})
	set.len = 0
}

func (set *Set) List() []string {
	set.lock.Lock()
	defer set.lock.Unlock()

	l := make([]string, 0)
	for k := range set.data {
		l = append(l, k)
	}
	return l
}
