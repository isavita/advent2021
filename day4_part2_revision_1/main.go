package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Rows     [][]int
	Columns  [][]int
	Diagonal []int
}

func NewBoard(rows [][]int) *Board {
	n := len(rows)
	columns := make([][]int, n)
	for i := range columns {
		columns[i] = make([]int, n)
	}
	for i, row := range rows {
		for j, num := range row {
			columns[j][i] = num
		}
	}
	diagonal := make([]int, n)
	for i := range rows {
		diagonal[i] = rows[i][i]
	}
	return &Board{
		Rows:     rows,
		Columns:  columns,
		Diagonal: diagonal,
	}
}

func (b *Board) MarkNumber(num int) {
	for i, row := range b.Rows {
		for j, val := range row {
			if val == num {
				b.Rows[i][j] = -1
				b.Columns[j][i] = -1
			}
		}
	}
}

func (b *Board) isWinner() bool {
	for _, row := range b.Rows {
		if allMarked(row) {
			return true
		}
	}
	for _, column := range b.Columns {
		if allMarked(column) {
			return true
		}
	}
	return false
}

func allMarked(nums []int) bool {
	for _, num := range nums {
		if num != -1 {
			return false
		}
	}
	return true
}

func (b *Board) Score(num int) int {
	sum := 0
	for _, row := range b.Rows {
		for _, num := range row {
			if num != -1 {
				sum += num
			}
		}
	}
	return sum * num
}

func main() {
	file, err := os.Open("day4_part2_revision_1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var boards []*Board
	scanner.Scan()
	numsStr := scanner.Text()
	nums := strings.Split(numsStr, ",")
	for scanner.Scan() {
		rows := make([][]int, 5)
		for i := range rows {
			scanner.Scan()
			line := strings.TrimSpace(scanner.Text())
			line = strings.ReplaceAll(line, "  ", " ")
			nums := strings.Split(line, " ")
			row := make([]int, 5)
			for j, numStr := range nums {
				num, _ := strconv.Atoi(numStr)
				row[j] = num
			}
			rows[i] = row
		}
		board := NewBoard(rows)
		boards = append(boards, board)
	}

	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)

		for j := 0; j < len(boards); j += 1 {
			board := boards[j]
			board.MarkNumber(num)
			if board.isWinner() {
				if len(boards) == 1 {
					fmt.Println(board.Score(num))
					return
				} else {
					boards = append(boards[:j], boards[j+1:]...)
					j = -1
				}
			}
		}
	}
}
