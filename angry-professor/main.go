package main

import "fmt"

/*
	Link:
	https://www.hackerrank.com/challenges/angry-professor/problem
*/

/*
	Run:
	go run main.go < test.in
*/

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n, k int
		fmt.Scanf("%d%d", &n, &k)

		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &a[j])
		}

		result := angryProfessor(k, a)
		fmt.Println(result)
	}
}

func angryProfessor(k int, a []int) string {
	onTime := 0
	for _, arrival := range a {
		if arrival <= 0 {
			onTime++
		}
	}
	if onTime < k {
		return "YES"
	}
	return "NO"
}
