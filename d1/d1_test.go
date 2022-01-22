package d1_test

import (
	"testing"

	"github.com/danr57/advent-of-code-2022/d1"
)

var (
	ans1   = 7
	ans2   = 5
	sample = []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	input = d1.ReadAndSplit("input")
	sol1  = 1759
	sol2  = 1805
)

func TestPart1(t *testing.T) {
	t.Parallel()

	t.Run("Sample", func(t *testing.T) {
		t.Parallel()

		r := d1.Part1(sample)

		if r != ans1 {
			t.Fatalf("Expected: %v. Got: %v", ans1, r)

		}
	})

	t.Run("Actual", func(t *testing.T) {
		t.Parallel()
		r := d1.Part1(input)

		if r != sol1 {
			t.Fatalf("Expected: %v. Got: %v", sol1, r)

		}

	})
}

func TestPart2(t *testing.T) {
	t.Parallel()

	t.Run("Sample", func(t *testing.T) {
		t.Parallel()

		r := d1.Part2(sample)

		if r != ans2 {
			t.Fatalf("Expected: %v. Got: %v", ans2, r)
		}
	})

	t.Run("Actual input", func(t *testing.T) {
		t.Parallel()

		r := d1.Part2(input)

		if r != sol2 {
			t.Fatalf("Expected: %v. Got: %v", sol2, r)
		}
	})
}
