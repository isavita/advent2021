package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	content, err := ioutil.ReadFile("day2_part1/input.txt")
	if err != nil {
		panic(err)
	}

	// Split input into lines
	lines := strings.Split(string(content), "\n")

	horizontalPosition := 0
	depth := 0

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
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}
	fmt.Println(horizontalPosition * depth)
}
