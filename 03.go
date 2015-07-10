// https://projecteuler.net/problem=3

package main

import "fmt"

func main() {
	var n int64 = 600851475143
	var d int64 = 2

	// https://en.wikipedia.org/wiki/Trial_division
	for d <= n {
		if n%d == 0 {
			n = n / d // sieve multiples of d
		} else {
			if d > 2 {
				d += 2 // next odd number
			} else {
				d = 3
			}
		}
	}
	fmt.Println(d) // 6857
}
