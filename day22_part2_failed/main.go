package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RebootStep struct {
	action string
	xStart int
	xEnd   int
	yStart int
	yEnd   int
	zStart int
	zEnd   int
}

type OctreeNode struct {
	value       bool
	children    [8]*OctreeNode
	hasChildren bool
}

func (n *OctreeNode) Get(x, y, z, size int) bool {
	if !n.hasChildren {
		return n.value
	}

	half := size / 2
	index := 0

	if x >= half {
		index |= 1
		x -= half
	}

	if y >= half {
		index |= 2
		y -= half
	}

	if z >= half {
		index |= 4
		z -= half
	}

	return n.children[index].Get(x, y, z, half)
}

func (n *OctreeNode) Set(x, y, z, size int, value bool) {
	if !n.hasChildren && n.value == value {
		return
	}

	half := size / 2
	index := 0

	if x >= half {
		index |= 1
		x -= half
	}

	if y >= half {
		index |= 2
		y -= half
	}

	if z >= half {
		index |= 4
		z -= half
	}

	if n.hasChildren {
		n.children[index].Set(x, y, z, half, value)
		return
	}

	if !n.value {
		n.children[index].value = value
	}

	n.hasChildren = true
	n.children[index] = &OctreeNode{value: value}
	for i := range n.children {
		if i == index {
			continue
		}
		n.children[i] = &OctreeNode{value: n.value}
	}

	n.children[index].Set(x, y, z, half, value)
	n.value = false
}

func main() {
	file, err := os.Open("day22_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rebootSteps := []RebootStep{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		step := parseRebootStep(line)
		rebootSteps = append(rebootSteps, step)
	}

	rebootSteps = mergeOverlappingRebootSteps(rebootSteps)
	root := &OctreeNode{}
	gridSize := 100

	for _, step := range rebootSteps {
		for x := step.xStart; x <= step.xEnd; x++ {
			for y := step.yStart; y <= step.yEnd; y++ {
				for z := step.zStart; z <= step.zEnd; z++ {
					root.Set(x+gridSize/2, y+gridSize/2, z+gridSize/2, gridSize, step.action == "on")
				}
			}
		}
	}

	onCubes := countOnCubes(root, gridSize)
	fmt.Println("Number of cubes that are on:", onCubes)
}

func parseRebootStep(line string) RebootStep {
	parts := strings.Split(line, " ")

	action := parts[0]
	parts = strings.Split(parts[1], ",")
	xRange := strings.Split(parts[0][2:], "..")
	yRange := strings.Split(parts[1][2:], "..")
	zRange := strings.Split(parts[2][2:], "..")

	xStart, _ := strconv.Atoi(xRange[0])
	xEnd, _ := strconv.Atoi(xRange[1])
	yStart, _ := strconv.Atoi(yRange[0])
	yEnd, _ := strconv.Atoi(yRange[1])
	zStart, _ := strconv.Atoi(zRange[0])
	zEnd, _ := strconv.Atoi(zRange[1])

	return RebootStep{action, xStart, xEnd, yStart, yEnd, zStart, zEnd}
}

func mergeOverlappingRebootSteps(rebootSteps []RebootStep) []RebootStep {
	mergedSteps := []RebootStep{}

	for _, step := range rebootSteps {
		merged := false
		for i := range mergedSteps {
			if step.action == mergedSteps[i].action &&
				step.xStart <= mergedSteps[i].xEnd && step.xEnd >= mergedSteps[i].xStart &&
				step.yStart <= mergedSteps[i].yEnd && step.yEnd >= mergedSteps[i].yStart &&
				step.zStart <= mergedSteps[i].zEnd && step.zEnd >= mergedSteps[i].zStart {
				mergedSteps[i].xStart = min(step.xStart, mergedSteps[i].xStart)
				mergedSteps[i].xEnd = max(step.xEnd, mergedSteps[i].xEnd)
				mergedSteps[i].yStart = min(step.yStart, mergedSteps[i].yStart)
				mergedSteps[i].yEnd = max(step.yEnd, mergedSteps[i].yEnd)
				mergedSteps[i].zStart = min(step.zStart, mergedSteps[i].zStart)
				mergedSteps[i].zEnd = max(step.zEnd, mergedSteps[i].zEnd)
				merged = true
				break
			}
		}
		if !merged {
			mergedSteps = append(mergedSteps, step)
		}
	}

	return mergedSteps
}

func countOnCubes(root *OctreeNode, size int) int {
	count := 0

	for x := -size / 2; x <= size/2; x++ {
		for y := -size / 2; y <= size/2; y++ {
			for z := -size / 2; z <= size/2; z++ {
				if root.Get(x+size/2, y+size/2, z+size/2, size) {
					count++
				}
			}
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
