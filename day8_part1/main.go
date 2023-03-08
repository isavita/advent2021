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

	counts := map[int]int{
		1: 0,
		4: 0,
		7: 0,
		8: 0,
	}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " | ")
		outputValue := strings.Split(splitLine[1], " ")

		for _, value := range outputValue {
			if len(value) == 2 {
				counts[1]++
			}
			if len(value) == 3 {
				counts[7]++
			}
			if len(value) == 4 {
				counts[4]++
			}
			if len(value) == 7 {
				counts[8]++
			}
		}
	}

	fmt.Println(counts[1] + counts[4] + counts[7] + counts[8])
}
