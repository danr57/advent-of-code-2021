package d3

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	inputLength int
)

func BinaryStrToInt64(in string) int64 {
	base := 2
	bitSize := 64

	val, err := strconv.ParseInt(in, base, bitSize)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func Part1(r []string) int {
	var gamma bytes.Buffer
	var epsilon bytes.Buffer

	for i := 0; i < inputLength; i++ {
		count0s := 0
		count1s := 0

		for _, num := range r {
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
	log.Println(gammaDecimal * epsilonDecimal)
	return int(gammaDecimal * epsilonDecimal)
}

func Part2(r []string) int {
	return 0
}

// Read input file and split line separated values into an []int readable by Part1() and Part2()
func ReadAndSplit(r string) []string {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file \"%v\": %v", r, err)
	}

	if len(input) < 1 {
		log.Fatalf("Error: file '%v' is empty. Please populate. err:", err)
	}

	report := strings.Split(string(input), "\n")

	inputLength = len(report[0])

	return report

}
