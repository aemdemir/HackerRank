package main

import (
	"fmt"
)

/*
	Link:
	https://www.hackerrank.com/challenges/library-fine/problem

/*
	Run:
	go run main.go < test.in
*/

// Date holds relevant information about a date.
type Date struct {
	Day   int
	Month int
	Year  int
}

func main() {
	actual := Date{}
	expected := Date{}

	fmt.Scanf("%d%d%d", &actual.Day, &actual.Month, &actual.Year)
	fmt.Scanf("%d%d%d", &expected.Day, &expected.Month, &expected.Year)

	fine := calculateFine(actual, expected)
	fmt.Println(fine)
}

func calculateFine(actual, expected Date) int {
	switch {
	case actual.Year != expected.Year:
		if actual.Year > expected.Year {
			return 10000
		}
		return 0
	case actual.Month != expected.Month:
		if actual.Month > expected.Month {
			return 500 * (actual.Month - expected.Month)
		}
		return 0
	case actual.Day != expected.Day:
		if actual.Day > expected.Day {
			return 15 * (actual.Day - expected.Day)
		}
		return 0
	default:
		return 0
	}
}
