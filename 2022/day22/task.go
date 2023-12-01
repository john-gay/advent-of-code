package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var input = "2022/day22/input.txt"

type Task struct {
	grid       map[Point]string
	start      Point
	directions []string
	max        Point
	facing     string
	cube       bool
	face       string
}

type Point struct {
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
	task := readInput()
	part1 := task.mapPath()

	task2 := readInput()
	task2.cube = true
	part2 := task2.mapPath()

	return part1, part2
}

func (t *Task) mapPath() int {
	position := t.start
	t.facing = "right"
	t.face = "A"
	for _, direction := range t.directions {
		switch direction {
		case "R":
			t.rotateClockwise()
		case "L":
			t.rotateCounterClockwise()
		default:
			position = t.move(position, direction)
		}
	}
	return t.password(position)
}

func (t *Task) password(position Point) int {
	var facingValue int
	switch t.facing {
	case "right":
		facingValue = 0
	case "left":
		facingValue = 2
	case "up":
		facingValue = 3
	case "down":
		facingValue = 1
	}
	return 1000*(position.y+1) + 4*(position.x+1) + facingValue
}

func (t *Task) move(position Point, direction string) Point {
	moves, _ := strconv.Atoi(direction)

	for i := 0; i < moves; i++ {
		var blocked bool
		var nextPosition Point
		switch t.facing {
		case "right":
			nextPosition, blocked = t.moveStep(position, Point{1, 0})
			if !blocked {
				position = nextPosition
			}
		case "up":
			nextPosition, blocked = t.moveStep(position, Point{0, -1})
			if !blocked {
				position = nextPosition
			}
		case "left":
			nextPosition, blocked = t.moveStep(position, Point{-1, 0})
			if !blocked {
				position = nextPosition
			}
		case "down":
			nextPosition, blocked = t.moveStep(position, Point{0, 1})
			if !blocked {
				position = nextPosition
			}
		}
	}

	return position
}

func (t *Task) moveStep(position Point, next Point) (Point, bool) {
	if t.cube == true {
		return t.moveStepCube(position, next)
	} else {
		return t.moveStepPlane(position, next)
	}
}

func (t *Task) moveStepCube(position Point, next Point) (Point, bool) {
	nextPosition := Point{position.x + next.x, position.y + next.y}

	if t.grid[nextPosition] == "#" {
		return position, true
	}

	switch t.face {
	case "A":
		if nextPosition.x < 50 {
			nextPosition = Point{0, 149 - nextPosition.y%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "E"
			t.facing = "right"
		} else if nextPosition.x > 99 {
			t.face = "B"
		} else if nextPosition.y < 0 {
			nextPosition = Point{0, 150 + nextPosition.x%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "F"
			t.facing = "right"
		} else if nextPosition.y > 49 {
			t.face = "C"
		}
	case "B":
		if nextPosition.x < 100 {
			t.face = "A"
		} else if nextPosition.x > 149 {
			nextPosition = Point{99, 149 - nextPosition.y%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "D"
			t.facing = "left"
		} else if nextPosition.y < 0 {
			nextPosition = Point{nextPosition.x % 50, 199}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "F"
			t.facing = "up"
		} else if nextPosition.y > 49 {
			nextPosition = Point{99, 50 + nextPosition.x%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "C"
			t.facing = "left"
		}
	case "C":
		if nextPosition.x < 50 {
			nextPosition = Point{nextPosition.y % 50, 100}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "E"
			t.facing = "down"
		} else if nextPosition.x > 99 {
			nextPosition = Point{100 + nextPosition.y%50, 49}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "B"
			t.facing = "up"
		} else if nextPosition.y < 50 {
			t.face = "A"
		} else if nextPosition.y > 99 {
			t.face = "D"
		}
	case "D":
		if nextPosition.x < 50 {
			t.face = "E"
		} else if nextPosition.x > 99 {
			nextPosition = Point{149, 49 - nextPosition.y%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "B"
			t.facing = "left"
		} else if nextPosition.y < 100 {
			t.face = "C"
		} else if nextPosition.y > 149 {
			nextPosition = Point{49, 150 + nextPosition.x%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "F"
			t.facing = "left"
		}
	case "E":
		if nextPosition.x < 0 {
			nextPosition = Point{50, 49 - nextPosition.y%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "A"
			t.facing = "right"
		} else if nextPosition.x > 49 {
			t.face = "D"
		} else if nextPosition.y < 100 {
			nextPosition = Point{50, 50 + nextPosition.x%50}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "C"
			t.facing = "right"
		} else if nextPosition.y > 149 {
			t.face = "F"
		}
	case "F":
		if nextPosition.x < 0 {
			nextPosition = Point{50 + nextPosition.y%50, 0}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "A"
			t.facing = "down"
		} else if nextPosition.x > 49 {
			nextPosition = Point{50 + nextPosition.y%50, 149}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "D"
			t.facing = "up"
		} else if nextPosition.y < 150 {
			t.face = "E"
		} else if nextPosition.y > 199 {
			nextPosition = Point{100 + nextPosition.x%50, 0}
			if t.grid[nextPosition] == "#" {
				return position, true
			}
			t.face = "B"
			t.facing = "down"
		}
	}

	return nextPosition, false
}

func (t *Task) moveStepPlane(position Point, next Point) (Point, bool) {
	var nextPosition Point
	nextPosition = Point{position.x + next.x, position.y + next.y}
	switch t.grid[nextPosition] {
	case "":
		return t.getNextPosition(nextPosition, next)
	case "#":
		return position, true
	case ".":
		return nextPosition, false
	}
	panic("invalid position")
}

func (t *Task) getNextPosition(position, next Point) (Point, bool) {
	var nextPosition Point
	if next.x != 0 {
		nextPosition = Point{getNext(position.x+next.x, t.max.x), position.y}
	} else {
		nextPosition = Point{position.x, getNext(position.y+next.y, t.max.y)}
	}

	switch t.grid[nextPosition] {
	case "":
		return t.getNextPosition(nextPosition, next)
	case "#":
		return position, true
	case ".":
		return nextPosition, false
	}
	panic("invalid cell")
}

func getNext(next, max int) int {
	if next < 0 {
		next = max
	}
	if next > max {
		next = 0
	}
	return next
}

func (t *Task) rotateCounterClockwise() {
	switch t.facing {
	case "right":
		t.facing = "up"
	case "up":
		t.facing = "left"
	case "left":
		t.facing = "down"
	case "down":
		t.facing = "right"
	}
}

func (t *Task) rotateClockwise() {
	switch t.facing {
	case "right":
		t.facing = "down"
	case "down":
		t.facing = "left"
	case "left":
		t.facing = "up"
	case "up":
		t.facing = "right"
	}
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	task := Task{
		grid: map[Point]string{},
		cube: false,
	}

	y := 0
	gridSetting := true
	startSet := false
	maxX := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			if gridSetting {
				for i, c := range in {
					if string(c) != " " {
						if !startSet {
							task.start = Point{i, y}
							startSet = true
						}
						task.grid[Point{i, y}] = string(c)
						if maxX < i {
							maxX = i
						}
					}
				}
				y++
			} else {
				moves := ""
				for _, c := range in {
					_, err = strconv.Atoi(string(c))
					if err != nil {
						task.directions = append(task.directions, moves, string(c))
						moves = ""
					} else {
						moves += string(c)
					}
				}
				if moves != "" {
					task.directions = append(task.directions, moves)
				}
			}
		} else {
			gridSetting = false
		}
	}

	task.max = Point{maxX, y - 1}

	return task
}
