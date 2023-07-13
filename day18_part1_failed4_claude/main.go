package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pair struct {
	left, right interface{}
}

func (p pair) magnitude() int {
	return 3*p.left.(int) + 2*p.right.(int)
}

func parse(input string) interface{} {
	var stack []interface{}
	for _, r := range input {
		switch r {
		case '[':
			stack = append(stack, &pair{})
		case ']':
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack[len(stack)-1] = pair{left, right}
		case ',':
			// no-op
		default:
			num, _ := strconv.Atoi(string(r))
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func add(a, b interface{}) interface{} {
	sum := pair{a, b}
	for {
		if reduced := reduceOnce(sum); reduced {
			continue
		}
		return sum
	}
}

func reduceOnce(sn interface{}) (interface{}, bool) {
	switch s := sn.(type) {
	case pair:
		if p, ok := explode(s); ok {
			return p, true
		}
		if p, ok := split(s); ok {
			return p, true
		}
	}
	return sn, false
}

func explode(sn pair) (pair, bool) {
	p, ok := explodeHelper(sn, 0)
	return p, ok
}

func explodeHelper(sn pair, depth int) (pair, bool) {
	if depth >= 4 && isRegular(sn.left) && isRegular(sn.right) {
		return pair{0, 0}, true
	}

	if isPair(sn.left) {
		if p, ok := explodeHelper(sn.left.(pair), depth+1); ok {
			return pair{p, sn.right}, true
		}
	}

	if isPair(sn.right) {
		if p, ok := explodeHelper(sn.right.(pair), depth+1); ok {
			return pair{sn.left, p}, true
		}
	}

	return sn, false
}

func split(sn pair) (pair, bool) {
	if isRegular(sn.left) && sn.left.(int) >= 10 {
		return pair{splitRegular(sn.left.(int)), sn.right}, true
	}

	if isPair(sn.left) {
		if p, ok := split(sn.left.(pair)); ok {
			return pair{p, sn.right}, true
		}
	}

	if isRegular(sn.right) && sn.right.(int) >= 10 {
		return pair{sn.left, splitRegular(sn.right.(int))}, true
	}

	if isPair(sn.right) {
		if p, ok := split(sn.right.(pair)); ok {
			return pair{sn.left, p}, true
		}
	}

	return sn, false
}

// Other helper functions

func part1(inputs []string) int {
	// same logic as before
}

func main() {
	inputs := readInput("day18_part1/tests.txt")
	result := part1(inputs)
	fmt.Println(result)
}

func readInput(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
