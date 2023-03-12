package main

import (
	"bufio"
	"fmt"
	"os"
)

type octopus struct {
	energy  int
	flashed bool
}

func main() {
	// Open the input file
	file, err := os.Open("day11_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the input file
	scanner := bufio.NewScanner(file)
	grid := make([][]octopus, 10)
	for i := 0; i < 10; i++ {
		grid[i] = make([]octopus, 10)
		scanner.Scan()
		line := scanner.Text()
		for j, c := range line {
			grid[i][j].energy = int(c - '0')
		}
	}

	// Simulate steps
	for step := 1; ; step++ {
		// Count the total number of flashes
		totalFlashes := 0
		// Increase the energy of each octopus
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				grid[i][j].energy++
			}
		}

		// Check which octopuses will flash
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if grid[i][j].energy > 9 && !grid[i][j].flashed {
					grid[i][j].flashed = true
					totalFlashes++
					// Increase the energy of adjacent octopuses
					for ii := i - 1; ii <= i+1; ii++ {
						for jj := j - 1; jj <= j+1; jj++ {
							if ii >= 0 && ii < 10 && jj >= 0 && jj < 10 && !(ii == i && jj == j) {
								grid[ii][jj].energy++
							}
						}
					}
					// Start the outer loop over again
					i, j = 0, -1
				}
			}
		}

		// Reset the energy of flashed octopuses
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if grid[i][j].flashed {
					grid[i][j].energy = 0
					grid[i][j].flashed = false
				}
			}
		}

		if totalFlashes == 100 {
			fmt.Println(step)
			break
		}
	}
}
