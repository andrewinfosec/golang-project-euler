// https://projecteuler.net/problem=15

package main

import "fmt"

const gridSize int = 20

func main() {
	var paths int = 1

	// "2*n choose n" becomes "40 choose 20"
	// https://en.wikipedia.org/wiki/Binomial_coefficient
	for i := 0; i < gridSize; i++ {
		paths = paths * ((2 * gridSize) - i)
		paths = paths / (i + 1)
	}
	fmt.Println(paths) // 137846528820
}
