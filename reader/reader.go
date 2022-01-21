package reader

import (
	"log"
	"os"
	"strings"
)

func ReadAndSplit(s string) []string {
	i, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("Error reading input file:", err)
	}

	if len(i) < 1 {
		log.Fatalf("Error: %v is empty. Please populate and retry", s)
	}

	return strings.Split(string(i), "\n")

}
