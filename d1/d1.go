package d1

import (
	"log"
	"strconv"
)

func Part1(s []string) int {

	x := 0
	var count int

	for _, n := range s[1:] {

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
