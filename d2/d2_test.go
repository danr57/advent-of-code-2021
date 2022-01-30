package d2_test

import (
	"log"
	"testing"

	"github.com/danr57/advent-of-code-2022/d2"
)

var (
	ans1   = 150
	ans2   = 900
	sample = []d2.Move{
		{Direction: "forward", Distance: 5},
		{Direction: "down", Distance: 5},
		{Direction: "forward", Distance: 8},
		{Direction: "up", Distance: 3},
		{Direction: "down", Distance: 8},
		{Direction: "forward", Distance: 2}}
	actual = d2.ReadAndSplit("input")
	sol1   = 1990000
	sol2   = 1975421260
)

func TestPart1(t *testing.T) {
	t.Parallel()

	t.Run("Day 3.1: Sample", func(t *testing.T) {
		t.Parallel()

		if r := d2.Part1(sample); r != ans1 {
			t.Fatalf("Expected: %v. Got: %v.", ans1, r)
		}
	})

	t.Run("Day 3.1: Actual", func(t *testing.T) {
		t.Parallel()

		if r := d2.Part1(actual); r != sol1 {
			log.Fatalf("Expedted: %v. Got: %v.", sol1, r)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Parallel()

	t.Run("Day 3.2: Sample: ", func(t *testing.T) {
		t.Parallel()

		if r := d2.Part2(sample); r != ans2 {
			t.Fatalf("Expected: %v. Got: %v.", ans2, r)
		}
	})

	t.Run("Day 3.2: Actual: ", func(t *testing.T) {
		t.Parallel()

		if r := d2.Part2(actual); r != sol2 {
			log.Fatalf("Expected: %v. Got: %v.", sol2, r)
		}
	})
}
