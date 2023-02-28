package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prevDepth int = math.MaxInt
	var count int
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing depth: %v\n", err)
			os.Exit(1)
		}
		if depth > prevDepth {
			count++
		}
		prevDepth = depth
	}
	fmt.Println(count)
}
