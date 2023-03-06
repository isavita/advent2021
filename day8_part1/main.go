package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day8_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := map[string]int{
		"abcefg":  0,
		"bc":      1,
		"abged":   2,
		"abcdg":   3,
		"bcfg":    4,
		"acdfg":   5,
		"acdefg":  6,
		"abc":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	uniquePatterns := make(map[string]bool)
	for pattern := range patterns {
		uniquePatterns[pattern] = true
	}

	counts := map[int]int{
		1: 0,
		4: 0,
		7: 0,
		8: 0,
	}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " | ")

		patternsSeen := strings.Split(splitLine[0], " ")
		outputValue := splitLine[1]

		uniquePatternsSeen := make(map[string]bool)
		for _, pattern := range patternsSeen {
			uniquePatternsSeen[pattern] = true
		}

		digitMapping := make(map[string]int)
		for pattern := range uniquePatternsSeen {
			for p := range uniquePatterns {
				if pattern == p {
					digitMapping[pattern] = patterns[p]
					break
				}
			}
		}

		var decodedOutputValue string
		for _, pattern := range patternsSeen {
			decodedOutputValue += fmt.Sprintf("%d", digitMapping[pattern])
		}

		for _, digit := range []int{1, 4, 7, 8} {
			counts[digit] += strings.Count(decodedOutputValue, fmt.Sprintf("%d", digit))
		}
	}

	var totalCount int
	for _, count := range counts {
		totalCount += count
	}

	fmt.Println(totalCount)
}
