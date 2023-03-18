# Prompt 1 - GPT-4
Context: You are a prominent competitive programmer. You have been invited to participate in a programming competition. You are given a task to solve.
Instructions: Write a golang program that solves the task and read the input from "day13_part1/test.txt".
Task:
"""
The incredible pressures at this depth are starting to put a strain on your submarine. The submarine has polymerization equipment that would produce suitable materials to reinforce the submarine, and the nearby volcanically-active caves should even have the necessary input elements in sufficient quantities.

The submarine manual contains instructions for finding the optimal polymer formula; specifically, it offers a polymer template and a list of pair insertion rules (your puzzle input). You just need to work out what polymer would result after repeating the pair insertion process a few times.

For example:

NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
The first line is the polymer template - this is the starting point of the process.

The following section defines the pair insertion rules. A rule like AB -> C means that when elements A and B are immediately adjacent, element C should be inserted between them. These insertions all happen simultaneously.

So, starting with the polymer template NNCB, the first step simultaneously considers all three pairs:

The first pair (NN) matches the rule NN -> C, so element C is inserted between the first N and the second N.
The second pair (NC) matches the rule NC -> B, so element B is inserted between the N and the C.
The third pair (CB) matches the rule CB -> H, so element H is inserted between the C and the B.
Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.

After the first step of this process, the polymer becomes NCNBCHB.

Here are the results of a few steps using the above rules:

Template:     NNCB
After step 1: NCNBCHB
After step 2: NBCCNBBBCBHCB
After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073. After step 10, B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times; taking the quantity of the most common element (B, 1749) and subtracting the quantity of the least common element (H, 161) produces 1749 - 161 = 1588.

Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result. What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?
"""
To solve this task, you can follow these steps:

Read the input from the file "day13_part1/test.txt".
Parse the input into the polymer template and a list of pair insertion rules.
Implement a function to perform one step of pair insertion.
Apply the function 10 times to the polymer template.
Count the occurrences of each element in the resulting polymer.
Find the most and least common elements and subtract their quantities.
Here's the Golang program that solves this task:
```go
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
	polymer, rules := readFile("day13_part1/test.txt")
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
```
This program reads the input from the file, applies 10 steps of pair insertion, and calculates the difference between the most and least common elements in the resulting polymer.
"""
