package main

import (
	"fmt"
)

/*
	https://www.hackerrank.com/challenges/day-of-the-programmer/problem
*/

// Calendar is an enumerator for Calendar systems
type Calendar int

// Calendars
const (
	Julian Calendar = iota
	Gregorian
)

const (
	dayOfTheProgrammer = 256
)

func main() {
	var year int
	fmt.Scanf("%d\n", &year)

	if year == 1918 {
		fmt.Println("26.09.1918")
		return
	}

	days := buildDaysIn(year)

	var dd int
	var mm string
	for month, daysPassed := range days {
		daysLeft := dayOfTheProgrammer - daysPassed

		if daysLeft <= 0 {
			dd = dayOfTheProgrammer - days[month-1]
			mm = fmt.Sprint(month)
			if month < 10 { mm = fmt.Sprintf("0%d", month) }
			break
		}
	}

	fmt.Printf("%d.%s.%d\n", dd, mm, year)
}

func buildDaysIn(year int) [13]int {
	var leapYear bool

	switch {
	case year < 1918:
		leapYear = isLeap(year, Julian)
	case 1918 < year:
		leapYear = isLeap(year, Gregorian)
	}

	feb := 28
	if leapYear {
		feb = 29
	}

	return [13]int{
		0,
		31,
		31 + feb,
		31 + feb + 31,
		31 + feb + 31 + 30,
		31 + feb + 31 + 30 + 31,
		31 + feb + 31 + 30 + 31 + 30,
		31 + feb + 31 + 30 + 31 + 30 + 31,
		31 + feb + 31 + 30 + 31 + 30 + 31 + 31,
		31 + feb + 31 + 30 + 31 + 30 + 31 + 31 + 30,
		31 + feb + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
		31 + feb + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
		31 + feb + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
	}
}

func isLeap(year int, calendar Calendar) bool {
	switch calendar {
	case Julian:
		return year%4 == 0
	case Gregorian:
		return year%4 == 0 && (year%100 != 0 || year%400 == 0)
	}
	panic("Unknown calendar!")
}
