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
	wideGrid := map[coord]string{}
	moves := []string{}
	i := 0
	start, wideStart := coord{0, 0}, coord{0, 0}

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

					grid[coord{j, i}] = p[j]

					if p[j] == "@" {
						start = coord{j, i}
						wideStart = coord{j*2, i}
						wideGrid[coord{j*2, i}] = p[j]
					} else if p[j] == "O" {
						wideGrid[coord{j*2, i}] = "["
						wideGrid[coord{j*2+1, i}] = "]"
					} else {
						wideGrid[coord{j*2, i}] = p[j]
						wideGrid[coord{j*2+1, i}] = p[j]
					}
				}
				i++
			} else {
				moves = append(moves, p...)
			}
		}
	}

	makeMoves(grid, moves, start)

	for k, v := range grid {
		if v == "O" {
			part1 += k.x + 100 * k.y
		}
	}

	makeMoves(wideGrid, moves, wideStart)

	for k, v := range wideGrid {
		if v == "[" {
			part2 += k.x + 100 * k.y
		}
	}

	return part1, part2
}

func makeMoves(grid map[coord]string, moves []string, pos coord) {
	for _, m := range moves {
		var move coord
		switch m {
		case "^":
			move = coord{0, -1}
		case "v":
			move = coord{0, 1}
		case ">":
			move = coord{1, 0}
		case "<":
			move = coord{-1, 0}
		}

		if movable(grid, pos, move, m) {
			nextPos := coord{pos.x + move.x, pos.y + move.y}
			if grid[nextPos] == "O" || grid[nextPos] == "[" || grid[nextPos] == "]" {
				moveBox(grid, nextPos, move, m)
			}
			grid[nextPos] = "@"
			delete(grid, pos)

			pos = nextPos
		}
	}
}

func moveBox(grid map[coord]string, pos coord, move coord, m string) {
	prevPos := coord{pos.x - move.x, pos.y - move.y}
	nextPos := coord{pos.x + move.x, pos.y + move.y}
	if grid[pos] == "#" {
		return
	}
	if grid[pos] == "O" {
		moveBox(grid, nextPos, move, m)
	}
	if grid[pos] == "[" {
		if m == "^" || m == "v" {
			moveBox(grid, nextPos, move, m)
			moveBox(grid, coord{nextPos.x + 1, nextPos.y}, move, m)
			delete(grid, coord{pos.x + 1, pos.y})
		} else {
			moveBox(grid, nextPos, move, m)
		}
	}
	if grid[pos] == "]" {
		if m == "^" || m == "v" {
			moveBox(grid, nextPos, move, m)
			moveBox(grid, coord{nextPos.x - 1, nextPos.y}, move, m)
			delete(grid, coord{pos.x -1, pos.y})
		} else {
			moveBox(grid, nextPos, move, m)
		}
	}

	grid[pos] = grid[prevPos]
}

func movable(grid map[coord]string, pos coord, move coord, m string) bool {
	nextPos := coord{pos.x + move.x, pos.y + move.y}
	if grid[nextPos] == "#" {
		return false
	}
	if grid[nextPos] == "O" {
		return movable(grid, nextPos, move, m)
	}
	if grid[nextPos] == "[" {
		if m == "^" || m == "v" {
			left := movable(grid, nextPos, move, m)
			right := movable(grid, coord{nextPos.x + 1, nextPos.y}, move, m)
			return left && right
		} else {
			return movable(grid, nextPos, move, m)
		}
	}
	if grid[nextPos] == "]" {
		if m == "^" || m == "v" {
			right := movable(grid, nextPos, move, m)
			left := movable(grid, coord{nextPos.x - 1, nextPos.y}, move, m)
			return left && right
		} else {
			return movable(grid, nextPos, move, m)
		}
	}

	return true
}

func printGrid(grid map[coord]string) {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for pos := range grid {
		minX = min(minX, pos.x)
		minY = min(minY, pos.y)
		maxX = max(maxX, pos.x)
		maxY = max(maxY, pos.y)
	}

	fmt.Println()

	for y := minY; y <= maxY; y++ {
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
