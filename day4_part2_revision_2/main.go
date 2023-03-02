package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numbers [][]int
}

func (b Board) isWinner() bool {
	for i := 0; i < 5; i++ {
		if b.numbers[i][0] != -1 && b.numbers[i][0]+b.numbers[i][1]+b.numbers[i][2]+b.numbers[i][3]+b.numbers[i][4] == -5 {
			return true
		}
		if b.numbers[0][i] != -1 && b.numbers[0][i]+b.numbers[1][i]+b.numbers[2][i]+b.numbers[3][i]+b.numbers[4][i] == -5 {
			return true
		}
	}
	return false
}

func (b Board) sumUnmarked() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] != -1 {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum
}

func lastWinner(boards []Board, nums []int) (int, int) {
	n := len(nums)
	lastBoard := -1
	lastNum := -1
	for i := 0; i < n; i++ {
		for j := 0; j < len(boards); j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if boards[j].numbers[k][l] == nums[i] {
						boards[j].numbers[k][l] = -1
					}
				}
			}
			if boards[j].isWinner() {
				lastBoard = j
				lastNum = nums[i]
			}
		}
		if lastBoard != -1 {
			break
		}
	}
	return lastBoard, lastNum
}

func main() {
	file, _ := os.Open("day4_part2_revision_2/test.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var boards []Board
	var currentBoard Board
	isNewBoard := true
	for scanner.Scan() {
		line := strings.TrimLeft(scanner.Text(), " ")
		if line == "" {
			if !isNewBoard {
				boards = append(boards, currentBoard)
			}
			currentBoard = Board{numbers: make([][]int, 5)}
			isNewBoard = true
			continue
		}
		if isNewBoard {
			nums := strings.Split(strings.TrimSpace(line), " ")
			for i := 0; i < 5; i++ {
				currentBoard.numbers[i] = make([]int, 5)
				currentBoard.numbers[i][0], _ = strconv.Atoi(nums[i])
			}
			isNewBoard = false
		} else {
			nums := strings.Split(strings.TrimSpace(line), " ")
			for i := 0; i < 5; i++ {
				currentBoard.numbers[i][len(boards)%5], _ = strconv.Atoi(nums[i])
			}
		}
	}
	if !isNewBoard {
		boards = append(boards, currentBoard)
	}

	if scanner.Scan() {
		numsStr := strings.Split(scanner.Text(), ",")
		nums := make([]int, len(numsStr))
		for i, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums[i] = num
		}

		// Part Two: find the last winning board
		lastBoard, lastNum := lastWinner(boards, nums)
		fmt.Println("Last winning board:", lastBoard)
		fmt.Println("Last number called:", lastNum)
		fmt.Println("Final score:", boards[lastBoard].sumUnmarked()*lastNum)
	}
}
