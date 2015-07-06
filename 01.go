// https://projecteuler.net/problem=1

package main

import "fmt"

func main() {
	var c int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			c = c + i
		}
	}
	fmt.Println(c) // 233168
}
