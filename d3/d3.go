// Package d3 contains the solution for https://adventofcode.com/2021/day/3
package d3

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// BinaryStrToInt64 Gets decimal value of binary value in string format.
func BinaryStrToInt64(in string) int64 {
	base := 2
	bitSize := 64

	val, err := strconv.ParseInt(in, base, bitSize)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

/* Part1 Finds most/least common digit at each index for all readings,
returns multiplied decimal versions of these 2 values.*/
func Part1(input []string) int {
	var gamma bytes.Buffer // most common digits per index

	var epsilon bytes.Buffer // least common digits per index

	inputLength := len(input[0])

	for i := 0; i < inputLength; i++ {
		count0s := 0
		count1s := 0

		for _, num := range input {
			switch num[i : i+1] {
			case "0":
				count0s++
			case "1":
				count1s++
			}
		}

		if count0s > count1s {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		} else {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		}
	}

	gammaDecimal := BinaryStrToInt64(gamma.String())
	epsilonDecimal := BinaryStrToInt64(epsilon.String())

	return int(gammaDecimal * epsilonDecimal)
}

// Part2 will contain the solution for Day 3 Part 2 of AOC2021.
// TODO: Complete Part2 solution.
func Part2(r []string) int {
	return 0
}

/* ReadAndSplit takes input file and splits line separated values
into an []int readable by Part1() and Part2().*/
func ReadAndSplit(r string) []string {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", r, err)
	}

	if len(input) < 1 {
		log.Fatalf("Error: file '%v' is empty. Please populate. err:", err)
	}

	report := strings.Split(string(input), "\n")

	return report
}
