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
	file, err := os.Open("day13_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dots := []Coordinate{}
	var foldInstruction string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold along") {
			foldInstruction = line
			break
		}

		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		dots = append(dots, Coordinate{x: x, y: y})
	}

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

	visibleDots := map[Coordinate]bool{}
	for _, dot := range dots {
		visibleDots[dot] = true
	}

	fmt.Println(len(visibleDots))
}
