package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var input = "2025/day11/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodes := map[string][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			parts := strings.Split(in, ": ")
			node := parts[0]
			neighbors := strings.Split(parts[1], " ")
			nodes[node] = neighbors
		}
	}

	part1, part2 := 0, 0

	start := nodes["you"]
	memo := make(map[string]int)
	for _, s := range start {
		part1 += followPaths(s, nodes, map[string]bool{}, map[string]bool{}, memo)
	}

	start = nodes["svr"]
	memo = make(map[string]int)
	mustVisit := map[string]bool{"dac": true, "fft": true}
	for _, s := range start {
		part2 += followPaths(s, nodes, mustVisit, map[string]bool{}, memo)
	}

	return part1, part2
}

func followPaths(current string, nodes map[string][]string, mustVisit map[string]bool, visited map[string]bool, memo map[string]int) int {
	if visited[current] {
		return 0
	}

	mustVisitKeys := make([]string, 0, len(mustVisit))
	for k := range mustVisit {
		mustVisitKeys = append(mustVisitKeys, k)
	}
	sort.Strings(mustVisitKeys)
	memoKey := current + "|" + strings.Join(mustVisitKeys, ",")

	if val, ok := memo[memoKey]; ok {
		return val
	}

	newVisited := make(map[string]bool)
	for k, v := range visited {
		newVisited[k] = v
	}
	newVisited[current] = true

	newMustVisit := make(map[string]bool)
	for k, v := range mustVisit {
		newMustVisit[k] = v
	}
	delete(newMustVisit, current)

	if current == "out" {
		if len(newMustVisit) == 0 {
			memo[memoKey] = 1
			return 1
		}
		memo[memoKey] = 0
		return 0
	}

	count := 0
	for _, s := range nodes[current] {
		count += followPaths(s, nodes, newMustVisit, newVisited, memo)
	}

	memo[memoKey] = count
	return count
}
