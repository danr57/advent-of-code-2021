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
func Part1(input []string) int64 {
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

	return gammaDecimal * epsilonDecimal
}

// Part2 verifies the life support rating of the submarine.
func Part2(input []string) int64 {
	var (
		oxygenRating = input
		co2Rating    = input
	)

	for index := 0; index < len(input[0]); index++ {
		oxygenRating = findRatings(oxygenRating, true, index)
		co2Rating = findRatings(co2Rating, false, index)
	}

	return BinaryStrToInt64(oxygenRating[0]) * BinaryStrToInt64(co2Rating[0])
}

// finds most common bit at index position and purges it.
func findRatings(input []string, keep bool, index int) []string {
	count0s := 0
	count1s := 0

	for _, num := range input {
		switch num[index : index+1] {
		case "0":
			count0s++
		case "1":
			count1s++
		}
	}

	if len(input) > 1 {
		if count0s > count1s {
			input = purgeValues(input, "0", index, keep)
		} else {
			input = purgeValues(input, "1", index, keep)
		}
	}

	return input
}

func purgeValues(input []string, match string, index int, keep bool) []string {
	var result []string

	for _, num := range input {
		if len(input) <= 1 {
			result = append(result, num)

			break
		}

		if (num[index:index+1] == match) == keep {
			result = append(result, num)
		}
	}

	return result
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
