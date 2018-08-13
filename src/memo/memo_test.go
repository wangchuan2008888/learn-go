package memo

import (
	"testing"
	"fmt"
)

type memoizeFunction func(int) interface{}

var Fabonacci memoizeFunction

func Memoize(funtion memoizeFunction) memoizeFunction {
	cache := make(map[int]interface{})
	return func(i int) interface{} {
		if value, found := cache[i]; found {
			return value
		} else {
			value := funtion(i)
			cache[i] = value
			return value
		}
	}
}
func Init() {
	Fabonacci = Memoize(func(i int) interface{} {
		if i < 2 {
			return 1
		} else{
			return Fabonacci(i-1).(int) + Fabonacci(i-2).(int)
		}
	})
}
func Test_memo(t *testing.T) {
	Init()
	a := []int{1,2,3,4,5,6,7,8,9,10}
	for i :=range a{
		fmt.Println(Fabonacci(i))

	}
}
