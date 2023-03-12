package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	scores = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	openToCloseBrck = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
)

type result struct {
	line  string
	score int
}

func main() {
	file, err := os.Open("day10_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var results []result

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
					corrupted = true
					goto end
				}
			case ']':
				if len(stack) > 0 && stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					corrupted = true
					goto end
				}
			case '}':
				if len(stack) > 0 && stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
				} else {
					corrupted = true
					goto end
				}
			case '>':
				if len(stack) > 0 && stack[len(stack)-1] == '<' {
					stack = stack[:len(stack)-1]
				} else {
					corrupted = true
					goto end
				}
			}
		}

	end:
		if !corrupted {
			// incomplete line, complete it
			stack := []rune{}
			for _, c := range line {
				switch c {
				case '(', '[', '{', '<':
					stack = append(stack, c)
				case ')':
					if len(stack) > 0 && stack[len(stack)-1] == '(' {
						stack = stack[:len(stack)-1]
					} else {
						stack = append(stack, c)
					}
				case ']':
					if len(stack) > 0 && stack[len(stack)-1] == '[' {
						stack = stack[:len(stack)-1]
					} else {
						stack = append(stack, c)
					}
				case '}':
					if len(stack) > 0 && stack[len(stack)-1] == '{' {
						stack = stack[:len(stack)-1]
					} else {
						stack = append(stack, c)
					}
				case '>':
					if len(stack) > 0 && stack[len(stack)-1] == '<' {
						stack = stack[:len(stack)-1]
					} else {
						stack = append(stack, c)
					}
				}
			}
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score = score*5 + scores[openToCloseBrck[stack[i]]]
			}
			results = append(results, result{line, score})
		}
	}

	// sort results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].score < results[j].score
	})

	// find middle score
	middle := results[len(results)/2].score

	fmt.Printf("Middle score: %d\n", middle)
}
