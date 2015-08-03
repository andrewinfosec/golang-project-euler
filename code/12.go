// https://projecteuler.net/problem=12

package main

import (
	"fmt"
	"math"
)

func divisors(n int64) int64 {
	var c int64 = 0

	for i := int64(1); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			c = c + 2
		}
	}
	return c
}

func main() {
	var i, j int64 = 0, 1
	var next int64

	for {
		next = i + j
		if divisors(next) > 500 {
			fmt.Println(next) // 76576500
			break
		}
		i = next
		j++
	}
}
