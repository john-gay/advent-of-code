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

var input = "2023/day10/input.txt"

type task struct {
	grid  map[coord]string
	start coord
	maxX int
	maxY int
}

type position struct {
	coord coord
	direction string
}

type coord struct {
	x int
	y int
}

func main() {
	start := time.Now()

	task := parse()

	part1, longestStart := task.furthestPoint()
	part2 := task.tilesEnclosed(longestStart)

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (t *task) tilesEnclosed(start position) int {
	largerGrid := t.markPath(start, t.doubleGrid())
	t.grid = largerGrid
	t.maxX = t.maxX*2
	t.maxY = t.maxY*2

	t.floodFill(coord{t.maxX /2 -5, t.maxY /2})
	enclosed := 0

	// only count even rows
	for i := 0; i < t.maxX; i += 2 {
		for j := 0; j < t.maxY; j += 2 {
			if t.grid[coord{i,j}] == "I" {
				enclosed++
			}
		}
	}
	// t.printGrid()

	return enclosed
}

func (t *task) floodFill(c coord) {
	val, _ := t.grid[c]
	if (val == "X" || val == "I") {
		return
	}

	if c.x < 0 || c.y < 0 || c.x > t.maxX || c.y > t.maxY {
		return
	}

	t.grid[c] = "I"

	t.floodFill(coord{c.x+1, c.y})
	t.floodFill(coord{c.x-1, c.y})
	t.floodFill(coord{c.x, c.y+1})
	t.floodFill(coord{c.x, c.y-1})

	return
}

func (t *task) furthestPoint() (int, position) {
	startingCoords := []position{
		{coord{t.start.x+1, t.start.y}, "E"},
		{coord{t.start.x-1, t.start.y}, "W"},
		{coord{t.start.x, t.start.y-1}, "N"},
		{coord{t.start.x, t.start.y+1}, "S"},
	}
	max := 0
	var longestPosition position
	for _, start := range startingCoords {
		steps := t.checkPath(start)
		if steps > 0 && steps / 2 > max {
			max = steps / 2
			longestPosition = start
		}
	}
	return max, longestPosition
}

func middleCoord(coordA, coordB coord) coord {
	return coord{(coordA.x*2 + coordB.x*2) / 2, (coordA.y*2 + coordB.y*2) / 2}
}

func (t *task) markPath(start position, largerGrid map[coord]string) map[coord]string {
	currCoord := start.coord
	direction := start.direction
	largerGrid[middleCoord(t.start, currCoord)] = "X"
	for {
		nextStep, _, finished := t.nextStep(currCoord, &direction)
		largerGrid[middleCoord(currCoord, nextStep)] = "X"
		largerGrid[coord{currCoord.x*2,currCoord.y*2}] = "X"
		if finished {
			return largerGrid
		}
		currCoord = nextStep
	}
}

func (t *task) checkPath(start position) int {
	currCoord := start.coord
	direction := start.direction
	steps := 1
	for {
		nextStep, blocked, finished := t.nextStep(currCoord, &direction)
		if finished {
			return steps
		}
		if blocked {
			return -999
		}
		steps++
		currCoord = nextStep
	}
}

func (t *task) nextStep(currCoord coord, direction *string) (coord, bool, bool) {
	blocked := false
	var nextStep coord
	curr, ok := t.grid[currCoord]
	if !ok {
		curr = "."
	}
	switch curr {
	case "|":
		if *direction == "N" {
			nextStep = coord{currCoord.x, currCoord.y - 1}
		} else if *direction == "S" {
			nextStep = coord{currCoord.x, currCoord.y + 1}
		} else {
			blocked = true
		}
	case "-":
		if *direction == "E" {
			nextStep = coord{currCoord.x + 1, currCoord.y}
		} else if *direction == "W" {
			nextStep = coord{currCoord.x - 1, currCoord.y}
		} else {
			blocked = true
		}
	case "L":
		if *direction == "S" {
			nextStep = coord{currCoord.x + 1, currCoord.y}
			*direction = "E"
		} else if *direction == "W" {
			nextStep = coord{currCoord.x, currCoord.y - 1}
			*direction = "N"
		} else {
			blocked = true
		}
	case "J":
		if *direction == "S" {
			nextStep = coord{currCoord.x - 1, currCoord.y}
			*direction = "W"
		} else if *direction == "E" {
			nextStep = coord{currCoord.x, currCoord.y - 1}
			*direction = "N"
		} else {
			blocked = true
		}
	case "7":
		if *direction == "N" {
			nextStep = coord{currCoord.x - 1, currCoord.y}
			*direction = "W"
		} else if *direction == "E" {
			nextStep = coord{currCoord.x, currCoord.y + 1}
			*direction = "S"
		} else {
			blocked = true
		}
	case "F":
		if *direction == "N" {
			nextStep = coord{currCoord.x + 1, currCoord.y}
			*direction = "E"
		} else if *direction == "W" {
			nextStep = coord{currCoord.x, currCoord.y + 1}
			*direction = "S"
		} else {
			blocked = true
		}
	case ".":
		blocked = true
	case "S":
		return coord{}, false, true
	}
	return nextStep, blocked, false
}

func parse() task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{
		grid: map[coord]string{},
		maxX: 0,
		maxY: 0,
	}

	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			row := strings.Split(in, "")
			t.maxX = len(row)
			for x, ch := range row {
				if t.grid[coord{x: x, y: y}] == "" {
					if ch != "." {
						t.grid[coord{x: x, y: y}] = ch
					}
					if ch == "S" {
						t.start = coord{x: x, y: y}
					}
					_, err := strconv.Atoi(ch)
					if err != nil {
						continue
					}
				}
			}
			y += 1
		}
	}

	t.maxY = y

	return t
}

func (t *task) doubleGrid() map[coord]string {
	grid := map[coord]string{}
	for c, val  := range t.grid {
		grid[coord{c.x*2, c.y*2}] = val
	}
	return grid
}

func (t *task) printGrid() {
	screen := make([][]string, t.maxY)
	for i := range screen {
		screen[i] = make([]string, t.maxX)
	}
	for i := range screen {
		for j := range screen[i] {
			screen[i][j] = "."
		}
	}
	for coord, val  := range t.grid {
		screen[coord.y][coord.x] = val
	}
	for i := 0; i < len(screen); i++ {
		fmt.Println(screen[i])
	}
}