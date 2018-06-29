package stack

import "errors"

type Stack [] interface{}

func (stack *Stack) Len() int {
	return len(*stack)
}

func (stack *Stack) Cap() int {
	return cap(*stack)
}

func (stack *Stack) Push(ele interface{}) {
	*stack = append(*stack, ele)
}

func (stack *Stack) Top() (interface{}, error) {
	if len(*stack) == 0 {
		return nil, errors.New("cannot Top en empty stack")
	}
	return (*stack)[len(*stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	_len := len(*stack)
	if _len == 0 {
		return nil, errors.New("cannot Pop en empty stack")
	}
	x := (*stack)[_len-1]
	*stack = (*stack)[:(_len - 1)]
	return x, nil
}
