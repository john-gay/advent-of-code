package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var input = "day24/input.txt"

type Task struct {
	blizzard   []Blizzard
	start, end Point
	maxX, maxY int
}

type Blizzard struct {
	grid map[Point][]string
}

type Point struct {
	x, y int
}

type TimePoint struct {
	t, x, y int
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
	task := readInput()

	part1 := task.findShortestPath(TimePoint{0, task.start.x, task.start.y}, task.end)

	part2 := task.getSnacks()

	return part1, part2
}

func (t *Task) getSnacks() int {
	minutes := t.findShortestPath(TimePoint{0, t.start.x, t.start.y}, t.end)
	minutes = t.findShortestPath(TimePoint{minutes, t.end.x, t.end.y}, t.start)
	minutes = t.findShortestPath(TimePoint{minutes, t.start.x, t.start.y}, t.end)
	return minutes
}

func (t *Task) findShortestPath(start TimePoint, end Point) int {
	visited := map[TimePoint]bool{start: true}
	queue := []TimePoint{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range []TimePoint{
			{current.t + 1, current.x, current.y - 1},
			{current.t + 1, current.x, current.y + 1},
			{current.t + 1, current.x + 1, current.y},
			{current.t + 1, current.x - 1, current.y},
			{current.t + 1, current.x, current.y},
		} {

			if next.x < 0 || next.y < 0 {
				continue
			}

			if next.x == end.x && next.y == end.y {
				return next.t
			}
			if len(t.blizzard[next.t].grid[Point{next.x, next.y}]) == 0 && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return -1
}

func (t *Task) createBlizzard(i int) {
	t.blizzard[i] = Blizzard{grid: map[Point][]string{}}
	for x := 0; x < t.maxX; x++ {
		for y := 0; y < t.maxY; y++ {
			for _, value := range t.blizzard[i-1].grid[Point{x, y}] {
				nextX, nextY, empty := t.nextBlizzardPosition(x, y, value)
				if !empty {
					t.blizzard[i].grid[Point{nextX, nextY}] = append(t.blizzard[i].grid[Point{nextX, nextY}], value)
				}
			}
		}
	}
}

func (t *Task) nextBlizzardPosition(x, y int, value string) (int, int, bool) {
	switch value {
	case "#":
		return x, y, false
	case ">":
		if t.blizzard[0].grid[Point{x + 1, y}][0] == "#" {
			return 1, y, false
		}
		return x + 1, y, false
	case "<":
		if t.blizzard[0].grid[Point{x - 1, y}][0] == "#" {
			return t.maxX - 2, y, false
		}
		return x - 1, y, false
	case "^":
		if t.blizzard[0].grid[Point{x, y - 1}][0] == "#" {
			return x, t.maxY - 2, false
		}
		return x, y - 1, false
	case "v":
		if t.blizzard[0].grid[Point{x, y + 1}][0] == "#" {
			return x, 1, false
		}
		return x, y + 1, false
	}

	return x, y, true
}

func (t *Task) printBlizzard(blizzard Blizzard) {
	screen := make([][]string, t.maxY)
	for i := range screen {
		screen[i] = make([]string, t.maxX)
	}
	for i := range screen {
		for j := range screen[i] {
			screen[i][j] = " "
		}
	}
	for point, value := range blizzard.grid {
		if len(value) > 1 {
			screen[point.y][point.x] = strconv.Itoa(len(value))
		} else {
			screen[point.y][point.x] = value[0]
		}
	}
	fmt.Println("")
	for i := 0; i < len(screen); i++ {
		fmt.Println(screen[i])
	}
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	task := Task{
		start: Point{-1, -1},
	}

	blizzard := Blizzard{
		grid: map[Point][]string{},
	}
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			for i, c := range in {
				blizzard.grid[Point{i, y}] = []string{string(c)}
				if task.start.x == -1 && task.start.y == -1 && string(c) == "." {
					task.start = Point{i, y}
				}
				if string(c) == "." {
					task.end = Point{i, y}
				}
				task.maxX = i + 1
			}
			y++
			task.maxY = y
		}
	}

	maxT := task.maxX * task.maxY
	blizzards := make([]Blizzard, maxT)
	blizzards[0] = blizzard

	task.blizzard = blizzards

	for i := 1; i < maxT; i++ {
		task.createBlizzard(i)
	}

	return task
}
