// Package d4 contains the solution for https://adventofcode.com/2021/day/4
package d4

import (
	"log"

	"github.com/danr57/advent-of-code-2022/d4/bingomaster"
)

// Part1 will find a winning bingo board and calculate its score.
func Part1(ans int) int {
	x, y := bingomaster.ReadAndSplit("d4/input")
	log.Printf("x: %v", x)
	log.Printf("y: %v", y)

	return ans
}

// Part2 TODO Finish.
func Part2(ans int) int {
	return ans
}
