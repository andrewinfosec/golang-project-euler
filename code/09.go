// https://projecteuler.net/problem=9

package main

import (
	"fmt"
)

func isTriplet(a, b, c int) bool {
	return a < b && b < c && a*a+b*b == c*c
}

// a^2 + b^2 = c^2 therefore a + b > c; a + b + c = 1000 therefore c <= 500
func main() {

Loop:
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			for c := 1; c <= 500; c++ {
				if isTriplet(a, b, c) {
					if a+b+c == 1000 {
						fmt.Println(a * b * c) // 31875000
						break Loop
					}
				}
			}
		}
	}
}
