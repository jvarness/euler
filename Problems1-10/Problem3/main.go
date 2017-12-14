package main

import "fmt"

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
func main() {
	x := int64(600851475143)
	fmt.Printf("Largest prime factor for %v: %v", x, largestPrimeFactor(x))
}

func largestPrimeFactor(num int64) int64 {
	var largestPrime int64

	for x := num; x >= int64(2); x-- {
		if num % x == 0 {
			isLargestPrime := true
			for v := int64(2); v < x; v++ {
				if x % v == 0 {
					isLargestPrime = false
					break
				}
			}
			if isLargestPrime {
				largestPrime = x
				break
			}
		}
	}

	return largestPrime
}
