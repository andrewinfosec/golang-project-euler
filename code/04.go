// https://projecteuler.net/problem=4

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for i, upper := 0, len(s)-1; i < len(s)/2; i++ {
		if s[i] != s[upper] {
			return false
		}
		upper--
	}
	return true
}

func main() {
	largest := 0
	for a := 100; a < 1000; a++ {
		for b := 100; b < 1000; b++ {
			product := a * b
			if isPalindrome(product) && product > largest {
				largest = product
			}
		}
	}
	fmt.Println(largest) // 906609
}
