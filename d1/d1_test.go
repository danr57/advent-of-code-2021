package d1_test

import (
	"testing"

	"github.com/danr57/advent-of-code-2022/d1"
)

var (
	ans    = 7
	sample = []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	input = []string{}
	sol   = 1759
)

func TestPart1(t *testing.T) {
	t.Run("Sample", func(t *testing.T) {

		r := d1.Part1(sample)

		if r != ans {
			t.Fatalf("Expected: %v. Got: %v", ans, r)

		}
	})

	t.Run("ok fr this time", func(t *testing.T) {
		r := d1.Part1(input)

		if r != sol {
			t.Fatalf("Expected: %v. Got: %v", sol, r)

		}

	})
}
