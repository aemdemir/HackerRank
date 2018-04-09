package main

import (
	"fmt"
)

/*
	Link:
	https://www.hackerrank.com/challenges/service-lane/problem
*/

/*
	Run:
	go run main.go < test.in
*/

func main() {
	var n, t int
	fmt.Scanf("%d%d", &n, &t)

	widths := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &widths[i])
	}

	for i := 0; i < t; i++ {
		var i, j int
		fmt.Scanf("%d%d", &i, &j)

		result := test(widths[i : j+1])
		fmt.Println(result)
	}
}

func test(widths []int) int {
	return min(widths)
}

func min(slice []int) int {
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if min > slice[i] {
			min = slice[i]
		}
	}
	return min
}
