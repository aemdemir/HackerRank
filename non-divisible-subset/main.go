package main

import (
	"fmt"
)

/*
	Link:
	https://www.hackerrank.com/challenges/non-divisible-subset/problem
*/

/*
	Run:
	go run main.go < test.in
*/

func main() {
	var n, k int
	fmt.Scanf("%d%d\n", &n, &k)

	var remainders = make([]int, k)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		remainders[a%k]++
	}

	count := 0
	for i := 1; i <= k/2; i++ {
		if i == k-i {
			count++
			continue
		}
		count += max(remainders[i], remainders[k-i])
	}
	if remainders[0] > 0 {
		count++
	}

	fmt.Println(count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
