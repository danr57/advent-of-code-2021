package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/danr57/advent-of-code-2022/d1"
)

func main() {

	d1input := ReadAndSplit("d1/input")
	log.Printf("===Day 1=== \nPart 1: %v", d1.Part1(d1input))

}

func ReadAndSplit(s string) []string {
	i, err := ioutil.ReadFile(s)

	if len(i) < 1 {
		log.Fatalf("Error: %v is empty. Please populate and retry", s)
	}

	if err != nil {
		log.Fatal("Error reading input file:", err)
	}

	return strings.Split(string(i), "\n")

}
