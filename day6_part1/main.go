package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	input, err := ioutil.ReadFile("day6_part1/input.txt")
	if err != nil {
		panic(err)
	}

	// Parse input into a slice of integers
	fishStr := strings.Split(strings.TrimSpace(string(input)), ",")
	fish := make([]int, len(fishStr))
	for i, f := range fishStr {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		fish[i] = n
	}

	// Simulate growth for 80 days
	for day := 1; day <= 80; day++ {
		counts := make(map[int]int)
		for i, f := range fish {
			// Decrement fish age and handle wrapping around
			if f == 0 {
				fish[i] = 6
				counts[8]++
			} else {
				fish[i]--
			}

			// Add fish to counts if it has reached 0 age
			if fish[i] == 0 {
				counts[6]++
			}
		}

		// Create new fish for each 8-age fish
		for i := 0; i < counts[8]; i++ {
			fish = append(fish, 8)
		}

		// Print the total number of fish for every 10 days
		if day%10 == 0 {
			total := len(fish)
			fmt.Printf("After %d days: %d\n", day, total)
		}
	}

	// Print the final total number of fish
	total := len(fish)
	fmt.Printf("After 80 days: %d\n", total)
}
