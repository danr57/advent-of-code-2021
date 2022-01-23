package d2

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Part1(route []string) int {
	x := 0
	y := 0

	for _, d := range route {
		t := strings.Split(string(d), " ")
		n, _ := strconv.Atoi(t[1])
		switch t[0] {
		case "up":
			y -= n
		case "down":
			y += n
		case "forward":
			x += n
		case "backward":
			x -= n
		}
	}
	return x * y
}

func Part2(route []string) int {
	a := 0
	x := 0
	y := 0

	for _, d := range route {
		t := strings.Split(string(d), " ")
		n, _ := strconv.Atoi(t[1])
		switch t[0] {
		case "up":
			a -= n
		case "down":
			a += n
		case "forward":
			x += n
			y -= (a * n)
		case "backward":
			x -= n
		}
	}

	// TODO: This is returning the answer as a negative int
	// TODO: Problem demands a positive result
	return x * y
}

// TODO: Process input to a better structure
func ReadAndSplit(r string) []string {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", r, err)
	}

	if len(input) < 1 {
		log.Fatalf("Data input empty. Please populate")
	}

	return strings.Split(string(input), "\n")
}
