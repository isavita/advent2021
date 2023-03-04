package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type segment struct {
	start, end point
	index      int
}

func main() {
	file, err := os.Open("day5_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var segments []segment

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		startParts := strings.Split(parts[0], ",")
		endParts := strings.Split(parts[1], ",")
		start := point{toInt(startParts[0]), toInt(startParts[1])}
		end := point{toInt(endParts[0]), toInt(endParts[1])}
		segments = append(segments, segment{start, end, len(segments)})
	}

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, seg := range segments {
		markSegment(seg, grid)
	}

	count := countOverlap(grid)
	fmt.Println(count)
}

func toInt(s string) int {
	i := 0
	for _, c := range s {
		i = i*10 + int(c-'0')
	}
	return i
}

func markSegment(seg segment, grid [][]int) {
	start, end := seg.start, seg.end
	if start.x == end.x {
		if start.y > end.y {
			start, end = end, start
		}
		for y := start.y; y <= end.y; y++ {
			grid[start.x][y]++
		}
	} else if start.y == end.y {
		if start.x > end.x {
			start, end = end, start
		}
		for x := start.x; x <= end.x; x++ {
			grid[x][start.y]++
		}
	} else if abs(end.x-start.x) == abs(end.y-start.y) {
		if start.x > end.x {
			start, end = end, start
		}
		if start.y > end.y {
			for x, y := start.x, start.y; x <= end.x && y >= end.y; x, y = x+1, y-1 {
				grid[x][y]++
			}
		} else {
			for x, y := start.x, start.y; x <= end.x && y <= end.y; x, y = x+1, y+1 {
				grid[x][y]++
			}
		}
	}
}

func countOverlap(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val >= 2 {
				count++
			}
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
