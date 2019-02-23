package main

import (
	"fmt"
)

/*
	Link:
	https://www.hackerrank.com/challenges/acm-icpc-team/problem
*
	Run:
	go run main.go < test.in
*/

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)

	attendees := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &attendees[i])
	}

	maxTopics := 0
	maxTeams := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			topics := 0
			for k := 0; k < m; k++ {
				// 49 is the representation of 1 in byte.
				if attendees[i][k] == 49 || attendees[j][k] == 49 {
					topics++
				}
			}
			if topics > maxTopics {
				maxTopics = topics
				maxTeams = 1
				continue
			}
			if topics != 0 && topics == maxTopics {
				maxTeams++
			}
		}
	}

	fmt.Println(maxTopics)
	fmt.Println(maxTeams)
}
