package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2023/day14/input.txt"

type task struct {
	grid map[coord]string
	maxX int
	maxY int
}

type coord struct {
	x int
	y int
}

func main() {
	start := time.Now()

	// 1999 - 1973 = 26
	// (1000000000 - 2014) / 26 = 0rem

	t := parse()
	part1 := t.totalLoad(false)

	t2 := parse()
	part2 := t2.totalLoad(true)

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2)) // manually worked out from repeating pattern

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (t *task) totalLoad(spin bool) int {
	sum := 0

	if spin {
		count := 0
		//iter := 1000000000
		iter := 2000

		for {
			if count == iter {
				break
			}
			previous := t.copyGrid()
			t.tiltNorth()
			//t.printGrid()
			t.tiltWest()

			//t.printGrid()
			t.tiltSouth()
			//t.printGrid()
			t.tiltEast()

			//t.printGrid()
			diff := t.compareGrid(previous)
			fmt.Println(count, diff, t.calcScore(0))
			if diff == 0 {
				fmt.Println(count)
				break
			}
			count++
		}
	} else {
		t.tiltNorth()
	}

	sum = t.calcScore(sum)

	return sum
}

func (t *task) calcScore(sum int) int {
	for key, value := range t.grid {
		if value == "O" {
			sum += t.maxY - key.y
		}
	}
	return sum
}

func (t *task) compareGrid(compareTo map[coord]string) int {
	diff := 0
	for x := 0; x < t.maxX; x++ {
		for y := 0; y < t.maxY; y++ {
			if t.grid[coord{x, y}] != compareTo[coord{x, y}] {
				diff++
			}
		}
	}
	return diff
}

func (t *task) copyGrid() map[coord]string {
	previous := map[coord]string{}
	for key, value := range t.grid {
		previous[key] = value
	}
	return previous
}

func (t *task) tiltNorth() {
	newGrid := map[coord]string{}
	for y := 0; y < t.maxY; y++ {
		for x := 0; x < t.maxX; x++ {
			key := coord{x, y}
			value, ok := t.grid[key]
			if ok {
				if value == "O" {
					newPos := key
					for y2 := key.y - 1; y2 >= 0; y2-- {
						_, ok := newGrid[coord{key.x, y2}]
						if ok {
							break
						}
						newPos = coord{key.x, y2}
					}
					newGrid[newPos] = value
				} else {
					newGrid[key] = value
				}
			}
		}
	}
	t.grid = newGrid
}

func (t *task) tiltSouth() {
	newGrid := map[coord]string{}
	for y := t.maxY - 1; y >= 0; y-- {
		for x := 0; x < t.maxX; x++ {
			key := coord{x, y}
			value, ok := t.grid[key]
			if ok {
				if value == "O" {
					newPos := key
					for y2 := key.y + 1; y2 < t.maxY; y2++ {
						_, ok := newGrid[coord{key.x, y2}]
						if ok {
							break
						}
						newPos = coord{key.x, y2}
					}
					newGrid[newPos] = value
				} else {
					newGrid[key] = value
				}
			}
		}
	}
	t.grid = newGrid
}

func (t *task) tiltEast() {
	newGrid := map[coord]string{}
	for x := t.maxX - 1; x >= 0; x-- {
		for y := 0; y < t.maxY; y++ {
			key := coord{x, y}
			value, ok := t.grid[key]
			if ok {
				if value == "O" {
					newPos := key
					for x2 := key.x + 1; x2 < t.maxX; x2++ {
						_, ok := newGrid[coord{x2, key.y}]
						if ok {
							break
						}
						newPos = coord{x2, key.y}
					}
					newGrid[newPos] = value
				} else {
					newGrid[key] = value
				}
			}
		}
	}
	t.grid = newGrid
}

func (t *task) tiltWest() {
	newGrid := map[coord]string{}
	for x := 0; x < t.maxX; x++ {
		for y := 0; y < t.maxY; y++ {
			key := coord{x, y}
			value, ok := t.grid[key]
			if ok {
				if value == "O" {
					newPos := key
					for x2 := key.x - 1; x2 >= 0; x2-- {
						_, ok := newGrid[coord{x2, key.y}]
						if ok {
							break
						}
						newPos = coord{x2, key.y}
					}
					newGrid[newPos] = value
				} else {
					newGrid[key] = value
				}
			}
		}
	}
	t.grid = newGrid
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
				if ch != "." {
					t.grid[coord{x, y}] = ch
				}
			}
			y += 1
		}
	}

	t.maxY = y

	return t
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
	for coord, val := range t.grid {
		screen[coord.y][coord.x] = val
	}
	for i := 0; i < len(screen); i++ {
		fmt.Println(screen[i])
	}
	fmt.Println("")
}
