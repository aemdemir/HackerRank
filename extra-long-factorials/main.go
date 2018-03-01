package main

import (
	"fmt"
	"math/big"
)

/*
	https://www.hackerrank.com/challenges/extra-long-factorials/problem
*/

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(factorial(n))
}

func factorial(n int) *big.Int {
	if n == 1 {
		return big.NewInt(1)
	}
	return mul(big.NewInt(int64(n)), factorial(n-1))
}

func mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
