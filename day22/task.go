package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var input = "day22/input.txt"

type Task struct {
	grid map[Point]string
	start Point
	directions []string
}

type Point struct {
	x,y int
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

	fmt.Println(task)

	part1 := task.mapPath()

	part2 := 0

	return part1, part2
}

func (t *Task) mapPath() int {
	position := t.start
	facing := "right"
	for _, direction := range t.directions {
		fmt.Println(position)
		switch direction {
		case "R":
			facing = rotateClockwise(facing)
		case "L":
			facing = rotateClockwise(facing)
		default:
			position = t.move(position, direction, facing)
		}
	}

	return 0
}

func (t *Task) move(position Point, direction, facing string) Point {
	
}

func rotateCounterClockwise(facing string) string {
	switch facing {
	case "right":
		return "up"
	case "up":
		return "left"
	case "left":
		return "down"
	case "down":
		return "right"
	}
	panic("invalid facing")
}

func rotateClockwise(facing string) string {
	switch facing {
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	case "up":
		return "right"
	}
	panic("invalid facing")
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	task := Task{
		grid: map[Point]string{},
	}

	y := 0
	gridSetting := true
	startSet := false

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
					}
				}
				y++
			} else {
				moves := ""
				for _, c := range in {
					_, err := strconv.Atoi(string(c)) 
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

	return task
}
