package main

import (
	"fmt"
)

/*
	Link:
	https://www.hackerrank.com/challenges/chocolate-feast/problem
*
	Run:
	go run main.go < test.in
*/

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n, c, m int
		fmt.Scanf("%d%d%d", &n, &c, &m)

		result := chocolateFeast(n, c, m)
		fmt.Println(result)
	}
}

func chocolateFeast(n int, c int, m int) int {
	chocolate := n / c
	wrappers := chocolate

	for wrappers >= m {
		r := wrappers % m
		wrappers = wrappers / m

		chocolate += wrappers
		wrappers += r
	}

	return chocolate
}
