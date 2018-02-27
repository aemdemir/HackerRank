package main

import (
	"fmt"
	"os"
	"unicode"
)

/*
	https://www.hackerrank.com/challenges/camelcase/problem
*/

func main() {
	var result = 1
	var s string
	fmt.Scanf("%s\n", &s)
	for _, r := range s {
		if unicode.IsUpper(r) {
			result++
		}
	}
	fmt.Fprintln(os.Stdout, result)
}
