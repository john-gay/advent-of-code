package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var input = "2024/day6/input.txt"

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

	grid := [][]string{}
	start := []int{0, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			grid = append(grid, strings.Split(in, ""))

			for index, c := range grid[len(grid)-1] {
				if c == "^" {
					start = []int{len(grid) - 1, index}
				}
			}
		}
	}

	visited, _ := patrol(grid, start)

	part2 := 0
	for key := range visited {
		parts := strings.Split(key, ",")
		i, _ := strconv.Atoi(parts[0])
		j, _ := strconv.Atoi(parts[1])

		if i == start[0] && j == start[1] {
			continue
		}

		newGrid := make([][]string, len(grid))
		for i := range grid {
			newGrid[i] = make([]string, len(grid[i]))
			copy(newGrid[i], grid[i])
		}
		newGrid[i][j] = "#"
		_, stuck := patrol(newGrid, start)

		if stuck {
			part2++
		}
	}

	return len(visited), part2
}

func printGrid(newGrid [][]string) {
	for _, row := range newGrid {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}

func patrol(grid [][]string, start []int) (map[string]int, bool) {
	visited := map[string]int{}
	current := []int{start[0], start[1]}
	direction := "^"
	for {
		if numberOfVisits, ok := visited[fmt.Sprintf("%d,%d", current[0], current[1])]; ok {
			visited[fmt.Sprintf("%d,%d", current[0], current[1])]++

			if numberOfVisits >= 10 {
				return visited, true
			}
		} else {
			visited[fmt.Sprintf("%d,%d", current[0], current[1])] = 1
		}

		next := getNextLocation(direction, current)

		if next[0] < 0 || next[0] >= len(grid) || next[1] < 0 || next[1] >= len(grid[0]) {
			break
		}

		for {
			if grid[next[0]][next[1]] == "#" {
				direction = getNextDirection(direction)
				next = getNextLocation(direction, current)
			} else {
				break
			}
		}

		current = next
	}
	return visited, false
}

func getNextLocation(direction string, current []int) []int {
	if direction == "^" {
		return []int{current[0] - 1, current[1]}
	} else if direction == ">" {
		return []int{current[0], current[1] + 1}
	} else if direction == "v" {
		return []int{current[0] + 1, current[1]}
	} else {
		return []int{current[0], current[1] - 1}
	}
}

func getNextDirection(direction string) string {
	if direction == "^" {
		return ">"
	} else if direction == ">" {
		return "v"
	} else if direction == "v" {
		return "<"
	} else {
		return "^"
	}
}
