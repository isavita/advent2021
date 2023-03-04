package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day6_part2/input.txt")
	if err != nil {
		panic(err)
	}

	fishStr := strings.Split(strings.TrimSpace(string(input)), ",")
	fishCounts := make([]int64, 9)
	for _, f := range fishStr {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		fishCounts[n]++
	}

	for day := 1; day <= 256; day++ {
		newFishCount := fishCounts[0]
		for i := 0; i < 8; i++ {
			fishCounts[i] = fishCounts[i+1]
			if i == 6 {
				fishCounts[i] += newFishCount
			}
		}
		fishCounts[8] = newFishCount

		if day%10 == 0 {
			total := int64(0)
			for _, count := range fishCounts {
				total += count
			}
			fmt.Printf("After %d days: %d\n", day, total)
		}
	}

	total := int64(0)
	for _, count := range fishCounts {
		total += count
	}
	fmt.Printf("After 256 days: %d\n", total)

}
