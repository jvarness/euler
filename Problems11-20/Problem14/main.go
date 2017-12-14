package main

import (
	"fmt"
)

/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?
*/
func main() {
	longestCollatzChain := 0
	collatzChain := 0
	var collatzNumber int
	var longestCollatzNumber int

	for x := 1; x < 1e6; x++ {
		collatzChain = 0
		collatzNumber = x
		for collatzNumber != 1 {
			collatzChain++

			if collatzNumber%2 == 0 {
				collatzNumber = collatzNumber / 2
			} else {
				collatzNumber = (collatzNumber * 3) + 1
			}
		}

		if collatzChain > longestCollatzChain {
			longestCollatzChain = collatzChain
			longestCollatzNumber = x
		}
	}

	fmt.Printf("Number under 1e6 with the longest Collatz chain: %v", longestCollatzNumber)
}
