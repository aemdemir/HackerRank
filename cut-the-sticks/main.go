package main

import (
	"fmt"
	"math"
)

/*
	Link:
	https://www.hackerrank.com/challenges/cut-the-sticks/problem
*
	Run:
	go run main.go < test.in
*/

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int32, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	result := cutTheSticks(arr)
	for _, val := range result {
		fmt.Println(val)
	}
}

func cutTheSticks(arr []int32) []int32 {
	result := make([]int32, 0)

	for {
		cutLength := int32(math.MaxInt32)
		for _, val := range arr {
			if val == 0 {
				continue
			}
			if cutLength > val {
				cutLength = val
			}
		}

		if cutLength == int32(math.MaxInt32) {
			break
		}

		sticksCut := int32(0)
		for i, val := range arr {
			if val != 0 {
				arr[i] = val - cutLength
				sticksCut++
			}
		}

		result = append(result, sticksCut)
	}

	return result
}
