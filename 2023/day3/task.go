package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var input = "2023/day3/input.txt"

type Task struct {
	grid  map[Coord]string
	numbers []Number
	maxX int
	maxY int
}

type Number struct {
	coords []Coord
	value int
}

type Coord struct {
	x int
	y int
}

func main() {
	start := time.Now()

	task := parse()

	// task.printGrid()
	// fmt.Println(task.numbers)

	part1 := task.sumEngineParts()
	part2 := task.sumGearRatio()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func parse() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	task := Task{
		grid: map[Coord]string{},
		numbers: []Number{},
		maxX: 0,
		maxY: 0,
	}

	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			row := strings.Split(in, "")
			task.maxX = len(row)
			numString := ""
			for x, ch := range row {
				if task.grid[Coord{x: x, y: y}] == "" {
					if ch != "." {
					
						task.grid[Coord{x: x, y: y}] = ch
					}
					_, err := strconv.Atoi(ch)
					if err != nil {
						continue
					}
					numString = ch
					number := Number{coords: []Coord{{x: x, y: y}}}
					numString = task.checkForward(x, y, row, &number, numString)
					number.value, _ = strconv.Atoi(numString)
					task.numbers = append(task.numbers, number)
				}
			}
			y += 1
		}
	}

	task.maxY = y

	return task
}

func (t *Task) checkForward(x, y int, row []string, number *Number, numString string) string{
	if len(row) == x+1 {
		return numString
	}
	ch := row[x+1]
	_, err := strconv.Atoi(ch)
	if err != nil {
		return numString
	}
	number.coords = append(number.coords, Coord{x: x+1, y: y})
	t.grid[Coord{x: x+1, y: y}] = ch
	numString += ch
	return t.checkForward(x+1, y, row, number, numString)
}

func (t *Task) sumGearRatio() int {
	sum := 0
	for c, value := range t.grid {
		if value == "*" {
			adjacentNumbers := map[int]bool{}
			for i := c.y - 1; i <= c.y + 1; i++ {
				for j := c.x - 1; j <= c.x + 1; j++ {
					coord := Coord{x: j, y: i}
					for _, n := range t.numbers {
						if slices.Contains(n.coords, coord) {
							adjacentNumbers[n.value] = true
						}
					}
				}
			}
			if len(adjacentNumbers) == 2 {
				keys := []int{}
				for k := range adjacentNumbers {
					keys = append(keys, k)
				}
				sum += keys[0] * keys[1]
			}
		}
	}
	return sum
}


func (t *Task) sumEngineParts() int {
	sum := 0
	for _, number := range t.numbers {
		coords := t.getAdjacentCoords(number)
		symbolFound := false
		for _, checkCoord := range coords {
			if t.grid[checkCoord] != "" {
				symbolFound = true
				break
			}
		}
		if symbolFound {
			sum += number.value
		}
	}
	return sum
}

func (t *Task) getAdjacentCoords(n Number) []Coord {
	coords := []Coord{}

	y := n.coords[0].y
	minX := n.coords[0].x
	maxX := n.coords[len(n.coords)-1].x
	for i := y-1; i <= y+1; i++ {
		for j := minX - 1; j <= maxX + 1; j++ {
			coord := Coord{x: j, y: i}
			if i >= 0 && j >= 0 && i < t.maxY && j < t.maxX && !slices.Contains(n.coords, coord) {
				coords = append(coords, coord)
			}
		}
	}

	return coords
}


func (t *Task) printGrid() {
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