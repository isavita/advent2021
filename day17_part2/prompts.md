# Prompt 1 - Sage (from https://poe.com/)
Thank you! I have solved the program with the folloing golang code:
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day17_part1/input.txt")
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

	maxY := -1 << 30
	for xVel := -1000; xVel <= 1000; xVel++ {
		for yVel := -1000; yVel <= 1000; yVel++ {
			xPos, yPos := 0, 0
			curXVel, curYVel := xVel, yVel
			highestY := yPos
			for {
				xPos += curXVel
				yPos += curYVel

				if xPos >= xMin && xPos <= xMax && yPos >= yMin && yPos <= yMax {
					if highestY > maxY {
						maxY = highestY
					}
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
				if yPos > highestY {
					highestY = yPos
				}
			}
		}
	}

	fmt.Println("The highest y position the probe reaches:", maxY)
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
```
Instructions: The requirement changed and you are given a new task base on the previous to solve.
Task:
"""
--- Part Two ---
Maybe a fancy trick shot isn't the best idea; after all, you only have one probe, so you had better not miss.

To get the best idea of what your options are for launching the probe, you need to find every initial velocity that causes the probe to eventually be within the target area after any step.

In the above example, there are 112 different initial velocity values that meet these criteria:
How many distinct initial velocity values cause the probe to be within the target area after any step?
"""
A:
"""
To solve this task, we need to modify our existing code to keep track of all the initial velocities that cause the probe to eventually be within the target area after any step. We can do this by creating a map where the keys are the initial velocities (represented as strings), and the values are booleans indicating whether the probe eventually reaches the target area.

Here's the modified code in Go:
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day17_part1/input.txt")
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

	fmt.Println("Number of distinct initial velocity values:", len(velocities))
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
```
The main changes are:

We define a velocities map to keep track of the distinct initial velocities that cause the probe to eventually reach the target area.
Inside the nested loops over xVel and yVel, we add a new variable inTargetArea to keep track of whether the probe ever reaches the target area.
If the probe reaches the target area, we add the initial velocity to the velocities map.
After the loops, we output the size of the velocities map, which gives us the number of distinct initial velocities that cause the probe to eventually reach the target area.
This code should give us the answer to the task.
"""
