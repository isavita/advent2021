package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) (string, map[string]string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Text()

	rules := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}
	return polymer, rules
}

func applyInsertion(polymer string, rules map[string]string) string {
	var newPolymer strings.Builder
	for i := 0; i < len(polymer); i++ {
		newPolymer.WriteByte(polymer[i])
		if i < len(polymer)-1 {
			rule := rules[string(polymer[i])+string(polymer[i+1])]
			if rule != "" {
				newPolymer.WriteString(rule)
			}
		}
	}
	return newPolymer.String()
}

func countElements(polymer string) map[byte]int {
	count := make(map[byte]int)
	for i := 0; i < len(polymer); i++ {
		count[polymer[i]]++
	}
	return count
}

func main() {
	polymer, rules := readFile("day14_part1/input.txt")
	for i := 0; i < 10; i++ {
		polymer = applyInsertion(polymer, rules)
	}
	elementCount := countElements(polymer)

	maxCount := 0
	minCount := len(polymer)
	for _, count := range elementCount {
		if count > maxCount {
			maxCount = count
		}
		if count < minCount {
			minCount = count
		}
	}

	fmt.Println(maxCount - minCount)
}
