package stack

import (
	"fmt"
)

// declaring the new stack type
type Stack struct {
	data []int
}

// constructor for creating new stack type
func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

// checks if there are any elements on the stack
func (stack *Stack) isEmpty() bool {
	return len(stack.data) == 0
}

// returns the element at the top of the stack, or else return an error
func (stack *Stack) Peek() (int, error) {
	if stack.isEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return stack.data[len(stack.data)-1], nil
}

// adds an element to the top of the stack
func (stack *Stack) Push(a int) {
	stack.data = append(stack.data, a)
}

// removes an element from the stack top and returns its value or throws error if stack is empty
func (stack *Stack) Pop() (int, error) {
	if len(stack.data) == 0 {
		return 0, fmt.Errorf("stack is emptpy")
	}
	element := stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return element, nil
}
