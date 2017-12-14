package main

import (
	"fmt"
)

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
func main() {
	var primes []int64

	for x := int64(2); len(primes) <= 10001; x++ {
		isPrime := true
		for _, v := range primes {
			if x%v == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, x)
		}
	}

	fmt.Print(primes[10000])
}
