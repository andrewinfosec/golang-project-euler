// https://projecteuler.net/problem=16

package main

import (
  "fmt"
  "strconv"
  "math/big"
)

func digitSum(s string) (sum int) {
  sum = 0
  for _, c := range s {
    i, err := strconv.Atoi(string(c))
    if err != nil {
      panic(err)
    }
    sum += i
  }
  return sum
}

func main() {
    n := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()
    fmt.Println(digitSum(n)) // 1366
}
