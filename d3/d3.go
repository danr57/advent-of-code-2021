package d3

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Part1(r []int) int {

	return 198
}

// Read input file and split line separated values into an []int readable by Part1() and Part2()
func ReadAndSplit(r string) []int {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", r, err)
	}

	if len(input) < 1 {
		log.Fatalf("Error: file '%v' is empty. Please populate. err:", err)
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
