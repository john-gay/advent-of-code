package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2023/day11/input.txt"

type task struct {
	galaxies  map[coord]string
	emptyRows map[int]bool
	emptyCols map[int]bool
	maxX int
	maxY int
}

type coord struct {
	x int
	y int
}

func main() {
	start := time.Now()

	task := parse()

	part1 := task.sumDistances(false)
	part2 := task.sumDistances(true)

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func distance(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func (t *task) distanceBetween(start, end coord, older bool) int {
	dX := distance(start.x, end.x)
	dY := distance(start.y, end.y)

	sX, eX := start.x, end.x
	if eX < sX {
		sX = end.x
		eX = start.x
	}
	sY, eY := start.y, end.y
	if eY < sY {
		sY = end.y
		eY = start.y
	}
	for i := sX; i < eX; i++ {
		if t.emptyCols[i] {
			if older {
				dX += 999999
			} else {
				dX++
			}
		}
	}
	for i := sY; i < eY; i++ {
		if t.emptyRows[i] {
			if older {
				dY += 999999
			} else {
				dY++
			}
		}
	}

	return dX + dY
}

func (t *task) sumDistances(older bool) int {
	sum := 0
	visited := map[coord]bool{}
	for start := range t.galaxies {
		for end := range t.galaxies {
			if start != end && !visited[end] {
				sum += t.distanceBetween(start, end, older)
			}
		}
		visited[start] = true
	}
	return sum
}

func parse() task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{
		galaxies: map[coord]string{},
		emptyRows: map[int]bool{},
		emptyCols: map[int]bool{},
	}

	y := 0
	colsUsed := map[int]bool{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			row := strings.Split(in, "")
			t.maxX = len(row) - 1
			rowEmpty := true
			for x, ch := range row {
				if ch != "." {
					t.galaxies[coord{x: x, y: y}] = ch
					colsUsed[x] = true
					rowEmpty = false
				}
			}
			if rowEmpty {
				t.emptyRows[y] = true
			}
			y += 1
		}
	}
	t.maxY = y
	for i := 0; i <= t.maxX; i++ {
		if colsUsed[i] {
			continue
		}
		t.emptyCols[i] = true
	}

	return t
}
