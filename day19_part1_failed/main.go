package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	id       int
	beacons  [][3]int
	absolute [3]int
}

func rotate(coordinates [3]int, axis int, clockwise bool) [3]int {
	var rotated [3]int
	copy(rotated[:], coordinates[:])

	if axis == 0 {
		if clockwise {
			rotated[1], rotated[2] = -coordinates[2], coordinates[1]
		} else {
			rotated[1], rotated[2] = coordinates[2], -coordinates[1]
		}
	} else if axis == 1 {
		if clockwise {
			rotated[0], rotated[2] = coordinates[2], -coordinates[0]
		} else {
			rotated[0], rotated[2] = -coordinates[2], coordinates[0]
		}
	} else {
		if clockwise {
			rotated[0], rotated[1] = -coordinates[1], coordinates[0]
		} else {
			rotated[0], rotated[1] = coordinates[1], -coordinates[0]
		}
	}

	return rotated
}

func rotateAll(coords [][3]int, axis int, clockwise bool) [][3]int {
	var rotated [][3]int
	for _, coord := range coords {
		rotated = append(rotated, rotate(coord, axis, clockwise))
	}
	return rotated
}

func findOverlap(s1, s2 *Scanner) ([12][3]int, bool) {
	matches := make(map[[3]int]int)

	for _, b1 := range s1.beacons {
		for _, b2 := range s2.beacons {
			if b1 == b2 {
				matches[b1]++
			}
		}
	}

	if len(matches) >= 12 {
		commonBeacons := [12][3]int{}
		i := 0
		for beacon := range matches {
			commonBeacons[i] = beacon
			i++
			if i == 12 {
				break
			}
		}
		return commonBeacons, true
	}

	return [12][3]int{}, false
}

func main() {
	file, err := os.Open("day19_part1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanners := []*Scanner{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "--- scanner ") {
			idStr := strings.TrimPrefix(line, "--- scanner ")
			id, _ := strconv.Atoi(idStr[:len(idStr)-4])
			sc := &Scanner{id: id}
			scanners = append(scanners, sc)
		} else {
			coords := strings.Split(line, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			z, _ := strconv.Atoi(coords[2])
			scanners[len(scanners)-1].beacons = append(scanners[len(scanners)-1].beacons, [3]int{x, y, z})
		}
	}

	found := 1
outerLoop:
	for found < len(scanners) {
		for i, s1 := range scanners[:found] {
			for _, s2 := range scanners[found:] {
				scannerMatched := false
				for axis := 0; axis < 3 && !scannerMatched; axis++ {
					for _, clockwise := range []bool{true, false} {
						s2RotatedBeacons := rotateAll(s2.beacons, axis, clockwise)
						if _, ok := findOverlap(s1, &Scanner{id: s2.id, beacons: s2RotatedBeacons}); ok {
							s1Common, _ := findOverlap(scanners[0], s1)
							s2Common, _ := findOverlap(scanners[0], &Scanner{id: s2.id, beacons: s2RotatedBeacons})

							for i := 0; i < 12; i++ {
								diff := [3]int{s1Common[i][0] - s2Common[i][0], s1Common[i][1] - s2Common[i][1], s1Common[i][2] - s2Common[i][2]}
								if i == 0 {
									s2.absolute = diff
								} else if s2.absolute != diff {
									break
								}
							}

							found++
							scannerMatched = true
							break
						}
					}
				}
				if scannerMatched {
					break
				}
			}
			if found > i+1 {
				break
			} else if i+1 == len(scanners[:found]) {
				break outerLoop
			}
		}
	}

	beacons := make(map[[3]int]struct{})
	for _, sc := range scanners {
		for _, beacon := range sc.beacons {
			abs := [3]int{sc.absolute[0] + beacon[0], sc.absolute[1] + beacon[1], sc.absolute[2] + beacon[2]}
			beacons[abs] = struct{}{}
		}
	}

	fmt.Println("Total beacons:", len(beacons))
}
