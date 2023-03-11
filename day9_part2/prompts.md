# Prompt 1 - Continue from day9_part1 session
Instructions: Write a go program that solves this problem and read the input from "day9_part2/test.txt". The part one of the problem is already solve you need to Only solve the part two of the problem that start after "--- Part Two ---".
Task:
"""
These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal vents release smoke into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678
Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?

Your puzzle answer was 554.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Next, you need to find the largest basins so you know what areas are most important to avoid.

A basin is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin, although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low point. The example above has four basins.

The top-left basin, size 3:

2199943210
3987894921
9856789892
8767896789
9899965678
The top-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
The middle basin, size 14:

2199943210
3987894921
9856789892
8767896789
9899965678
The bottom-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
Find the three largest basins and multiply their sizes together. In the above example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?
"""
A:
"""
Sure, here's a solution in Go for the part two of the problem:
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Open the input file
    file, err := os.Open("day9_part2/test.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Read the input into a 2D slice
    var heightMap [][]int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        row := make([]int, len(line))
        for i, char := range line {
            height, _ := strconv.Atoi(string(char))
            row[i] = height
        }
        heightMap = append(heightMap, row)
    }

    // Find the low points and their basins
    var basins [][][2]int
    visited := make(map[[2]int]bool)
    for i := 0; i < len(heightMap); i++ {
        for j := 0; j < len(heightMap[0]); j++ {
            location := [2]int{i, j}
            if isLowPoint(heightMap, i, j) && !visited[location] {
                basin := getBasin(heightMap, location, visited)
                basins = append(basins, basin)
            }
        }
    }

    // Find the three largest basins and multiply their sizes
    largestBasins := findLargestBasins(basins, 3)
    basinSizeProduct := 1
    for _, basin := range largestBasins {
        basinSizeProduct *= len(basin)
    }

    // Print the product of sizes of the three largest basins
    fmt.Println(basinSizeProduct)
}

// Check if the given location is a low point
func isLowPoint(heightMap [][]int, i int, j int) bool {
    height := heightMap[i][j]
    if i > 0 && heightMap[i-1][j] < height {
        return false
    }
    if i < len(heightMap)-1 && heightMap[i+1][j] < height {
        return false
    }
    if j > 0 && heightMap[i][j-1] < height {
        return false
    }
    if j < len(heightMap[0])-1 && heightMap[i][j+1] < height {
        return false
    }
    return true
}

// Get the basin of the given low point location
func getBasin(heightMap [][]int, location [2]int, visited map[[2]int]bool) [][2]int {
    queue := []([2]int){location}
    basin := [][2]int{}
    visited[location] = true
    for len(queue) > 0 {
        currLocation := queue[0]
        queue = queue[1:]
        basin = append(basin, currLocation)
        i, j := currLocation[0], currLocation[1]
        if i > 0 && !visited[[2]int{i-1, j}] && heightMap[i-1][j] >= heightMap[location[0]][location[1]] {
            visited[[2]int{i-1, j}] = true
            queue = append(queue, [2]int{i-1, j})
        }
        if i < len(heightMap)-1 && !visited[[2]int{i+1, j}] && heightMap[i+1][j] >= heightMap[location[0]][location[1]] {
            visited[[2]int{i+1, j}] = true
            queue =    append(queue, [2]int{i+1, j})
        }
        if j > 0 && !visited[[2]int{i, j-1}] && heightMap[i][j-1] >= heightMap[location[0]][location[1]] {
            visited[[2]int{i, j-1}] = true
            queue = append(queue, [2]int{i, j-1})
        }
        if j < len(heightMap[0])-1 && !visited[[2]int{i, j+1}] && heightMap[i][j+1] >= heightMap[location[0]][location[1]] {
            visited[[2]int{i, j+1}] = true
            queue = append(queue, [2]int{i, j+1})
        }
    }
    return basin
}

// Find the largest basins and return them in descending order of their sizes
func findLargestBasins(basins [][][2]int, numLargest int) [][][2]int {
    sizes := make([]int, len(basins))
    for i, basin := range basins {
        sizes[i] = len(basin)
    }
    largestSizes := make([]int, numLargest)
    largestBasins := make([][][2]int, numLargest)
    for i := 0; i < numLargest; i++ {
        largestSize := -1
        largestIndex := -1
        for j, size := range sizes {
            if size > largestSize {
                largestSize = size
                largestIndex = j
            }
        }
        if largestIndex >= 0 {
            largestSizes[i] = largestSize
            largestBasins[i] = basins[largestIndex]
            sizes[largestIndex] = -1
        }
    }
    return largestBasins
}
```
This program reads the input from the "day9_part2/test.txt" file and stores it in a 2D slice called heightMap. It then iterates over each low point location in the height map and gets its basin using the getBasin function. The basins are stored in a slice of slices called basins. The program then finds the three largest basins using the findLargestBasins function and multiplies their sizes to get the final answer.

The getBasin function uses a breadth-first search algorithm to find all locations that flow downward to the given low point location. The locations are stored in a slice called basin and returned by the function.

The findLargestBasins function takes a slice of basins and finds the largest basins by their size. It returns the largest basins in descending order of their sizes.

Note that this program assumes that the input file is in the correct format and contains only valid input. Error handling can be added to make the program more robust.
"""
