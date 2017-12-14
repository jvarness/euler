package main

import (
	"fmt"
)
/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func main() {
	for x := 20; ; x++ {
		var y int
		for y = 1; y < 20; y++ {
			if x%y != 0 {
				break
			}
		}
		if y == 20 {
			fmt.Println(x)
		}
	}
}
