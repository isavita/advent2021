package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("day8_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " | ")
		inputValues := strings.Split(splitLine[0], " ")
		outputValues := strings.Split(splitLine[1], " ")

		digitMap := make(map[string]int)

		for _, value := range inputValues {
			switch len(value) {
			case 2:
				digitMap[sortString(value)] = 1
			case 3:
				digitMap[sortString(value)] = 7
			case 4:
				digitMap[sortString(value)] = 4
			case 5:
				digitMap[sortString(value)] = deduceDigitForValueWithLen5(inputValues, value)
			case 6:
				digitMap[sortString(value)] = deduceDigitForValueWithLen6(inputValues, value)
			case 7:
				digitMap[sortString(value)] = 8
			}
		}

		sum := 0
		for _, value := range outputValues {
			if digit, ok := digitMap[sortString(value)]; ok {
				sum = sum*10 + digit
			}
		}
		numbers = append(numbers, sum)
	}

	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println(sum)
}

func deduceDigitForValueWithLen5(inputValues []string, value string) int {
	indOne := findByLength(inputValues, 2)
	indFour := findByLength(inputValues, 4)
	if commonChars(value, inputValues[indOne]) == 2 {
		return 3
	} else if commonChars(value, inputValues[indFour]) == 3 {
		return 5
	} else {
		return 2
	}
}

func deduceDigitForValueWithLen6(inputValues []string, value string) int {
	indFour := findByLength(inputValues, 4)
	indOne := findByLength(inputValues, 2)
	if commonChars(value, inputValues[indFour]) == 4 {
		return 9
	} else if commonChars(value, inputValues[indOne]) == 2 {
		return 0
	} else {
		return 6
	}
}

func findByLength(values []string, length int) int {
	for index, value := range values {
		if len(value) == length {
			return index
		}
	}
	return -1
}

func commonChars(s1 string, s2 string) int {
	m := make(map[rune]bool)
	count := 0
	for _, c := range s1 {
		m[c] = true
	}
	for _, c := range s2 {
		if _, ok := m[c]; ok {
			count++
		}
	}
	return count
}

func sortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
