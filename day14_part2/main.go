package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day14_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Text()
	pairRules := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		rule := strings.Split(line, " -> ")
		pairRules[rule[0]] = rule[1]
	}

	pairCounts := make(map[string]map[string]int)
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		if pairCounts[string(pair[0])] == nil {
			pairCounts[string(pair[0])] = make(map[string]int)
		}
		pairCounts[string(pair[0])][string(pair[1])]++
	}

	for step := 0; step < 40; step++ {
		newPairCounts := make(map[string]map[string]int)
		for first, secondCounts := range pairCounts {
			for second, count := range secondCounts {
				inserted := pairRules[first+second]
				if newPairCounts[first] == nil {
					newPairCounts[first] = make(map[string]int)
				}
				newPairCounts[first][inserted] += count
				if newPairCounts[inserted] == nil {
					newPairCounts[inserted] = make(map[string]int)
				}
				newPairCounts[inserted][second] += count
			}
		}
		pairCounts = newPairCounts
	}

	elementCounts := make(map[string]int)
	for _, secondCounts := range pairCounts {
		for second, count := range secondCounts {
			elementCounts[second] += count
		}
	}

	// Add the count of the first element in the polymer
	elementCounts[string(polymer[0])]++

	mostCommon := ""
	leastCommon := ""
	for element, count := range elementCounts {
		if mostCommon == "" || elementCounts[mostCommon] < count {
			mostCommon = element
		}
		if leastCommon == "" || elementCounts[leastCommon] > count {
			leastCommon = element
		}
	}

	fmt.Println(elementCounts[mostCommon] - elementCounts[leastCommon])
}
