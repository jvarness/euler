package main

import "fmt"

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
func main () {
	val := int64(600851475143)
	fmt.Printf("Largest prime factor for %v: %v", val, largestPrimeFactor(val));
}

func largestPrimeFactor(num int64) int64 {
	primes := sieve(num)
	
	largestPrime := int64(0)

	for _, v := range primes {
		if num % v == 0 && num != v {
			largestPrime = v
		}
	}

	return largestPrime
}

func sieve(num int64) []int64 {
	var unsieved, sieved []int64
	unsieved = make([]int64, num)
	for i := int64(2); i <= num; i++ {
		unsieved = append(unsieved, i)
	}

	for len(unsieved) > 0 {
		v := unsieved[0]
		sieved = append(sieved, v)
		unsieved = unsieved[1:]

		for i := 0; i < len(unsieved); i++ {
			p := unsieved[i]
			if p % v == 0 {
				if i < len(unsieved) {
					unsieved = append(unsieved[:i], unsieved[i+1:]...)
				} else {
					unsieved = unsieved[:i]
				}
			}
		}
	}
	
	return sieved
}