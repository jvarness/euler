package main

import (
	"fmt"
)

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
