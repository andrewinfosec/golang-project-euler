// https://projecteuler.net/problem=10

package main

import (
	"fmt"
)

func main() {
	var n int = 2000000 - 1 //  "below two million"

	list := make([]bool, n+1)
	for i, _ := range list {
		list[i] = true
	}
	list[0], list[1] = false, false // 0 and 1 are not prime numbers

	// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
	begin := 1
	for {
		if begin*begin > n {
			break
		}
	Loop:
		for i, v := range list {
			if v == true && i > begin {
				// sieve multiples of i
				for j := i + i; j <= n; j += i {
					list[j] = false
				}
				begin = i
				break Loop
			}
		}
	}

	var c int64 = 0
	for i, v := range list {
		if v == true {
			c += int64(i)
		}
	}
	fmt.Println(c) // 142913828922
}
