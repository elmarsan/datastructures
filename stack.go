package main

type Stack struct {
	stack []uint16
	last  *uint16
}

func NewStack() *Stack {
	return &Stack{
		stack: []uint16{},
	}
}

// pop returns last element inserted on stack.
func (sp *Stack) pop() *uint16 {
	if sp.last == nil {
		return nil
	}

	last := sp.last

	newStack := make([]uint16, len(sp.stack)-1)
	copy(newStack, sp.stack[:1])
	sp.stack = newStack

	if len(sp.stack) > 0 {
		sp.last = &sp.stack[len(sp.stack)-1]
	} else {
		sp.last = nil
	}

	return last
}

// push put val at top of stack.
func (sp *Stack) push(val uint16) {
	sp.stack = append(sp.stack, val)
	sp.last = &sp.stack[len(sp.stack)-1]
}
