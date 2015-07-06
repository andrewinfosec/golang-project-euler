// https://projecteuler.net/problem=2

package main

import "fmt"

func main() {
	var count int

	a, b := 1, 2
	count = 2 // because b%2 == 0

	for {
		next := a + b
		if next > 4000000 {
			break
		}
		if next%2 == 0 {
			count = count + next
		}
		a, b = b, next
	}
	fmt.Println(count) // 4613732
}
