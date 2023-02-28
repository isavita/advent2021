package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1_part2/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var measurements []int
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing depth: %v\n", err)
			os.Exit(1)
		}
		measurements = append(measurements, depth)
	}

	var prevSum int = -1
	var count int
	for i := 0; i < len(measurements)-2; i++ {
		sum := measurements[i] + measurements[i+1] + measurements[i+2]
		if sum > prevSum && prevSum > -1 {
			count++
		}
		prevSum = sum
	}
	fmt.Println(count)
}
