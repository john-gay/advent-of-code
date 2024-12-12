package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2024/day12/input.txt"

type coord struct {
	x int
	y int
}

var allVisited = map[coord]bool{}
var valueVisited = map[coord]bool{}

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
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			for j := 0; j < len(p); j++ {
				grid[coord{i, j}] = p[j]
			}
			i++
		}
	}

	for k, v := range grid {
		valueVisited = map[coord]bool{}
		perimeter := findPerimeter(grid, k, v)
		area := len(valueVisited)
		part1 += perimeter * area

		sides := findPerimeter(grid, k, v)
		part2 += sides * area
	}

	return part1, part2
}

func findPerimeter(grid map[coord]string, c coord, value string) int {
	if allVisited[c] {
		return 0
	}

	allVisited[c] = true
	valueVisited[c] = true

	directions := []coord{
		{c.x, c.y + 1},
		{c.x, c.y - 1},
		{c.x + 1, c.y},
		{c.x - 1, c.y},
	}

	perimeter := 0

	for _, dir := range directions {
		if grid[dir] != value {
			perimeter += 1
		} else {
			perimeter += findPerimeter(grid, dir, value)
		}
	}

	return perimeter
}
