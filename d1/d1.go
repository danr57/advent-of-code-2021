package d1

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Part1(s []int) int {

	x := 0
	var count int

	for _, n := range s[1:] {

		if n > x {
			count++
		}
		x = n
	}

	return count
}

func Part2(s []int) int {

	return 69
}

func ReadAndSplit(s string) []int {
	input, err := os.ReadFile(filepath.Clean(s))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", s, err)
	}

	p := strings.Split(string(input), "\n")
	report := make([]int, len(p))

	for i, v := range p {

		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error converting value %v to int: %v", v, err)
		}

		report[i] = value

	}

	return report

}
