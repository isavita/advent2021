package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day17/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ", ")
	xRange := strings.Split(parts[0][15:], "..")
	yRange := strings.Split(parts[1][2:], "..")
	xMin, _ := strconv.Atoi(xRange[0])
	xMax, _ := strconv.Atoi(xRange[1])
	yMin, _ := strconv.Atoi(yRange[0])
	yMax, _ := strconv.Atoi(yRange[1])

	velocities := make(map[string]bool)
	for xVel := -1000; xVel <= 1000; xVel++ {
		for yVel := -1000; yVel <= 1000; yVel++ {
			xPos, yPos := 0, 0
			curXVel, curYVel := xVel, yVel
			inTargetArea := false
			for {
				xPos += curXVel
				yPos += curYVel

				if xPos >= xMin && xPos <= xMax && yPos >= yMin && yPos <= yMax {
					inTargetArea = true
					break
				}

				if isMovingAway(xPos, yPos, curXVel, curYVel, xMin, xMax, yMin, yMax) {
					break
				}

				if curXVel > 0 {
					curXVel--
				} else if curXVel < 0 {
					curXVel++
				}

				curYVel--
			}

			if inTargetArea {
				velocityKey := fmt.Sprintf("%d,%d", xVel, yVel)
				velocities[velocityKey] = true
			}
		}
	}

	fmt.Println(len(velocities))
}

func isMovingAway(xPos int, yPos int, xVel int, yVel int, xMin int, xMax int, yMin int, yMax int) bool {
	if xPos < xMin && xVel < 0 {
		return true
	}
	if xPos > xMax && xVel > 0 {
		return true
	}
	if yPos < yMin && yVel < 0 {
		return true
	}
	return false
}
