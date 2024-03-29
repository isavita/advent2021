package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var vals []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	prevSum := 0
	count := 0
	for i := 2; i < len(vals); i++ {
		currSum := vals[i-2] + vals[i-1] + vals[i]
		if currSum > prevSum {
			count++
		}
		prevSum = currSum
	}

	fmt.Println(count)
}
