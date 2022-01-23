package d3_test

import (
	"log"
	"testing"

	"github.com/danr57/advent-of-code-2022/d3"
)

var (
	ans1   = 198
	sample = []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
)

func TestPart1(t *testing.T) {
	t.Parallel()

	if r := d3.Part1(sample); r != ans1 {
		log.Fatalf("Expected: %v. Got: %v", ans1, r)
	}
}

func TestBinaryStrToInt64(t *testing.T) {
	t.Parallel()

	if r := d3.BinaryStrToInt64("10110"); r != 22 {
		log.Fatalf("Expected: %v. Got: %v", 22, r)
	}
}
