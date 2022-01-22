package main

import (
	"log"

	"github.com/danr57/advent-of-code-2022/d1"
)

func main() {

	d1input := d1.ReadAndSplit("d1/input")
	log.Printf("===Day 1=== \nPart 1: %v \nPart 2: %v", d1.Part1(d1input), d1.Part2(d1input))

}
