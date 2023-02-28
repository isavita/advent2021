package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	content, err := ioutil.ReadFile("day2_part2/input.txt")
	if err != nil {
		panic(err)
	}

	// Split input into lines
	lines := strings.Split(string(content), "\n")

	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizontalPosition += distance
			depth += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	fmt.Println(horizontalPosition * depth)
}
