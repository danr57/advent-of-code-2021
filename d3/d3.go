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

// Part2 verifies the life support rating of the submarine.
func Part2(input []string) int {
	var (
		inputLength  = len(input[0])
		newInput     = input
		oxygenRating int
		co2Rating    int
	)

	newInput := input

	for index := 0; index < inputLength; index++ {
		count0s := 0

		count1s := 0

		for _, num := range newInput {
			switch num[index : index+1] {
			case "0":
				count0s++
			case "1":
				count1s++
			}
		}

		if len(input) > 1 {
			switch {
			case count0s > count1s:
				newInput = purgeValues(input, "0", index)
			case count0s < count1s:
				newInput = purgeValues(input, "1", index)
			case count0s == count1s:
				newInput = purgeValues(input, "1", index)
			}
		}

		input = newInput
		log.Println("Input Length: ", len(input))
	}

	return len(input)
}

func purgeValues(input []string, match string, index int) []string {
	purgedEntries := 0

	for _, num := range input {
		// log.Println("REACHED") // TODO remove debug line

		if num[index:index+1] != match {
			input[index] = input[len(input)-1]
			purgedEntries++
		}
	}

	log.Println("Purged entries:", purgedEntries) // TODO remove debug lines
	log.Println("Current Input Length: ", len(input))

	return input[:len(input)-purgedEntries]
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
