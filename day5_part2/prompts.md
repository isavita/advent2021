# Prompt 1
Instructions: Could you modify the previous code to take in account the new requirement?
Task: 
"""
Unfortunately, considering only horizontal and vertical lines doesn't give you the full picture; you need to also consider diagonal lines.

Because of the limits of the hydrothermal vent mapping system, the lines in your list will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees. In other words:

An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
Considering all lines from the above example would now produce the following diagram:

1.1....11.
.111...2..
..2.1.111.
...1.2.2..
.112313211
...1.2....
..1...1...
.1.....1..
1.......1.
222111....
You still need to determine the number of points where at least two lines overlap. In the above example, this is still anywhere in the diagram with a 2 or larger - now a total of 12 points.

Consider all of the lines. At how many points do at least two lines overlap?
"""
A: 
"""
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type segment struct {
	start, end point
	index int
}

func main() {
	file, err := os.Open("day5_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var segments []segment

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		startParts := strings.Split(parts[0], ",")
		endParts := strings.Split(parts[1], ",")
		start := point{toInt(startParts[0]), toInt(startParts[1])}
		end := point{toInt(endParts[0]), toInt(endParts[1])}
		segments = append(segments, segment{start, end, len(segments)})
	}

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, seg := range segments {
		markSegment(seg, grid)
	}

	count := countOverlap(grid)
	fmt.Println(count)
}

func toInt(s string) int {
	i := 0
	for _, c := range s {
		i = i*10 + int(c-'0')
	}
	return i
}

func markSegment(seg segment, grid [][]int) {
	start, end := seg.start, seg.end
	if start.x == end.x {
		if start.y > end.y {
			start, end = end, start
		}
		for y := start.y; y <= end.y; y++ {
			grid[start.x][y] |= 1 << seg.index
		}
	} else if start.y == end.y {
		if start.x > end.x {
			start, end = end, start
		}
		for x := start.x; x <= end.x; x++ {
			grid[x][start.y] |= 1 << seg.index
		}
	} else if abs(end.x-start.x) == abs(end.y-start.y) {
		if start.x > end.x {
			start, end = end, start
		}
		if start.y > end.y {
			for x, y := start.x, start.y; x <= end.x && y >= end.y; x, y = x+1, y-1 {
				grid[x][y] |= 1 << seg.index
			}
		} else {
			for x, y := start.x, start.y; x <= end.x && y <= end.y; x, y = x+1, y+1 {
				grid[x][y] |= 1 << seg.index
			}
		}
	}
}

func countOverlap(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if bitsSet(val) > 1 {
				count++
			}
		}
	}
	return count
}

func bitsSet(val int) int {
	count := 0
	for val != 0 {
		count++
		val &= val - 1
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```
The implementation now checks if a line segment is diagonal, and if so, marks all the points on the diagonal line using a loop. The countOverlap function remains the same.
"""
# Prompt 2
Could you remove the bit manipulation code and replace it with a more readable solution?
A:
"""
Certainly! Here's a modified implementation that does not use bit manipulation:
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type segment struct {
	start, end point
	index int
}

func main() {
	file, err := os.Open("day5_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var segments []segment

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		startParts := strings.Split(parts[0], ",")
		endParts := strings.Split(parts[1], ",")
		start := point{toInt(startParts[0]), toInt(startParts[1])}
		end := point{toInt(endParts[0]), toInt(endParts[1])}
		segments = append(segments, segment{start, end, len(segments)})
	}

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, seg := range segments {
		markSegment(seg, grid)
	}

	count := countOverlap(grid)
	fmt.Println(count)
}

func toInt(s string) int {
	i := 0
	for _, c := range s {
		i = i*10 + int(c-'0')
	}
	return i
}

func markSegment(seg segment, grid [][]int) {
	start, end := seg.start, seg.end
	if start.x == end.x {
		if start.y > end.y {
			start, end = end, start
		}
		for y := start.y; y <= end.y; y++ {
			grid[start.x][y]++
		}
	} else if start.y == end.y {
		if start.x > end.x {
			start, end = end, start
		}
		for x := start.x; x <= end.x; x++ {
			grid[x][start.y]++
		}
	} else if abs(end.x-start.x) == abs(end.y-start.y) {
		if start.x > end.x {
			start, end = end, start
		}
		if start.y > end.y {
			for x, y := start.x, start.y; x <= end.x && y >= end.y; x, y = x+1, y-1 {
				grid[x][y]++
			}
		} else {
			for x, y := start.x, start.y; x <= end.x && y <= end.y; x, y = x+1, y+1 {
				grid[x][y]++
			}
		}
	}
}

func countOverlap(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val >= 2 {
				count++
			}
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```
Instead of using bit manipulation, the modified implementation simply increments the count for each grid point that is covered by a line segment, and then counts the number of points with a count greater than or equal to 2. This should be easier to understand for those who are not familiar with bit manipulation. The markSegment function is modified to increment the count for each point on the diagonal line as well. The countOverlap function remains the same.
"""
