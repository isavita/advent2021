package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row, col int
}

func main() {
	grid := readInput("day15_part1/input.txt")
	minRisk := findMinRiskPath(grid)
	fmt.Println("Minimum total risk:", minRisk)
}

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for i, ch := range strings.TrimSpace(line) {
			row[i], _ = strconv.Atoi(string(ch))
		}

		grid = append(grid, row)
	}

	return grid
}

func findMinRiskPath(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	dist := make([][]int, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = 1<<31 - 1 // Initialize to a large number
		}
	}

	start := Position{0, 0}
	dist[start.row][start.col] = 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for {
		minDist := 1<<31 - 1
		cur := Position{-1, -1}

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if !visited[i][j] && dist[i][j] < minDist {
					minDist = dist[i][j]
					cur = Position{i, j}
				}
			}
		}

		if cur.row == -1 && cur.col == -1 {
			break
		}

		visited[cur.row][cur.col] = true

		for _, dir := range dirs {
			nextRow := cur.row + dir[0]
			nextCol := cur.col + dir[1]

			if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && !visited[nextRow][nextCol] {
				newDist := dist[cur.row][cur.col] + grid[nextRow][nextCol]

				if newDist < dist[nextRow][nextCol] {
					dist[nextRow][nextCol] = newDist
				}
			}
		}
	}

	return dist[rows-1][cols-1]
}
