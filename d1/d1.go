package d1

import (
	"log"
	"strconv"
)

var (
	count int = 0
	x     int = 0
)

func Part1(s []string) int {
	for _, n := range s {

		t, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("Error converting strings to int:", err)
		}

		if t > x {
			count++
		}
		x = t
	}

	return count
}
