package main

import (
	"fmt"
	"math"
)

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
