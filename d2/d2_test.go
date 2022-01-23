package d2_test

import (
	"testing"

	"github.com/danr57/advent-of-code-2022/d2"
)

var (
	ans1   = 150
	ans2   = 900
	sample = []d2.Direction{
		{Direction: "forward", Distance: 5},
		{Direction: "down", Distance: 5},
		{Direction: "forward", Distance: 8},
		{Direction: "up", Distance: 3},
		{Direction: "down", Distance: 8},
		{Direction: "forward", Distance: 2}}
)

func TestPart1(t *testing.T) {
	t.Parallel()

	t.Run("Sample", func(t *testing.T) {
		t.Parallel()

		r := d2.Part1(sample)

		if r != ans1 {
			t.Fatalf("Expected: %v. Got: %v.", ans1, r)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Parallel()

	t.Run("Actual", func(t *testing.T) {
		t.Parallel()

		r := d2.Part2(sample)

		if r != ans2 {
			t.Fatalf("Expected: %v. Got: %v.", ans2, r)
		}
	})
}
