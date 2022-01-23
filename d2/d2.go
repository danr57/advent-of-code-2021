package d2

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Describes a move of the submarine.
type Move struct {
	Direction string
	Distance  int
}

func Part1(route []Move) int {
	x := 0
	y := 0

	for _, d := range route {
		switch d.Direction {
		case "up":
			y -= d.Distance
		case "down":
			y += d.Distance
		case "forward":
			x += d.Distance
		}
	}

	return x * y
}

func Part2(route []Move) int {
	a := 0
	x := 0
	y := 0

	for _, d := range route {
		switch d.Direction {
		case "up":
			a += d.Distance
		case "down":
			a -= d.Distance
		case "forward":
			x += d.Distance
			y -= (a * d.Distance)
		}
	}

	return x * y
}

// Formats input file into a submarine route as []Moves.
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

	for i, dir := range list {
		split := strings.Split(string(dir), " ")

		distance, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Error converting distance %v to int. Error: %v", split[1], err)
		}

		step := Move{split[0], distance}

		route[i] = step
	}

	return route
}
