package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func depthFirstSearch(path []string, visited map[string]bool, graph map[string][]string, count *int) {
	current := path[len(path)-1]
	if current == "end" {
		*count++
		return
	}

	for _, neighbor := range graph[current] {
		if !visited[neighbor] || strings.ToUpper(neighbor) == neighbor {
			visited[neighbor] = true
			path = append(path, neighbor)
			depthFirstSearch(path, visited, graph, count)
			visited[neighbor] = false
			path = path[:len(path)-1]
		}
	}
}

func main() {
	file, err := os.Open("day12_part1/input.txt")
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
	visited := make(map[string]bool)
	visited[start] = true
	count := 0

	depthFirstSearch(path, visited, graph, &count)

	fmt.Println("Number of valid paths:", count)
}
