package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var input = "2024/day15/input.txt"

type coord struct {
	x, y int
}

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

	part1, part2 := 0, 0

	grid := map[coord]string{}
	moves := []string{}
	i := 0
	start := coord{0, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			if p[0] == "#" {
				for j := 0; j < len(p); j++ {
					if p[j] == "." {
						continue
					}
					if p[j] == "@" {
						start = coord{i, j}
					}
					grid[coord{i, j}] = p[j]
				}
				i++
			} else {
				moves = append(moves, p...)
			}
		}
	}

	printGrid(grid)

	makeMoves(grid, moves, start)

	printGrid(grid)

	return part1, part2
}

func makeMoves(grid map[coord]string, moves []string, pos coord) {
	for i, m := range moves {
		var move coord
		switch m {
		case "^":
			move = coord{-1, 0}
		case "v":
			move = coord{1, 0}
		case ">":
			move = coord{0, 1}
		case "<":
			move = coord{0, -1}
		}

		if movable(grid, pos, move) {
			moveChain(grid, pos, move)
			delete(grid, pos)
			pos.x += move.x
			pos.y += move.y
			grid[pos] = "@"
		}

		fmt.Println(i, m)
		printGrid(grid)
	}
}

func moveChain(grid map[coord]string, pos coord, move coord) bool {
	chain := []coord{}
	current := pos

	for {
		nextPos := coord{current.x + move.x, current.y + move.y}

		if grid[nextPos] == "#" {
			return false
		}

		chain = append(chain, current)

		if grid[nextPos] == "." {
			chain = append(chain, nextPos)
			break
		}

		if grid[nextPos] == "0" {
			current = nextPos
			continue
		}

		break
	}

	for i := len(chain) - 1; i > 0; i-- {
		current := chain[i]
		prev := chain[i-1]
		grid[current] = grid[prev]
	}

	delete(grid, chain[0])

	return true
}

func movable(grid map[coord]string, pos coord, move coord) bool {
	nextPos := coord{pos.x + move.x, pos.y + move.y}
	if grid[nextPos] == "#" {
		return false
	}
	if grid[nextPos] == "0" {
		return movable(grid, nextPos, move)
	}

	return true
}

func printGrid(grid map[coord]string) {
	// Find grid boundaries
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for pos := range grid {
		minX = min(minX, pos.x)
		minY = min(minY, pos.y)
		maxX = max(maxX, pos.x)
		maxY = max(maxY, pos.y)
	}

	// Print column numbers
	fmt.Print("   ")
	for x := minX; x <= maxX; x++ {
		fmt.Printf("%d", abs(x)%10)
	}
	fmt.Println()

	// Print grid rows
	for y := minY; y <= maxY; y++ {
		fmt.Printf("%2d ", abs(y)%10)
		for x := minX; x <= maxX; x++ {
			if val, exists := grid[coord{x, y}]; exists {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
