package main

func main() {
	// 实现一个栈（先进后出），push,pop,getMin
}

type stack struct {
	values []int
	size   int
}

// 实现一个栈（先进后出），push,pop,getMin
type myStack struct {
	stackData *stack
	stackMin  *stack
}

func NewMyStack() *myStack {
	s := new(myStack)
	s.stackData = newStack()
	s.stackMin = newStack()
	return s
}

func (ms *myStack) Push(v int) {
	ms.stackData.push(v)
	if ms.stackMin.size == 0 {
		ms.stackMin.push(v)
	} else {
		min := ms.stackMin.peek()
		if v <= min {
			ms.stackMin.push(v)
		} else {
			ms.stackMin.push(min)
		}
	}
	return
}

func (ms *myStack) Pop() int {
	if ms.stackData.size == 0 {
		panic("no data")
	}

	ms.stackMin.pop()
	return ms.stackData.pop()
}

func (ms *myStack) GetMin() int {
	if ms.stackData.size == 0 {
		panic("no data")
	}
	return ms.stackMin.peek()
}

func newStack() *stack {
	return new(stack)
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
	s.size++
}

func (s *stack) pop() int {
	if s.size == 0 {
		panic("no data")
	}

	v := s.values[s.size]
	s.values[s.size] = 0
	s.size--
	return v
}

func (s *stack) peek() int {
	if s.size == 0 {
		panic("no data")
	}
	return s.values[s.size]
}
