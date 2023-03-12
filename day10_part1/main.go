package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day10_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		// check if line is corrupted
		stack := []rune{}
		corrupted := false
		for _, c := range line {
			switch c {
			case '(', '[', '{', '<':
				stack = append(stack, c)
			case ')':
				if len(stack) > 0 && stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					score += 3
					corrupted = true
					goto end
				}
			case ']':
				if len(stack) > 0 && stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					score += 57
					corrupted = true
					goto end
				}
			case '}':
				if len(stack) > 0 && stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
				} else {
					score += 1197
					corrupted = true
					goto end
				}
			case '>':
				if len(stack) > 0 && stack[len(stack)-1] == '<' {
					stack = stack[:len(stack)-1]
				} else {
					score += 25137
					corrupted = true
					goto end
				}
			}
		}

	end:
		if corrupted {
			fmt.Printf("Corrupted line: %s\n", line)
		}
	}

	fmt.Printf("Total syntax error score: %d\n", score)
}
