// Package d1 contains the solution for https://adventofcode.com/2021/day/1
package d1

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Part1 Analyzes depth readings to count how many times the readings increase.
func Part1(route []int) int {
	prev := 0
	count := 0

	for _, n := range route[1:] {
		if n > prev {
			count++
		}

		prev = n
	}

	return count
}

// Part2 similarly analyzes readings again but in sliding windows of 3.
func Part2(route []int) int {
	window := 3
	count := 0
	prev := 0

	for index := range route[:len(route)-(window-1)] {
		if index == 0 {
			prev = route[index] + route[index+1] + route[index+2]

			continue
		}

		reading := route[index] + route[index+1] + route[index+2]

		if reading > prev {
			count++
		}

		prev = reading
	}

	return count
}

// ReadAndSplit input file into an []int readable by Part1() and Part2().
func ReadAndSplit(r string) []int {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", r, err)
	}

	p := strings.Split(string(input), "\n")
	report := make([]int, len(p))

	for index, v := range p {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error converting value %v to int: %v", v, err)
		}

		report[index] = value
	}

	return report
}
