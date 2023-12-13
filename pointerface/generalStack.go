package pointerface

import "errors"

// Stack in package pointerface is a general stack in which you can push different types of elements.
type Stack []interface{}

// Length tells the number of elements in the stack.
func (stack *Stack) Length() int {
	return len(*stack)
}

// IsEmpty provides a neat way to check if the stack is empty.
func (stack *Stack) IsEmpty() bool {
	return stack.Length() == 0
}

// Push let you push a set of elements into the stack in one go.
func (stack *Stack) Push(elements ...interface{}) {
	for this := range elements {
		*stack = append(*stack, elements[this])
	}
}

// Pop pops out the top element in the stack.
func (stack *Stack) Pop() (interface{}, error) {
	stk := *stack
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	} else {
		topEle := stk[stk.Length()-1]
		*stack = stk[:stk.Length()-1]
		return topEle, nil
	}
}
