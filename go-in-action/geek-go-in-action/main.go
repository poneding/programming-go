package main

import "fmt"

func main() {
	intset := &IntSet{
		data: make(map[int]bool),
	}
	intset.Add(4)
	fmt.Println(intset.Contains(4))
	fmt.Println(intset.Contains(5))
}

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}
func (set *IntSet) Add(x int) {
	set.data[x] = true
}
func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}
