package my_font

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func Test_font(t *testing.T) {
	fmt.Println(New("aaa", 10))
	fmt.Println(New("bbb", 10))
	fmt.Println(New("ccc", 10))
	fmt.Println(New("ddd", 10))
	fmt.Println(New("aaa", 3))
	fmt.Println(New("aaa", 1000))
}

var sum = 0

func print(job <-chan int, done chan<- struct{}, i int) {
	for {
		value := <-job
		if value == -1 {
			break
		}
		//fmt.Println(i, value)
		time.Sleep(time.Duration(rand.Int63() % 10000000))

		sum += value
	}

	done <- struct{}{}
}
func TestGoRouten(t *testing.T) {
	ch := make(chan int, 1003)
	done := make(chan struct{}, 3)
	fmt.Println(cap(ch))
	for i := 0; i < 1000; i++ {
		ch <- i+1
	}
	for i := 0; i < 3; i++ {
		ch <- -1
	}
	close(ch)
	fmt.Println("=============step 1===================")
	for i := 0; i < 3; i++ {
		go print(ch, done, i)

	}
	fmt.Println("=============step 2===================")
	for i := 0; i < 3; i++ {
		<-done
	}
	fmt.Println(sum)

}
