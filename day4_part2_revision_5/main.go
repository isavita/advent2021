package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	grid       [5][5]int
	markedGrid [5][5]bool
}

func main() {
	file, _ := os.Open("day4_part2_revision_5/input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var boards []board

	scanner.Scan()
	numStrs := strings.Split(scanner.Text(), ",")
	var numbers []int
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}

	for scanner.Scan() {
		var grid [5][5]int
		for i := 0; i < 5; i++ {
			scanner.Scan()
			row := scanner.Text()
			row = strings.Trim(row, " ")
			row = strings.ReplaceAll(row, "  ", " ")
			numsStrs := strings.Split(row, " ")
			for j, numStr := range numsStrs {
				if numStr != "" {
					num, _ := strconv.Atoi(numStr)
					grid[i][j] = num
				}
			}
		}
		boards = append(boards, board{grid: grid})
	}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(boards); j++ {
			board := boards[j]
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if board.grid[k][l] == numbers[i] {
						board.markedGrid[k][l] = true
						boards[j] = board
					}
				}
			}
			if isWinner(board) {
				if len(boards) > 1 {
					boards = append(boards[:j], boards[j+1:]...)
					j = -1
				} else {
					fmt.Println(calcScore(board, numbers[i]))
					return
				}
			}
		}
	}
}

func isWinner(board board) bool {
	for i := 0; i < 5; i++ {
		rowWin := true
		colWin := true
		for j := 0; j < 5; j++ {
			if !board.markedGrid[i][j] {
				rowWin = false
			}
			if !board.markedGrid[j][i] {
				colWin = false
			}
		}
		if rowWin || colWin {
			return true
		}
	}
	return false
}

func calcScore(board board, num int) int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.markedGrid[i][j] {
				sum += board.grid[i][j]
			}
		}
	}
	return sum * num
}
