package main

import (
	"fmt"
	"math"
)

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
func main() {
	var a, b, c float64
	for a = float64(1); a <= 1000; a++ {
		for b = float64(1); b <= 1000; b++ {
			c = math.Sqrt(math.Pow(a, float64(2)) + math.Pow(b, float64(2)))
			if a+b+c == 1000 {
				break
			}
		}
		if a+b+c == 1000 {
			break
		}
	}

	fmt.Println(a * b * c)
	fmt.Printf("a: %v, b: %v, c: %v", a, b, c)
}
