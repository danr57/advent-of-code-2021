package main

import (
	"log"

	"github.com/danr57/advent-of-code-2021/d1"
	"github.com/danr57/advent-of-code-2021/d2"
	"github.com/danr57/advent-of-code-2021/d3"
	"github.com/danr57/advent-of-code-2021/d4"
)

func main() {
	d1input := d1.ReadAndSplit("d1/input")
	log.Printf("DAY 1 { Part 1: %v. Part 2: %v. }", d1.Part1(d1input), d1.Part2(d1input))

	d2input := d2.ReadAndSplit("d2/input")
	log.Printf("DAY 2 { Part 1: %v. Part 2: %v. }", d2.Part1(d2input), d2.Part2(d2input))

	d3input := d3.ReadAndSplit("d3/input")
	log.Printf("DAY 3 { Part 1: %v. Part 2: %v. }", d3.Part1(d3input), d3.Part2(d3input))

	// draws, bob := bingomaster.ReadAndSplit("d4.input")

	log.Printf("DAY 4 { Part 1: %v. Part 2: %v, }", d4.Part1(120), d4.Part2(120))
}
