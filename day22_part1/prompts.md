# Prompt 1 - GPT-4
Context: You are a prominent competitive programmer. You have been invited to participate in a programming competition. You are given a task to solve.
Instructions: Write a golang program that solves the task and read the input from "day22_part1/test.txt".
Task:
"""
Operating at these extreme ocean depths has overloaded the submarine's reactor; it needs to be rebooted.

The reactor core is made up of a large 3-dimensional grid made up entirely of cubes, one cube per integer 3-dimensional coordinate (x,y,z). Each cube can be either on or off; at the start of the reboot process, they are all off. (Could it be an old model of a reactor you've seen before?)

To reboot the reactor, you just need to set all of the cubes to either on or off by following a list of reboot steps (your puzzle input). Each step specifies a cuboid (the set of all cubes that have coordinates which fall within ranges for x, y, and z) and whether to turn all of the cubes in that cuboid on or off.

For example, given these reboot steps:

on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..10
The first step (on x=10..12,y=10..12,z=10..12) turns on a 3x3x3 cuboid consisting of 27 cubes:

10,10,10
10,10,11
10,10,12
10,11,10
10,11,11
10,11,12
10,12,10
10,12,11
10,12,12
11,10,10
11,10,11
11,10,12
11,11,10
11,11,11
11,11,12
11,12,10
11,12,11
11,12,12
12,10,10
12,10,11
12,10,12
12,11,10
12,11,11
12,11,12
12,12,10
12,12,11
12,12,12
The second step (on x=11..13,y=11..13,z=11..13) turns on a 3x3x3 cuboid that overlaps with the first. As a result, only 19 additional cubes turn on; the rest are already on from the previous step:

11,11,13
11,12,13
11,13,11
11,13,12
11,13,13
12,11,13
12,12,13
12,13,11
12,13,12
12,13,13
13,11,11
13,11,12
13,11,13
13,12,11
13,12,12
13,12,13
13,13,11
13,13,12
13,13,13
The third step (off x=9..11,y=9..11,z=9..11) turns off a 3x3x3 cuboid that overlaps partially with some cubes that are on, ultimately turning off 8 cubes:

10,10,10
10,10,11
10,11,10
10,11,11
11,10,10
11,10,11
11,11,10
11,11,11
The final step (on x=10..10,y=10..10,z=10..10) turns on a single cube, 10,10,10. After this last step, 39 cubes are on.

The initialization procedure only uses cubes that have x, y, and z positions of at least -50 and at most 50. For now, ignore cubes outside this region.

Here is a larger example:

on x=-20..26,y=-36..17,z=-47..7
on x=-20..33,y=-21..23,z=-26..28
on x=-22..28,y=-29..23,z=-38..16
on x=-46..7,y=-6..46,z=-50..-1
on x=-49..1,y=-3..46,z=-24..28
on x=2..47,y=-22..22,z=-23..27
on x=-27..23,y=-28..26,z=-21..29
on x=-39..5,y=-6..47,z=-3..44
on x=-30..21,y=-8..43,z=-13..34
on x=-22..26,y=-27..20,z=-29..19
off x=-48..-32,y=26..41,z=-47..-37
on x=-12..35,y=6..50,z=-50..-2
off x=-48..-32,y=-32..-16,z=-15..-5
on x=-18..26,y=-33..15,z=-7..46
off x=-40..-22,y=-38..-28,z=23..41
on x=-16..35,y=-41..10,z=-47..6
off x=-32..-23,y=11..30,z=-14..3
on x=-49..-5,y=-3..45,z=-29..18
off x=18..30,y=-20..-8,z=-3..13
on x=-41..9,y=-7..43,z=-33..15
on x=-54112..-39298,y=-85059..-49293,z=-27449..7877
on x=967..23432,y=45373..81175,z=27513..53682
The last two steps are fully outside the initialization procedure area; all other steps are fully within it. After executing these steps in the initialization procedure region, 590784 cubes are on.

Execute the reboot steps. Afterward, considering only cubes in the region x=-50..50,y=-50..50,z=-50..50, how many cubes are on?
"""
A:
"""
Here is a Golang program that solves the given task. It reads the input file "day22_part1/test.txt", parses the reboot steps, and calculates the number of cubes that are on after executing the steps.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RebootStep struct {
	action    string
	xStart    int
	xEnd      int
	yStart    int
	yEnd      int
	zStart    int
	zEnd      int
}

func main() {
	file, err := os.Open("day22_part1/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rebootSteps := []RebootStep{}

	for scanner.Scan() {
		step := parseRebootStep(scanner.Text())
		rebootSteps = append(rebootSteps, step)
	}

	minCoord, maxCoord := -50, 50
	cubeGrid := createCubeGrid(minCoord, maxCoord)
	executeRebootSteps(cubeGrid, rebootSteps)
	onCubes := countOnCubes(cubeGrid)

	fmt.Println("Number of cubes that are on:", onCubes)
}

func parseRebootStep(line string) RebootStep {
	parts := strings.Split(line, " ")

	action := parts[0]
	xRange := strings.Split(parts[1][2:], "..")
	yRange := strings.Split(parts[2][2:], "..")
	zRange := strings.Split(parts[3][2:], "..")

	xStart, _ := strconv.Atoi(xRange[0])
	xEnd, _ := strconv.Atoi(xRange[1])
	yStart, _ := strconv.Atoi(yRange[0])
	yEnd, _ := strconv.Atoi(yRange[1])
	zStart, _ := strconv.Atoi(zRange[0])
	zEnd, _ := strconv.Atoi(zRange[1])

	return RebootStep{action, xStart, xEnd, yStart, yEnd, zStart, zEnd}
}

func createCubeGrid(minCoord, maxCoord int) [][][]bool {
	gridSize := maxCoord - minCoord + 1
	grid := make([][][]bool, gridSize)

	for i := range grid {
		grid[i] = make([][]bool, gridSize)
		for j := range grid[i] {
			grid[i][j] = make([]bool, gridSize)
		}
	}

	return grid
}

func executeRebootSteps(cubeGrid [][][]bool, rebootSteps []RebootStep) {
	for _, step := range rebootSteps {
		for x := step.xStart; x <= step.xEnd; x++ {
			for y := step.yStart; y <= step.yEnd; y++ {
				for z := step.zStart; z <= step.zEnd; z++ {
					if x >= -50 && x <= 50 && y >= -50 && y <= 50 && z >= -50 && z <= 50 {
						cubeGrid[x+50][y+50][z+50] = step.action == "on"
					}
				}
			}
		}
	}
}

func countOnCubes(cubeGrid [][][]bool) int {
	count := 0

	for i := range cubeGrid {
		for j := range cubeGrid[i] {
			for k := range cubeGrid[i][j] {
				if cubeGrid[i][j][k] {
					count++
				}
			}
		}
	}

	return count
}
```
This program reads the input file, creates a 3-dimensional grid of cubes, processes each reboot step, and counts the number of cubes that are on after executing all the steps.
"""
