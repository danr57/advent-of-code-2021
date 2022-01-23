package main

import (
	"log"

	"github.com/danr57/advent-of-code-2022/d1"
	"github.com/danr57/advent-of-code-2022/d2"
)

func main() {

	d1input := d1.ReadAndSplit("d1/input")
	log.Printf("===Day 1=== \nPart 1: %v \nPart 2: %v", d1.Part1(d1input), d1.Part2(d1input))

	d2input := d2.ReadAndSplit("d2/input")
	log.Printf("===Day 2=== \nPart 1: %v \nPart 2: %v", d2.Part1(d2input), d2.Part2(d2input))
}
