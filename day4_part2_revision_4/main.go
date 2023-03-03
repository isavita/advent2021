package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Board struct {
	Numbers [5][5]int
	Marked  [5][5]bool
}

func main() {
	input, err := ioutil.ReadFile("day4_part2_revision_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	calledNumbers := make([]int, 0)
	for _, numberStr := range strings.Split(lines[0], ",") {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}
		calledNumbers = append(calledNumbers, number)
	}

	var boards []Board
	if len(lines) > 0 {
		for i := 1; i < len(lines); i += 6 {
			boardLines := lines[i : i+5]
			board := Board{}
			for j, boardLine := range boardLines {
				for k, numberStr := range strings.Fields(boardLine) {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						log.Fatal(err)
					}
					board.Numbers[j][k] = number
				}
			}
			boards = append(boards, board)
		}
	} else {
		log.Fatal("Invalid input file format")
	}

	var winBoard Board
	var winNumber int
	for _, number := range calledNumbers {
		for _, board := range boards {
			board.mark(number)
			if board.wins() && number > winNumber {
				winBoard, winNumber = board, number
			}
		}
	}

	score := winBoard.unmarkedSum() * winNumber
	fmt.Println(score)
}

func (b *Board) mark(number int) {
	for i, row := range b.Numbers {
		for j, n := range row {
			if n == number {
				b.Marked[i][j] = true
			}
		}
	}
}

func (b Board) wins() bool {
	for i := 0; i < 5; i++ {
		rowWin := true
		colWin := true
		for j := 0; j < 5; j++ {
			if !b.Marked[i][j] {
				rowWin = false
			}
			if !b.Marked[j][i] {
				colWin = false
			}
		}
		if rowWin || colWin {
			return true
		}
	}
	return false
}

func (b Board) unmarkedSum() int {
	sum := 0
	for _, row := range b.Numbers {
		for _, n := range row {
			if !b.MarkedContains(n) {
				sum += n
			}
		}
	}
	return sum
}

func (b Board) MarkedContains(number int) bool {
	for i, row := range b.Numbers {
		for j, n := range row {
			if n == number && b.Marked[i][j] {
				return true
			}
		}
	}
	return false
}
