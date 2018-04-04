package main

import (
	"fmt"
	"math"
)

/*
	Link:
	https://www.hackerrank.com/challenges/sherlock-and-squares/problem
*/

/*
	Run:
	go run main.go < test.in
*/

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scanf("%d%d\n", &a, &b)
		result := squares(a, b)
		fmt.Println(result)
	}
}

func squares(a, b int) int {
	result := math.Floor(math.Sqrt(float64(b))) - math.Ceil(math.Sqrt(float64(a))) + 1
	return int(result)
}
