# Prompt 1 - GPT-4
Context: You are a prominent competitive programmer. You have been invited to participate in a programming competition. You are given a task to solve.
Instructions: Write a golang program that solves the task and read the input from "day15_part1/test.txt".
Task:
"""
You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still fit, though; the main problem is that the walls of the cave are covered in chitons, and it would be best not to bump any of them.

The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern resembles a square; a quick scan of chiton density produces a map of risk level throughout the cave (your puzzle input). For example:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
You start in the top left position, your destination is the bottom right position, and you cannot move diagonally. The number at each position is its risk level; to determine the total risk of an entire path, add up the risk levels of each position you enter (that is, don't count the risk level of your starting position unless you enter it; leaving it adds no risk to your total).

Your goal is to find a path with the lowest total risk. In this example, a path with the lowest total risk is highlighted here:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
The total risk of this path is 40 (the starting position is never entered, so its risk is not counted).

What is the lowest total risk of any path from the top left to the bottom right?
"""
A:
"""
To solve this problem, you can use Dijkstra's algorithm to find the shortest path from the top-left to the bottom-right corner of the grid. Here's a Golang program that reads input from a file named "day15_part1/test.txt" and solves the task:
```go
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
	grid := readInput("day15_part1/test.txt")
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
```
This program reads the input from the file, creates a 2D grid, and applies Dijkstra's algorithm to find the lowest total risk path. Finally, it prints the minimum total risk for the path.
"""
