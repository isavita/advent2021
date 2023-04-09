package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct {
	x, y int
}

type Amphipod struct {
	position Point
	energy   int
}

func main() {
	content, err := ioutil.ReadFile("day23_part1/test.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	minEnergy := minEnergyToOrganizeAmphipods(grid)
	fmt.Println(minEnergy)
}

func minEnergyToOrganizeAmphipods(grid [][]rune) int {
	energyCost := map[rune]int{
		'A': 1,
		'B': 10,
		'C': 100,
		'D': 1000,
	}

	amphipodPositions := make(map[rune][]Amphipod)
	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' && cell != '#' {
				amphipodPositions[cell] = append(amphipodPositions[cell], Amphipod{position: Point{x, y}})
			}
		}
	}

	return bfsBacktracking(grid, amphipodPositions, energyCost)
}

func bfsBacktracking(grid [][]rune, amphipodPositions map[rune][]Amphipod, energyCost map[rune]int) int {
	directions := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var explore func(current rune, visited map[Point]bool, energy int) int
	explore = func(current rune, visited map[Point]bool, energy int) int {
		if current > 'D' {
			return energy
		}

		minEnergy := 1 << 30
		for i, amphipod := range amphipodPositions[current] {
			if _, ok := visited[amphipod.position]; ok {
				continue
			}

			visited[amphipod.position] = true
			for _, dir := range directions {
				newX, newY := amphipod.position.x+dir.x, amphipod.position.y+dir.y
				if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[0]) && grid[newY][newX] == '.' {
					energySpent := energyCost[current]
					amphipodPositions[current][i].energy += energySpent
					minEnergy = min(minEnergy, explore(current+1, visited, energy+energySpent))
					amphipodPositions[current][i].energy -= energySpent
				}
			}

			delete(visited, amphipod.position)
		}

		return minEnergy
	}

	return explore('A', make(map[Point]bool), 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
