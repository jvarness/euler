package main

import (
	"fmt"
)

/*
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func main() {
	var sum = 0
	var square = 0
	for x := 1; x <= 100; x++ {
		sum += x
		square += (x * x)
	}
	fmt.Print((sum * sum) - square)
}
