package main

import (
	"fmt"
)

/*
What is the value of the first triangle number to have over five hundred divisors?
*/
func main() {
	found := false
	var triangleNumber int64
	iterations := int64(0)
	for !found {
		iterations++
		triangleNumber += iterations

		evenDivisors := 0
		for x := int64(1); x < triangleNumber; x++ {
			if triangleNumber%x == 0 {
				if triangleNumber/x == x {
					evenDivisors++
				} else {
					evenDivisors += 2
					if triangleNumber/x < x {
						break
					}
				}
			}
		}

		if evenDivisors >= 500 {
			found = true
		}
	}

	fmt.Printf("Triangle number with over 500 divisors: %v", triangleNumber)
}
