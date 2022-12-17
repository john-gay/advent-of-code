package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "day17/input.txt"

type Task struct {
	grid []Point
	jets []string
	shapes [][]Point
	width int
}

type Point struct {
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

	// fmt.Println(t)

	part1 := t.play(2022)

	// t = readInput()
	// part2 := t.play(1000000000000)

	return part1, 0
}

func (t *Task) play(rocks int) int {
	maxHeight := 0
	shapeIndex := 0
	jetIndex := 0

	// t.printGrid()

	for rock := 0; rock < rocks; rock++ {

		// fmt.Println(rock+1)
		// t.printGrid()

		var shape []Point
		shapeIndex, shape = t.getShape(shapeIndex)
		current := shapeLocation(2, maxHeight + 3, shape)
		
		// fmt.Println(fmt.Sprintf("start: %+v", current))

		falling := true
		for falling {
			var jet string
			jetIndex, jet = t.getJet(jetIndex)
			// fmt.Println(jet)
			if t.canPush(current, jet) {
				current = t.jetPush(current, jet)
			}
			// fmt.Println(fmt.Printf("after jet: %+v", current))
			if t.canFall(current) {
				current = fall(current)
			} else {
				falling = false
			}
		}

		// fmt.Println(fmt.Sprintf("after falling: %+v", current))

		t.setShape(current)
		height := shapeHeight(current)
		if maxHeight < height {
			maxHeight = height
		}

		// fmt.Println(fmt.Sprintf("rock: %d, height %d", rock, height))
	}
	// t.printGrid()

	return maxHeight
}

func shapeHeight(shape []Point) int {
	maxY := 0
	for _, point := range shape {
		if point.y > maxY {
			maxY = point.y
		}
	}
	return maxY + 1
}

func (t *Task) setShape(shape []Point) {
	for _, point := range shape {
		t.grid = append(t.grid, point)
	}
}

func fall(current []Point) []Point {
	next := []Point{}
	for _, point := range current {
		next = append(next, Point{point.x, point.y - 1})
	}
	return next
}

func (t *Task) canFall(current []Point) bool {
	// fmt.Println("In can fall")
	// fmt.Println(t.grid)
	// fmt.Println(current)
	for _, point := range current {
		if point.y -1 < 0 {
			// fmt.Println("Can't fall, y=0")

			return false
		}
		for _, set := range t.grid {
			below := point
			below.y -= 1
			if below == set {
				// fmt.Println("Can't fall, grid set")

				return false
			}
		}
	}
	// fmt.Println("Can fall")
	return true
}

func shapeLocation(x, y int, shape []Point) []Point {
	loc := []Point{}
	for _, point := range shape {
		loc = append(loc, Point{point.x + x, point.y + y})
	}
	return loc
}

func (t *Task) getShape(index int) (int, []Point) {
	if index < len(t.shapes) {
		return index + 1, t.shapes[index]
	}
	return 1, t.shapes[0]
}

func (t *Task) getJet(index int) (int, string) {
	if index < len(t.jets) {
		return index + 1, t.jets[index]
	}
	return 1, t.jets[0]
}

func (t *Task) jetPush(current []Point, jet string) []Point {
	if jet == ">" {
		next := []Point{}
		for _, point := range current {
			next = append(next, Point{point.x + 1, point.y})
		}
		return next
	} else {
		next := []Point{}
		for _, point := range current {
			next = append(next, Point{point.x - 1, point.y})
		}
		return next
	}
}

func (t *Task) canPush(current []Point, jet string) bool {
	if jet == ">" {
		for _, point := range current {
			if point.x + 1 >= 7 {
				return false
			}
			for _, set := range t.grid {
				next := point
				next.x += 1
				if next == set {
					return false
				}
			}
		}
		return true
	} else {
		for _, point := range current {
			if point.x - 1 < 0 {
				return false
			}
			for _, set := range t.grid {
				next := point
				next.x -= 1
				if next == set {
					return false
				}
			}
		}
		return true
	}
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := Task{
		grid: []Point{},
		jets: []string{},
		shapes: [][]Point{},
		width: 7,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			parts := strings.Split(in, "")
			for _, p := range parts {
				t.jets = append(t.jets, p)
			}
		}
	}

	t.shapes = append(t.shapes,
		[]Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		[]Point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
		[]Point{{2, 2}, {2, 1}, {2, 0}, {1, 0},{0, 0}},
		[]Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		[]Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}})

	return t
}

func (t *Task) printGrid() {
	screen := make([][]string, 25)
	for i := range screen {
		screen[i] = make([]string, 7)
	}
	for i := range screen {
		for j := range screen[i] {
			// screen[i][j] = fmt.Sprintf("%d, %d", i, j)
			screen[i][j] = "."
		}
	}
	for _, point := range t.grid {
		screen[point.y][point.x] = "#"
	}
	for i := len(screen)-1; i >= 0; i-- {
		fmt.Println(screen[i])
	}
}