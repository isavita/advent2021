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
	file, _ := os.Open("day4_part1/input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var boards []board

	// Read the first line of the input as comma-separated numbers
	scanner.Scan()
	numStrs := strings.Split(scanner.Text(), ",")
	var numbers []int
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}

	// Read the rest of the input as boards
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

	// Mark the rows and columns of each board
	for _, num := range numbers {
		for i, board := range boards {
			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if board.grid[j][k] == num {
						board.markedGrid[j][k] = true
						boards[i] = board
					}
				}
			}
		}

		for _, board := range boards {
			if isWinner(board) {
				fmt.Println(calcScore(board, num))
				return
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
