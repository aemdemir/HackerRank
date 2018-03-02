package main

import (
	"fmt"
	"math"
)

/*
	Link:
	https://www.hackerrank.com/challenges/encryption/problem
*
	Run:
	go run main.go < test.in
*/

func main() {
	var text string
	fmt.Scanln(&text)

	len := len(text)
	row := int(math.Floor(math.Sqrt(float64(len))))
	var column int

	switch {
	case row*row >= len:
		column = row
	case row*(row+1) >= len:
		column = row + 1
	case (row+1)*(row+1) >= len:
		row, column = row+1, row+1
	}

	var result string
	for c := 0; c < column; c++ {
		for i := c; i < len; i += column {
			result = result + string(text[i])
		}
		result = result + " "
	}

	fmt.Println(result)
}
