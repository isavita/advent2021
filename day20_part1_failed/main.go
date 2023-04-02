package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) (string, [][]byte) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var algorithm string
	var image [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if algorithm == "" {
			algorithm = line
		} else if strings.TrimSpace(line) != "" {
			image = append(image, []byte(line))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return algorithm, image
}

func applyAlgorithm(image [][]byte, algorithm string) [][]byte {
	newImage := make([][]byte, len(image)+2)
	for i := range newImage {
		newImage[i] = make([]byte, len(image[0])+2)
	}

	for y := -1; y <= len(image); y++ {
		for x := -1; x <= len(image[0]); x++ {
			index := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if x+dx >= 0 && x+dx < len(image[0]) && y+dy >= 0 && y+dy < len(image) && image[y+dy][x+dx] == '#' {
						index |= 1 << uint(8-(dy+1)*3-(dx+1))
					}
				}
			}

			newImage[y+1][x+1] = algorithm[index]
		}
	}

	return newImage
}

func countLitPixels(image [][]byte) int {
	count := 0
	for _, row := range image {
		for _, pixel := range row {
			if pixel == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	algorithm, image := readInput("day20_part1/input.txt")

	image = applyAlgorithm(image, algorithm)
	image = applyAlgorithm(image, algorithm)

	litPixels := countLitPixels(image)
	fmt.Println("Lit pixels after applying the image enhancement algorithm twice:", litPixels)
}
