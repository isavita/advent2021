package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	file, err := os.Open("day3_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the diagnostic report from the file
	scanner := bufio.NewScanner(file)
	var report []string
	for scanner.Scan() {
		report = append(report, scanner.Text())
	}

	// Calculate the gamma rate
	gamma := ""
	for i := 0; i < len(report[0]); i++ {
		zeros, ones := 0, 0
		for _, number := range report {
			if number[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		if ones > zeros {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	// Calculate the epsilon rate
	epsilon := ""
	for i := 0; i < len(report[0]); i++ {
		zeros, ones := 0, 0
		for _, number := range report {
			if number[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		if ones < zeros {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}

	// Convert gamma and epsilon to decimal
	gammaDecimal := binaryToDecimal(gamma)
	epsilonDecimal := binaryToDecimal(epsilon)

	// Calculate the power consumption
	power := gammaDecimal * epsilonDecimal

	// Print the result
	fmt.Println(power)
}

// Converts a binary string to a decimal integer
func binaryToDecimal(binary string) int {
	decimal := 0
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			decimal += 1 << (len(binary) - i - 1)
		}
	}
	return decimal
}
