package stack

import (
	"testing"
	"fmt"
)

func Test1(t *testing.T) {
	var a = make(Stack,0)
	fmt.Println(a.Len())
	fmt.Println(a.Cap())
	a.Push(10)
	fmt.Println(a)
	fmt.Println(a.Cap())
	i := 0
	for i < 10 {
		a.Push(i)
		i += 1
		fmt.Println(a)
	}
	i = 0
	for i < 10 {
		a.Pop()
		i += 1
		fmt.Println(a)
	}

}
