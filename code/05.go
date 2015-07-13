// https://projecteuler.net/problem=5

package main

import (
	"fmt"
)

func main() {
	numbers := []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11}

	c := 2520
Loop:
	for {
		c++
		for _, n := range numbers {
			if c%n != 0 {
				continue Loop
			}
		}
		break
	}
	fmt.Println(c) // 232792560
}
