package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2024/day4/input.txt"

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			grid = append(grid, strings.Split(in, ""))
		}
	}

	part1 := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] != "X" {
				continue
			}

			if j-3 >= 0 {
				w := grid[i][j] + grid[i][j-1] + grid[i][j-2] + grid[i][j-3]
				if w == "XMAS" {
					part1++
				}
			}

			if j+3 < len(row) {
				e := grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i][j+3]
				if e == "XMAS" {
					part1++
				}
			}

			if i-3 >= 0 {
				n := grid[i][j] + grid[i-1][j] + grid[i-2][j] + grid[i-3][j]
				if n == "XMAS" {
					part1++
				}
			}

			if i+3 < len(grid) {
				n := grid[i][j] + grid[i+1][j] + grid[i+2][j] + grid[i+3][j]
				if n == "XMAS" {
					part1++
				}
			}

			if i-3 >= 0 && j-3 >= 0 {
				nw := grid[i][j] + grid[i-1][j-1] + grid[i-2][j-2] + grid[i-3][j-3]
				if nw == "XMAS" {
					part1++
				}
			}

			if i-3 >= 0 && j+3 < len(row) {
				ne := grid[i][j] + grid[i-1][j+1] + grid[i-2][j+2] + grid[i-3][j+3]
				if ne == "XMAS" {
					part1++
				}
			}

			if i+3 < len(grid) && j-3 >= 0 {
				sw := grid[i][j] + grid[i+1][j-1] + grid[i+2][j-2] + grid[i+3][j-3]
				if sw == "XMAS" {
					part1++
				}
			}

			if i+3 < len(grid) && j+3 < len(row) {
				se := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2] + grid[i+3][j+3]
				if se == "XMAS" {
					part1++
				}
			}
		}
	}

	part2 := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] != "A" {
				continue
			}

			if j-1 >= 0 && j+1 < len(row) && i-1 >= 0 && i+1 < len(grid) {
				x1 := grid[i-1][j-1] + grid[i][j] + grid[i+1][j+1]
				x2 := grid[i+1][j-1] + grid[i][j] + grid[i-1][j+1]
				if (x1 == "MAS" || x1 == "SAM") && (x2 == "MAS" || x2 == "SAM") {
					part2++
				}
			}
		}
	}

	return part1, part2
}
