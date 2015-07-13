// https://projecteuler.net/problem=6

package main

import (
	"fmt"
)

func main() {
	sumOfSquares, sum := 0, 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		sum += i
	}
	squareOfSum := sum * sum

	fmt.Println(squareOfSum - sumOfSquares) //25164150
}
