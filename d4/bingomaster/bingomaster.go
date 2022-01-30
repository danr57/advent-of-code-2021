package bingomaster

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Board type stores each bingo scorecard.
type Board struct {
	tiles         [5][5][2]int
	winner        bool
	winningNumber int
}

// ReadAndSplit processes input into draw list and Boards.
func ReadAndSplit(r string) ([]int, []Board) {
	input, err := os.ReadFile(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file '%v': %v", r, err)
	}

	p := strings.Split(string(input[0]), " ")

	draws := make([]int, len(p))

	for i, n := range p {
		draws[i], _ = strconv.Atoi(n)
	}

	tmp := [5][5][2]int{{{1, 2}, {2, 3}, {3, 4}}}

	var bob []Board

	board := Board{tmp, true, 69}
	bob = append(bob, board)

	return draws, bob
}
