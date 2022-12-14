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

var input = "day14/input.txt"

type task struct {
	grid   map[coord]string
	min    coord
	max    coord
	origin coord
}

type coord struct {
	x int
	y int
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
	t := readInput()

	part1 := t.simulateSand(false)

	t2 := readInput()
	t2.max.y += 2
	part2 := t2.simulateSand(true)

	return part1, part2
}

func (t *task) simulateSand(floor bool) int {
	simulating := true
	sandSum := 0
	for simulating {
		next := t.nextBlock(floor, t.origin)
		if next == nil {
			simulating = false
		} else {
			t.grid[*next] = "o"
			sandSum++
		}
	}
	return sandSum
}

func (t *task) nextBlock(floor bool, c coord) *coord {
	if floor {
		if t.grid[t.origin] == "o" {
			return nil
		} else if c.y+1 == t.max.y {
			return &c
		}
	} else {
		if c.y > t.max.y {
			return nil
		}
	}
	if t.grid[coord{c.x, c.y + 1}] == "" {
		return t.nextBlock(floor, coord{c.x, c.y + 1})
	} else if t.grid[coord{c.x - 1, c.y + 1}] == "" {
		return t.nextBlock(floor, coord{c.x - 1, c.y + 1})
	} else if t.grid[coord{c.x + 1, c.y + 1}] == "" {
		return t.nextBlock(floor, coord{c.x + 1, c.y + 1})
	} else {
		return &c
	}
}

func readInput() task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{
		grid:   map[coord]string{},
		origin: coord{x: 500, y: 0},
	}

	var minX, maxX, minY, maxY *int
	var lines [][][]int
	var nodes [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			parts := strings.Split(in, " -> ")
			for _, part := range parts {
				cell := strings.Split(part, ",")
				x, _ := strconv.Atoi(cell[0])
				y, _ := strconv.Atoi(cell[1])
				if minX == nil || *minX > x {
					minX = &x
				}
				if maxX == nil || *maxX < x {
					maxX = &x
				}
				if minY == nil || *minY > y {
					minY = &y
				}
				if maxY == nil || *maxY < y {
					maxY = &y
				}

				nodes = append(nodes, []int{x, y})
			}
			lines = append(lines, nodes)
			nodes = [][]int{}
		}
	}
	t.min = coord{x: *minX, y: *minY}
	t.max = coord{x: *maxX, y: *maxY}

	var previous []int
	var x1, x2, y1, y2 int
	for _, line := range lines {
		previous = nil
		for _, node := range line {
			if previous == nil {
				previous = node
				continue
			}
			if previous[0] == node[0] {
				x1 = node[0]
				y1 = node[1]
				y2 = previous[1]
				if previous[1] < node[1] {
					y1 = previous[1]
					y2 = node[1]
				}
				for i := y1; i <= y2; i++ {
					t.grid[coord{x: x1, y: i}] = "#"
				}
			} else {
				x1 = node[0]
				x2 = previous[0]
				y1 = node[1]
				if previous[0] < node[0] {
					x1 = previous[0]
					x2 = node[0]
				}
				for i := x1; i <= x2; i++ {
					t.grid[coord{x: i, y: y1}] = "#"
				}
			}
			previous = node
		}
	}

	return t
}

func (t *task) printGrid() {
	screen := make([][]string, t.max.y+1)
	for i := range screen {
		screen[i] = make([]string, t.max.x-t.min.x+1)
	}
	for i := range screen {
		for j := range screen[i] {
			screen[i][j] = "."
		}
	}
	for co, value := range t.grid {
		screen[co.y][co.x-t.min.x] = value
	}
	screen[t.origin.y][t.origin.x-t.min.x] = "+"
	for _, row := range screen {
		fmt.Println(row)
	}
}
