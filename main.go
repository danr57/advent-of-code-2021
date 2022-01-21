package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/danr57/advent-of-code-2022/d1"
)

func main() {
	log.Println("Day 1")
	d1input := readAndSplit("d1/input")
	log.Println(d1.Part1(d1input))

}

func readAndSplit(s string) []string {
	i, err := ioutil.ReadFile(s)

	if err != nil {
		log.Fatal("Error reading input file:", err)
	}

	return strings.Split(string(i), "\n")

}
