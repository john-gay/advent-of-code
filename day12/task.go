package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var input = "day12/input.txt"

type task struct {
	grid  map[coord]int
	start coord
	end   coord
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
	t := task{
		grid: map[coord]int{},
	}

	t.populateGrid()

	return t.breadthFirstSearch()
}

func (t *task) breadthFirstSearch() (int, int) {
	distance := map[coord]int{t.end: 0}
	queue := []coord{t.end}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range []coord{
			{current.x, current.y - 1},
			{current.x, current.y + 1},
			{current.x + 1, current.y},
			{current.x - 1, current.y}} {

			alreadyChecked := false
			if distance[next] != 0 {
				alreadyChecked = true
			}

			coordExists := false
			if t.grid[next] != 0 {
				coordExists = true
			}

			if !alreadyChecked && coordExists && t.grid[current] <= t.grid[next]+1 {
				distance[next] = distance[current] + 1
				queue = append(queue, next)
			}
		}
	}

	shortest := distance[t.start]
	for location, dist := range distance {
		if t.grid[location] == 1 && dist < shortest {
			shortest = dist
		}
	}

	return distance[t.start], shortest
}

func (t *task) populateGrid() {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			for j, p := range in {
				c := coord{x: j, y: index}
				if string(p) == "S" {
					t.start = c
					t.grid[t.start] = 1
				} else if string(p) == "E" {
					t.end = c
					t.grid[t.end] = 26
				} else {
					t.grid[c] = int(p) - 96
				}
			}
			index++
		}
	}
}
