package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type SnailfishNumber interface{}

type RegularNumber int

func add(left, right SnailfishNumber) SnailfishNumber {
	return []SnailfishNumber{left, right}
}

func explode(sn SnailfishNumber) SnailfishNumber {
	if exploded, ok := tryExplode(sn); ok {
		return exploded
	}

	switch v := sn.(type) {
	case []SnailfishNumber:
		for i, el := range v {
			v[i] = explode(el)
		}
	}

	return sn
}

func tryExplode(sn SnailfishNumber) (SnailfishNumber, bool) {
	nested := 0
	left, right := sn, sn

	for {
		switch v := left.(type) {
		case []SnailfishNumber:
			if len(v) == 2 {
				nested++
				left = v[0]
			} else {
				break
			}
		case RegularNumber:
			break
		}

		switch v := right.(type) {
		case []SnailfishNumber:
			if len(v) == 2 {
				nested++
				right = v[1]
			} else {
				break
			}
		case RegularNumber:
			break
		}

		if nested == 4 {
			if leftPair, lok := left.([]SnailfishNumber); lok {
				if leftNum, lnok := leftPair[1].(RegularNumber); lnok {
					left = leftNum + 1
				}
			}
			if rightPair, rok := right.([]SnailfishNumber); rok {
				if rightNum, rnok := rightPair[0].(RegularNumber); rnok {
					right = rightNum + 1
				}
			}

			return []SnailfishNumber{left, right}, true
		}

		break
	}

	return nil, false
}

func split(sn SnailfishNumber) SnailfishNumber {
	if splitted, ok := trySplit(sn); ok {
		return splitted
	}

	switch v := sn.(type) {
	case []SnailfishNumber:
		for i, el := range v {
			v[i] = split(el)
		}
	}

	return sn
}

func trySplit(sn SnailfishNumber) (SnailfishNumber, bool) {
	if num, ok := sn.(RegularNumber); ok && num >= 10 {
		return []SnailfishNumber{RegularNumber(num / 2), RegularNumber((num + 1) / 2)}, true
	}
	return nil, false
}

func reduce(sn SnailfishNumber) SnailfishNumber {
	for {
		exploded := explode(sn)
		if !isEqual(exploded, sn) {
			sn = exploded
			continue
		}

		splitted := split(sn)
		if !isEqual(splitted, sn) {
			sn = splitted
			continue
		}

		break
	}

	return sn
}

func magnitude(sn SnailfishNumber) int {
	switch v := sn.(type) {
	case RegularNumber:
		return int(v)
	case []SnailfishNumber:
		return 3*magnitude(v[0]) + 2*magnitude(v[1])
	}
	return 0
}

func main() {
	data, err := ioutil.ReadFile("day18_part1/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var snailfishNumbers []SnailfishNumber
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		var snailfishNumber SnailfishNumber
		err = json.Unmarshal([]byte(line), &snailfishNumber)
		if err != nil {
			fmt.Println("Error unmarshaling JSON", err)
			return
		}
		snailfishNumbers = append(snailfishNumbers, snailfishNumber)
	}

	reducedSnailfishNumbers := make([]SnailfishNumber, len(snailfishNumbers))
	for i, sn := range snailfishNumbers {
		reducedSnailfishNumbers[i] = reduce(sn)
	}

	totalMagnitude := 0
	for _, sn := range reducedSnailfishNumbers {
		totalMagnitude += magnitude(sn)
	}

	fmt.Println("Total Magnitude:", totalMagnitude)
}

func isEqual(a, b SnailfishNumber) bool {
	switch aVal := a.(type) {
	case RegularNumber:
		if bVal, ok := b.(RegularNumber); ok {
			return aVal == bVal
		}
	case []SnailfishNumber:
		if bVal, ok := b.([]SnailfishNumber); ok {
			if len(aVal) != len(bVal) {
				return false
			}
			for i := range aVal {
				if !isEqual(aVal[i], bVal[i]) {
					return false
				}
			}
			return true
		}
	}
	return false
}
