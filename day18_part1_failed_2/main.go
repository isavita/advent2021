package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parseSnailFishNumbers(filename string) ([][]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	snailFishNumbers := make([][]interface{}, len(lines))

	for i, line := range lines {
		var snailFishNumber []interface{}
		err = json.Unmarshal([]byte(line), &snailFishNumber)
		if err != nil {
			return nil, err
		}
		snailFishNumbers[i] = snailFishNumber
	}

	return snailFishNumbers, nil
}

func addSnailFishNumbers(snailFishNumbers [][]interface{}) []interface{} {
	sum := snailFishNumbers[0]
	for i := 1; i < len(snailFishNumbers); i++ {
		sum = []interface{}{sum, snailFishNumbers[i]}
	}
	return sum
}

func explode(snailFishNumber []interface{}, parent []interface{}, index int, depth int) []interface{} {
	if depth >= 4 {
		allNumbers := true
		for _, element := range parent {
			if _, ok := element.(int); !ok {
				allNumbers = false
				break
			}
		}

		if allNumbers {
			left, right := snailFishNumber[0].(int), snailFishNumber[1].(int)

			if index > 0 {
				if leftNumber, ok := parent[index-1].(int); ok {
					parent[index-1] = leftNumber + left
				}
			}
			if index < len(parent)-1 {
				if rightNumber, ok := parent[index+1].(int); ok {
					parent[index+1] = rightNumber + right
				}
			}

			return []interface{}{0, 0}
		}
	}

	for i, element := range snailFishNumber {
		if nestedPair, ok := element.([]interface{}); ok {
			snailFishNumber[i] = explode(nestedPair, snailFishNumber, i, depth+1)
		}
	}

	return snailFishNumber
}

func split(snailFishNumber []interface{}) []interface{} {
	for i, element := range snailFishNumber {
		if number, ok := element.(int); ok && number >= 10 {
			snailFishNumber[i] = []interface{}{number / 2, number - number/2}
		} else if nestedPair, ok := element.([]interface{}); ok {
			snailFishNumber[i] = split(nestedPair)
		}
	}

	return snailFishNumber
}

func reduce(snailFishNumber []interface{}) []interface{} {
	prev := ""
	for {
		snailFishNumber = explode(snailFishNumber, nil, 0, 0)
		snailFishNumber = split(snailFishNumber)

		curr := fmt.Sprintf("%v", snailFishNumber)
		if curr == prev {
			break
		}
		prev = curr
	}

	return snailFishNumber
}

func magnitude(snailFishNumber []interface{}) int {
	left, right := snailFishNumber[0], snailFishNumber[1]

	var leftMag, rightMag int
	if leftNumber, ok := left.(int); ok {
		leftMag = leftNumber
	} else if leftPair, ok := left.([]interface{}); ok {
		leftMag = magnitude(leftPair)
	}

	if rightNumber, ok := right.(int); ok {
		rightMag = rightNumber
	} else if rightPair, ok := right.([]interface{}); ok {
		rightMag = magnitude(rightPair)
	}

	return 3*leftMag + 2*rightMag
}

func main() {
	snailFishNumbers, err := parseSnailFishNumbers("day18_part1/test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	sum := snailFishNumbers[0]
	for i := 1; i < len(snailFishNumbers); i++ {
		sum = addSnailFishNumbers([][]interface{}{sum, snailFishNumbers[i]})
		sum = reduce(sum)
	}

	mag := magnitude(sum)

	fmt.Println(mag)
}
