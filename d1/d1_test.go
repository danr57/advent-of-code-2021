package d1_test

import (
	"testing"

	"github.com/danr57/advent-of-code-2022/d1"
)

var sample = []string{
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

func TestPart1(t *testing.T) {
	t.Run("Sample", func(t *testing.T) {

		result := d1.Part1(sample)

		if result != 7 {
			t.Fatal("expected 7, got", result)

		}
	})
}
