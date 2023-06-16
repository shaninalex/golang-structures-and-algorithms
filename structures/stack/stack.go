package stack

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int64
}

type Stack struct {
	Top        int
	Capacity   int
	stackArray []Node
}

func (stack *Stack) init(capacity int) *Stack {
	stack.Top = -1
	stack.Capacity = capacity

	stack.stackArray = make([]Node, capacity)
	return stack
}

func NewStack(capacity int) *Stack {
	return new(Stack).init(capacity)
}

func (stack *Stack) Push(data Node) error {
	if stack.IsFull() {
		return errors.New("stack overflow")
	}

	stack.Top++
	stack.stackArray[stack.Top] = data
	return nil
}

func (stack *Stack) Pop() (*Node, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.Top]
	stack.Top--
	return &temp, nil
}

func (stack *Stack) Peek() (*Node, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.Top]
	return &temp, nil
}

func (stack *Stack) IsFull() bool {
	return stack.Top == int(stack.Capacity)-1
}

func (stack *Stack) IsEmpty() bool {
	return stack.Top == -1
}

func (stack *Stack) Size() uint {
	return uint(stack.Top + 1)
}

func (stack *Stack) Clear() {
	stack.stackArray = nil
	stack.Top = -1
}

func (stack *Stack) Print() {
	fmt.Print("[")
	for i := 0; i <= stack.Top; i++ {
		fmt.Print(stack.stackArray[i].Value)
		if i != stack.Top {
			fmt.Print(",")
		}
	}
	fmt.Print("]\n")
}
