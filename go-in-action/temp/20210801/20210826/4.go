package main

type stack1 struct {
	values []int
	size   int
}
type twoStackQueue struct {
	stackPush *stack1
	stackPop  *stack1
}

// 使用两个栈实现一个队列，push，pop，peek
func newTwoStackQueue() *twoStackQueue {
	q := new(twoStackQueue)
	q.stackPop = newStack1()
	q.stackPush = newStack1()
	return q
}

func (q *twoStackQueue) pushToStackPop() {
	if q.stackPop.size == 0 {
		for q.stackPush.size > 0 {
			q.stackPop.push(q.stackPush.pop())
		}
	}
}

func (q *twoStackQueue) push(v int) {
	q.stackPush.push(v)
	q.pushToStackPop()
}

func (q *twoStackQueue) pop() int {
	if q.stackPop.size == 0 && q.stackPush.size == 0 {
		panic("no data")
	}
	q.pushToStackPop()
	return q.stackPop.pop()
}

func (q *twoStackQueue) peek() int {
	if q.stackPop.size == 0 && q.stackPush.size == 0 {
		panic("no data")
	}
	q.pushToStackPop()
	return q.stackPop.peek()
}

func newStack1() *stack1 {
	return new(stack1)
}

func (s *stack1) push(v int) {
	s.values = append(s.values, v)
	s.size++
}

func (s *stack1) pop() int {
	if s.size == 0 {
		panic("no data")
	}

	v := s.values[s.size]
	s.values[s.size] = 0
	s.size--
	return v
}

func (s *stack1) peek() int {
	if s.size == 0 {
		panic("no data")
	}
	return s.values[s.size]
}
