package main

import (
	"fmt"
)

/*
	Link:
	https://www.hackerrank.com/challenges/almost-sorted/problem
*
	Run:
	go run main.go < test.in
*/

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var l, r int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			if l == 0 {
				l = i
			}
			r = i + 1
		}
	}

	arr2 := make([]int, n)
	copy(arr2, arr)

	swap(l, r, arr)
	if sorted(arr) {
		fmt.Println("yes")
		fmt.Println("swap", l+1, r+1)
		return
	}

	reverse(l, r, arr2)
	if sorted(arr2) {
		fmt.Println("yes")
		fmt.Println("reverse", l+1, r+1)
		return
	}

	fmt.Println("no")
}

func swap(l, r int, arr []int) {
	arr[l], arr[r] = arr[r], arr[l]
}

func reverse(l, r int, arr []int) {
	for i := 0; i < (r-l+1)/2; i++ {
		swap(l+i, r-i, arr)
	}
}

func sorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
