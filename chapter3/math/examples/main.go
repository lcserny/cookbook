package main

import (
	"cookbook/chapter3/math"
	"fmt"
)

func main() {
	math.Examples()
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}
