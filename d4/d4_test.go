package d4_test

import (
	"log"
	"testing"

	"github.com/danr57/advent-of-code-2022/d4"
)

var (
	ans1 = 4512
	ans2 = 0
	sol1 = 23177
	sol2 = 6804
)

func TestPart1(t *testing.T) {
	t.Parallel()

	t.Run("Day 4.1: Sample", func(t *testing.T) {
		t.Parallel()

		if r := d4.Part1(ans1); r != ans1 {
			t.Fatalf("Expected: %v. Got: %v", ans1, r)
		}
	})

	t.Run("Day4.1: Actual: ", func(t *testing.T) {
		t.Parallel()

		if r := d4.Part1(sol1); r != sol1 {
			log.Fatalf("Expected: %v. Got: %v.", sol1, r)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Parallel()

	t.Run("Day 4.2: Sample", func(t *testing.T) {
		t.Parallel()

		if r := d4.Part2(ans2); r != ans2 {
			t.Fatalf("Expected: %v. Got: %v", ans2, r)
		}
	})

	t.Run("Day 4.2: Actual: ", func(t *testing.T) {
		t.Parallel()

		if r := d4.Part2(sol2); r != sol2 {
			log.Fatalf("Expected: %v. Got: %v.", sol2, r)
		}
	})
}
