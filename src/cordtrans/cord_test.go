package main

import (
	"testing"
	"strings"
)

func createSolover(message chan string) chan string {
	upper := make(chan string)
	go func() {
		for {
			in := <-message
			upper <- strings.ToUpper(in)
		}
	}()
	return upper
}
func Test_main(t *testing.T) {
	message := make(chan string, 10)
	defer close(message)
	upper := createSolover(message)
	defer close(upper)

	input := [] string{"hello", "world"}
	for idx := 0; idx < len(input); idx++ {
		message <- input[idx]
		println(<-upper)
	}
}
