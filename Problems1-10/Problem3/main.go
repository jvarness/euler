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
	var primes []int64
	var largestPrime int64

	for x := int64(2); x <= num; x++ {
		isPrime := true
		for _, v := range primes {
			if x%v == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, x)
			if num%x == 0 {
				largestPrime = x
			}
		}
		if x%int64(100000) == 0 {
			fmt.Printf("Thru: %v", x)
		}
	}

	return largestPrime
}
