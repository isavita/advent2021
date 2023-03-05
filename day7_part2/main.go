package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day7_part2/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := strings.TrimSpace(string(data))
	numsStr := strings.Split(text, ",")

	var positions []int
	for _, numStr := range numsStr {
		position, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		positions = append(positions, position)
	}
	sort.Ints(positions)

	minFuel := -1
	var minPos int

	for i := 0; i < len(positions); i++ {
		pos := positions[i]
		fuel := 0
		for j := 0; j < len(positions); j++ {
			step := positions[j] - pos
			if step < 0 {
				step = -step
			}
			fuel += (step * (step + 1)) / 2
		}
		if minFuel == -1 || fuel < minFuel {
			minFuel = fuel
			minPos = pos
		}
	}

	fmt.Printf("The crabs should align at position %d with a total fuel cost of %d.\n", minPos, minFuel)
}
