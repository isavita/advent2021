package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day3_part2/input.txt")
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

	// Calculate the oxygen generator rating
	oxygenReport := make([]string, len(report))
	copy(oxygenReport, report)
	for i := 0; i < len(report[0]); i++ {
		zeros, ones := countBits(oxygenReport, i)
		if ones >= zeros {
			oxygenReport = filterNumbers(oxygenReport, i, '1')
		} else {
			oxygenReport = filterNumbers(oxygenReport, i, '0')
		}
		if len(oxygenReport) == 1 {
			break
		}
	}
	oxygenDecimal := binaryToDecimal(oxygenReport[0])

	// Calculate the CO2 scrubber rating
	co2Report := make([]string, len(report))
	copy(co2Report, report)
	for i := 0; i < len(report[0]); i++ {
		zeros, ones := countBits(co2Report, i)
		if ones < zeros {
			co2Report = filterNumbers(co2Report, i, '1')
		} else {
			co2Report = filterNumbers(co2Report, i, '0')
		}
		if len(co2Report) == 1 {
			break
		}
	}

	co2Decimal := binaryToDecimal(co2Report[0])

	// Calculate the life support rating
	life := oxygenDecimal * co2Decimal

	fmt.Println(life)
}

// A function that counts how many zeros and ones are in a given position of a slice of strings
func countBits(numbers []string, position int) (zeros int, ones int) {
	for _, number := range numbers {
		if number[position] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return
}

// A function that filters out the numbers that don't match a given bit in a given position
func filterNumbers(numbers []string, position int, bit byte) []string {
	filtered := make([]string, 0)
	for _, number := range numbers {
		if number[position] == bit {
			filtered = append(filtered, number)
		}
	}
	return filtered
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
