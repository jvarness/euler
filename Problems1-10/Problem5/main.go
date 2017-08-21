package main

import (
	"fmt"
)

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
