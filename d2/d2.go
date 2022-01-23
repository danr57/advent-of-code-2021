// Package d2 contains the solution for https://adventofcode.com/2021/day/2
package d2

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Move describes a move of the submarine during its route.
type Move struct {
	Direction string
	Distance  int
}

// Part1 takes the planned route of submarine and calculates final destination.
func Part1(route []Move) int {
	travel := 0
	depth := 0

	for _, move := range route {
		switch move.Direction {
		case "up":
			depth -= move.Distance
		case "down":
			depth += move.Distance
		case "forward":
			travel += move.Distance
		}
	}

	return travel * depth
}

// Part2 calculates final destination of the sub by adjusting aim over depth.
func Part2(route []Move) int {
	aim := 0
	movement := 0
	depth := 0

	for _, move := range route {
		switch move.Direction {
		case "up":
			aim += move.Distance
		case "down":
			aim -= move.Distance
		case "forward":
			movement += move.Distance
			depth -= (aim * move.Distance)
		}
	}

	return movement * depth
}

// ReadAndSplit input file into a submarine route as array of Moves.
func ReadAndSplit(r string) []Move {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", r, err)
	}

	if len(input) < 1 {
		log.Fatalf("Data input empty. Please populate")
	}

	list := strings.Split(string(input), "\n")

	route := make([]Move, len(list))

	for index, dir := range list {
		split := strings.Split(dir, " ")

		distance, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Error converting distance %v to int. Error: %v", split[1], err)
		}

		step := Move{split[0], distance}

		route[index] = step
	}

	return route
}
