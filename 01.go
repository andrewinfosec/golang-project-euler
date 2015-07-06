// https://projecteuler.net/problem=1

package main

import "fmt"

func main() {
	var count int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			count = count + i
		}
	}
	fmt.Println(count) // 233168
}
