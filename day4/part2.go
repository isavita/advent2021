package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day4/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(string(file))

	result := solve(input)
	fmt.Println(result)
}

func solve(input string) int {
	nums, boards := parseInput(input)

	lastWinningScore := -1
	alreadyWon := map[int]bool{}
	for _, n := range nums {
		for bi, b := range boards {
			if alreadyWon[bi] {
				continue
			}
			didWin := b.PickNum(n)
			if didWin {
				// WHICH BOARD WINS LAST
				lastWinningScore = b.Score() * n

				// mark board as already won
				alreadyWon[bi] = true
			}
		}
	}

	return lastWinningScore

}

// BoardState maintains a parsed board and a boolean matrix of cells that have
// been picked/marked
type BoardState struct {
	board  [][]int
	picked [][]bool
}

func NewBoardState(board [][]int) BoardState {
	picked := make([][]bool, len(board))
	for i := range picked {
		picked[i] = make([]bool, len(board[0]))
	}
	return BoardState{
		board:  board,
		picked: picked,
	}
}

func (b *BoardState) PickNum(num int) bool {
	for r, rows := range b.board {
		for c, v := range rows {
			if v == num {
				b.picked[r][c] = true
			}
		}
	}

	// is this fast enough to do on every "cycle"?
	// guess so. probably a constant time way to do this but oh well
	for i := 0; i < len(b.board); i++ {
		isFullRow, isFullCol := true, true
		// board is square so this works fine, otherwise would need another pair of nested loops
		for j := 0; j < len(b.board); j++ {
			// check row at index i
			if !b.picked[i][j] {
				isFullRow = false
			}
			// check col at index j
			if !b.picked[j][i] {
				isFullCol = false
			}
		}
		if isFullRow || isFullCol {
			// returns true if is winning board
			return true
		}
	}

	// false for incomplete board
	return false
}

func (b *BoardState) Score() int {
	var score int

	for r, rows := range b.board {
		for c, v := range rows {
			// adds up all the non-picked/marked cells
			if !b.picked[r][c] {
				score += v
			}
		}
	}

	return score
}

func parseInput(input string) (nums []int, boards []BoardState) {
	lines := strings.Split(input, "\n\n")

	for _, v := range strings.Split(lines[0], ",") {
		nums = append(nums, toInt(v))
	}

	for _, grid := range lines[1:] {
		b := [][]int{}
		for _, line := range strings.Split(grid, "\n") {
			line = strings.ReplaceAll(line, "  ", " ")
			for line[0] == ' ' {
				line = line[1:]
			}
			parts := strings.Split(line, " ")

			row := []int{}
			for _, p := range parts {
				row = append(row, toInt(p))
			}
			b = append(b, row)
		}

		boards = append(boards, NewBoardState(b))
	}
	return nums, boards
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
