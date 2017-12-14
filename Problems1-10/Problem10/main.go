package main

import (
	"fmt"
)
/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
func main() {
	var sum = int64(0)

	for x := int64(2); x <= int64(2e6); x++ {
		isPrime := true
		for v := int64(2); v < x; v++ {
			if x%v == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			sum = sum + x
		}
	}

	fmt.Print(sum)
}