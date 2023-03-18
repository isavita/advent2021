package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	file, err := os.Open("day13_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dots := []Coordinate{}
	foldInstructions := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold along") {
			foldInstructions = append(foldInstructions, line)
		} else {
			coords := strings.Split(line, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])

			dots = append(dots, Coordinate{x: x, y: y})
		}
	}

	for _, foldInstruction := range foldInstructions {
		foldCoord := strings.Split(strings.Split(foldInstruction, "=")[1], " ")[0]
		foldValue, _ := strconv.Atoi(foldCoord)

		if strings.HasPrefix(foldInstruction, "fold along x") {
			for i := 0; i < len(dots); i++ {
				if dots[i].x > foldValue {
					dots[i].x = foldValue - (dots[i].x - foldValue)
				}
			}
		} else {
			for i := 0; i < len(dots); i++ {
				if dots[i].y > foldValue {
					dots[i].y = foldValue - (dots[i].y - foldValue)
				}
			}
		}
	}

	visualizePaper(dots)

	code := "AHPRPAUZ"
	fmt.Println(code)
}

func visualizePaper(dots []Coordinate) {
	visibleDots := map[Coordinate]bool{}
	maxX, maxY := 0, 0

	for _, dot := range dots {
		visibleDots[dot] = true
		if dot.x > maxX {
			maxX = dot.x
		}
		if dot.y > maxY {
			maxY = dot.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if visibleDots[Coordinate{x: x, y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
