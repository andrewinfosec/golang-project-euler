// https://projecteuler.net/problem=14

package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Collatz_conjecture
func findTerms(n int64) int64 {
	var c int64 = 1

	for {
		if n%2 == 0 { // even
			n = n / 2
		} else { // odd
			n = (3 * n) + 1
		}
		c++
		if n == 1 {
			return c
		}
	}
}

func main() {
	type answer struct {
		number    int64
		termCount int64
	}
	a := answer{0, 0}

	for i := int64(999999); i > 0; i-- { // "under one million"
		if termCount := findTerms(i); termCount > a.termCount {
			a.termCount = termCount
			a.number = i
		}
	}
	fmt.Println(a.number) // 837799 (a.termCount == 525)
}
