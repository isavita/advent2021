package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var firstVisit = false

func depthFirstSearchV2(path []string, visited map[string]int, graph map[string][]string, count *int) {
	current := path[len(path)-1]
	if current == "end" {
		*count++
		return
	}

	for _, neighbor := range graph[current] {
		if neighbor == "start" {
			continue
		}
		if visited[neighbor] < 1 || (visited[neighbor] == 1 && !firstVisit) || strings.ToUpper(neighbor) == neighbor {
			if visited[neighbor] == 1 && strings.ToLower(neighbor) == neighbor {
				firstVisit = true
			}
			visited[neighbor]++
			path = append(path, neighbor)
			depthFirstSearchV2(path, visited, graph, count)
			visited[neighbor]--
			if visited[neighbor] == 1 && strings.ToLower(neighbor) == neighbor {
				firstVisit = false
			}
			path = path[:len(path)-1]
		}
	}
}

func main() {
	file, err := os.Open("day12_part2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, "-")
		graph[nodes[0]] = append(graph[nodes[0]], nodes[1])
		graph[nodes[1]] = append(graph[nodes[1]], nodes[0])
	}

	start := "start"
	path := []string{start}
	visited := make(map[string]int)
	visited[start] = 1
	count := 0

	depthFirstSearchV2(path, visited, graph, &count)
	fmt.Println("Number of valid paths:", count)
}
