package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2023/day16/input.txt"

type task struct {
	grid map[coord]string
	maxX int
	maxY int
}

type position struct {
	direction string
	location  coord
}

type coord struct {
	x int
	y int
}

func main() {
	start := time.Now()

	t := parse()
	t.printGrid()
	part1 := t.engergiseGrid(position{"E", coord{-1, 0}})
	part2 := t.maximiseEngergiseGrid()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (t *task) maximiseEngergiseGrid() int {
	maximum := 0
	n := 0
	for x := 0; x < t.maxX; x++ {
		n = t.engergiseGrid(position{"S", coord{x, -1}})
		if n > maximum {
			maximum = n
		}
		n = t.engergiseGrid(position{"N", coord{x, t.maxY}})
		if n > maximum {
			maximum = n
		}
	}

	for y := 0; y < t.maxX; y++ {
		n = t.engergiseGrid(position{"E", coord{-1, y}})
		if n > maximum {
			maximum = n
		}
		n = t.engergiseGrid(position{"W", coord{t.maxX, y}})
		if n > maximum {
			maximum = n
		}
	}

	return maximum
}

func (t *task) engergiseGrid(start position) int {
	visited := map[position]bool{}

	t.travel(start, visited)

	//t.printVisited(visited)

	coordMap := map[coord]bool{}
	for key := range visited {
		coordMap[key.location] = true
	}

	return len(coordMap)
}

func (t *task) travel(pos position, visited map[position]bool) {
	for {
		pos = nextPosition(pos)
		if visited[pos] {
			return
		}
		if pos.location.x < 0 || pos.location.x >= t.maxX || pos.location.y < 0 || pos.location.y >= t.maxY {
			return
		}
		visited[pos] = true

		switch t.grid[pos.location] {
		case "/":
			switch pos.direction {
			case "E":
				pos.direction = "N"
			case "N":
				pos.direction = "E"
			case "W":
				pos.direction = "S"
			case "S":
				pos.direction = "W"
			}
		case "\\":
			switch pos.direction {
			case "E":
				pos.direction = "S"
			case "N":
				pos.direction = "W"
			case "W":
				pos.direction = "N"
			case "S":
				pos.direction = "E"
			}
		case "|":
			if pos.direction == "E" || pos.direction == "W" {
				t.travel(position{"N", pos.location}, visited)
				t.travel(position{"S", pos.location}, visited)
				return
			}
		case "-":
			if pos.direction == "N" || pos.direction == "S" {
				t.travel(position{"E", pos.location}, visited)
				t.travel(position{"W", pos.location}, visited)
				return
			}
		}
	}
}

func nextPosition(pos position) position {
	if pos.direction == "E" {
		return position{pos.direction, coord{pos.location.x + 1, pos.location.y}}
	}
	if pos.direction == "N" {
		return position{pos.direction, coord{pos.location.x, pos.location.y - 1}}
	}
	if pos.direction == "W" {
		return position{pos.direction, coord{pos.location.x - 1, pos.location.y}}
	}
	if pos.direction == "S" {
		return position{pos.direction, coord{pos.location.x, pos.location.y + 1}}
	}
	panic(fmt.Sprintf("invalid direction: %s", pos.direction))
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

func (t *task) printVisited(visited map[position]bool) {
	screen := make([][]string, t.maxY)
	for i := range screen {
		screen[i] = make([]string, t.maxX)
	}
	for i := range screen {
		for j := range screen[i] {
			screen[i][j] = "."
		}
	}
	for p := range visited {
		screen[p.location.y][p.location.x] = "X"
	}
	for i := 0; i < len(screen); i++ {
		fmt.Println(screen[i])
	}
	fmt.Println("")
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
