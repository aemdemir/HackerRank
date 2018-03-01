package main

import (
	"fmt"
	"os"
)

/*
	https://www.hackerrank.com/challenges/caesar-cipher-1/problem
*/

func main() {
	var length, rotation int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &rotation)

	var result string
	for _, r := range input {
		switch {
		case 'a' <= r && r <= 'z':
			r = rotate(r, rotation, 'a', 'z')
		case 'A' <= r && r <= 'Z':
			r = rotate(r, rotation, 'A', 'Z')
		}
		result = result + string(r)
	}

	fmt.Fprintln(os.Stdout, result)
}

func rotate(r rune, rotation int, first rune, last rune) rune {
	length := last - first + 1
	distToFirst := r + rune(rotation) - first
	return first + (distToFirst % length)
}
