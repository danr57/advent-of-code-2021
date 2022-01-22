package d1

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Analyses depth readings to count how many times the readings increase
func Part1(r []int) int {

	prev := 0
	var count int

	for _, n := range r[1:] {

		if n > prev {
			count++
		}
		prev = n
	}

	return count
}

// Same as Part1() but analyses readings in sliding windows of 3
func Part2(r []int) int {
	window := 3
	var count int
	var prev int

	for i := range r[:len(r)-(window-1)] {
		if i == 0 {
			prev = r[i] + r[i+1] + r[i+2]
			continue
		}

		reading := r[i] + r[i+1] + r[i+2]

		if reading > prev {
			count++
		}
		prev = reading

	}
	log.Println(count)
	return count
}

// Read input file and split line separated values into an []int readable by Part1() and Part2()
func ReadAndSplit(r string) []int {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", r, err)
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
