package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("day9_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the input into a 2D slice
	var heightMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			height, _ := strconv.Atoi(string(char))
			row[i] = height
		}
		heightMap = append(heightMap, row)
	}

	// Find the low points and their basins
	var basins [][][2]int
	visited := make(map[[2]int]bool)
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[0]); j++ {
			location := [2]int{i, j}
			if isLowPoint(heightMap, i, j) && !visited[location] {
				tempBasin := getBasin(heightMap, location, visited)
				basin := [][2]int{}
				for _, location := range tempBasin {
					if heightMap[location[0]][location[1]] != 9 {
						basin = append(basin, location)
					}
				}
				basins = append(basins, basin)
			}
		}
	}

	// Find the three largest basins and multiply their sizes
	largestBasins := findLargestBasins(basins, 3)
	basinSizeProduct := 1
	for _, basin := range largestBasins {
		basinSizeProduct *= len(basin)
	}

	// Print the product of sizes of the three largest basins
	fmt.Println(basinSizeProduct)
}

// Check if the given location is a low point
func isLowPoint(heightMap [][]int, i int, j int) bool {
	height := heightMap[i][j]
	if i > 0 && heightMap[i-1][j] <= height {
		return false
	}
	if i < len(heightMap)-1 && heightMap[i+1][j] <= height {
		return false
	}
	if j > 0 && heightMap[i][j-1] <= height {
		return false
	}
	if j < len(heightMap[0])-1 && heightMap[i][j+1] <= height {
		return false
	}
	return true
}

// Get the basin of the given low point location
func getBasin(heightMap [][]int, location [2]int, visited map[[2]int]bool) [][2]int {
	queue := []([2]int){location}
	basin := [][2]int{}
	visited[location] = true
	for len(queue) > 0 {
		currLocation := queue[0]
		queue = queue[1:]
		basin = append(basin, currLocation)
		i, j := currLocation[0], currLocation[1]
		if i > 0 && !visited[[2]int{i - 1, j}] && heightMap[i-1][j] < 9 && heightMap[i-1][j] >= heightMap[location[0]][location[1]] {
			visited[[2]int{i - 1, j}] = true
			queue = append(queue, [2]int{i - 1, j})
		}
		if i < len(heightMap)-1 && !visited[[2]int{i + 1, j}] && heightMap[i+1][j] < 9 && heightMap[i+1][j] >= heightMap[location[0]][location[1]] {
			visited[[2]int{i + 1, j}] = true
			queue = append(queue, [2]int{i + 1, j})
		}
		if j > 0 && !visited[[2]int{i, j - 1}] && heightMap[i][j-1] < 9 && heightMap[i][j-1] >= heightMap[location[0]][location[1]] {
			visited[[2]int{i, j - 1}] = true
			queue = append(queue, [2]int{i, j - 1})
		}
		if j < len(heightMap[0])-1 && !visited[[2]int{i, j + 1}] && heightMap[i][j+1] < 9 && heightMap[i][j+1] >= heightMap[location[0]][location[1]] {
			visited[[2]int{i, j + 1}] = true
			queue = append(queue, [2]int{i, j + 1})
		}
	}
	return basin
}

// Find the largest basins and return them in descending order of their sizes
func findLargestBasins(basins [][][2]int, numLargest int) [][][2]int {
	sizes := make([]int, len(basins))
	for i, basin := range basins {
		sizes[i] = len(basin)
	}
	largestSizes := make([]int, numLargest)
	largestBasins := make([][][2]int, numLargest)
	for i := 0; i < numLargest; i++ {
		largestSize := -1
		largestIndex := -1
		for j, size := range sizes {
			if size > largestSize {
				largestSize = size
				largestIndex = j
			}
		}
		if largestIndex >= 0 {
			largestSizes[i] = largestSize
			largestBasins[i] = basins[largestIndex]
			sizes[largestIndex] = -1
		}
	}
	return largestBasins
}
