package main

import "fmt"

/*
	Link:
	https://www.hackerrank.com/challenges/the-time-in-words/problem
*/

/*
	Run:
	go run main.go < test.in
*/

var time = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	21: "twenty one",
	22: "twenty two",
	23: "twenty three",
	24: "twenty four",
	25: "twenty five",
	26: "twenty six",
	27: "twenty seven",
	28: "twenty eight",
	29: "twenty nine",
}

func main() {
	var h, m int
	fmt.Scanf("%d", &h)
	fmt.Scanf("%d", &m)

	result := timeInWords(h, m)
	fmt.Println(result)
}

func timeInWords(h int, m int) string {
	hour, nextHour := time[h], time[h+1]

	switch {
	case m == 0:
		return hour + " o' clock"
	case m == 1:
		return "one minute past " + hour
	case m == 15:
		return "quarter past " + hour
	case m == 30:
		return "half past " + hour
	case m == 45:
		return "quarter to " + nextHour
	case m < 30:
		return time[m] + " minutes past " + hour
	case m < 60:
		return time[60-m] + " minutes to " + nextHour
	}

	return ""
}
