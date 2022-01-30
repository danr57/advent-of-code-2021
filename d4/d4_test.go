package d4_test

import (
	"testing"

	"github.com/danr57/advent-of-code-2022/d4"
)

var (
	ans1 = 4512
	ans2 = 0
)

//TODO: Write tests for actual results

func TestPart1(t *testing.T) {
	t.Parallel()

	t.Run("Day 4.1: Sample", func(t *testing.T) {
		t.Parallel()

		r := d4.Part1()

		if r != ans1 {
			t.Fatalf("Expected: %v. Got: %v", ans1, r)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Parallel()

	t.Run("Day 4.2: Sample", func(t *testing.T) {
		t.Parallel()

		r := d4.Part2()

		if r != ans2 {
			t.Fatalf("Expected: %v. Got: %v", ans2, r)
		}
	})
}
