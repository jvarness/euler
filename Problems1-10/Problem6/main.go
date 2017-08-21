package main

import (
	"fmt"
)

func main() {
	var sum = 0
	var square = 0
	for x := 1; x <= 100; x++ {
		sum += x
		square += (x * x)
	}
	fmt.Print((sum * sum) - square)
}
