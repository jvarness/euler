package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
func main() {
	var num = 0
	for x := 999; x >= 100; x-- {
		for y := 999; y >= 100; y-- {
			var product = x * y
			var productstr = strconv.Itoa(x * y)
			if strings.Compare(productstr, reverse(productstr)) == 0 && product > num {
				num = product
			}
		}
	}

	fmt.Println(num)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
