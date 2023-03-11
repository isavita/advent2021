package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("day9_part1/input.txt")
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

	// Find the low points and their risk levels
	var riskLevelSum int
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[0]); j++ {
			if isLowPoint(heightMap, i, j) {
				riskLevel := heightMap[i][j] + 1
				riskLevelSum += riskLevel
			}
		}
	}

	// Print the sum of risk levels of all low points
	fmt.Println(riskLevelSum)
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
