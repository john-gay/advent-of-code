package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var input = "2022/day17/input.txt"

type Task struct {
	grid   []Point
	jets   []string
	shapes [][]Point
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

	fmt.Println(t)

	part1 := t.play(2022)

	t = readInput()
	part2 := t.play(1000000000000)

	return part1, part2
}

func (t *Task) play(rocks int) int {
	maxHeight := 0
	shapeIndex := 0
	jetIndex := 0

	rockSkip := []int{0, 0}
	// heightSkip := []int{178,337} // Sample skip values
	heightSkip := []int{6575, 9222} // Found by checking printed grid for repeats
	skipped := false

	for rock := 0; rock < rocks; rock++ {
		var shape []Point
		shapeIndex, shape = t.getShape(shapeIndex)
		current := shapeLocation(2, maxHeight+3, shape)

		falling := true
		for falling {
			var jet string
			jetIndex, jet = t.getJet(jetIndex)
			if t.canPush(current, jet) {
				current = t.jetPush(current, jet)
			}
			if t.canFall(current) {
				current = fall(current)
			} else {
				falling = false
			}
		}

		t.setShape(current)
		height := shapeHeight(current)

		if maxHeight < height {
			maxHeight = height
		}

		if height > heightSkip[0] && rockSkip[0] == 0 {
			fmt.Println(fmt.Sprintf("missed %d", heightSkip[0]))
			break
		}
		if height == heightSkip[0] {
			rockSkip[0] = rock
		}

		if height > heightSkip[1] && rockSkip[1] == 0 {
			fmt.Println(fmt.Sprintf("missed %d", heightSkip[1]))
			break
		}
		if height == heightSkip[1] {
			rockSkip[1] = rock
		}

		if rockSkip[1] != 0 && !skipped {
			fmt.Println(fmt.Sprintf("current rock: %d, current height: %d", rock, maxHeight))

			fmt.Println("skiping")
			skipped = true
			dR := rockSkip[1] - rockSkip[0]
			skips := int(math.Floor(float64(rocks-rock) / float64(dR)))
			rock = rock + dR*skips
			maxHeight = height + (heightSkip[1]-heightSkip[0])*skips
			fmt.Println(fmt.Sprintf("dr: %d, skips: %d, new rock: %d, new height: %d", dR, skips, rock, maxHeight))

			for x := 0; x <= 7; x++ {
				t.grid = append(t.grid, Point{x, maxHeight})
			}
		}
	}
	// t.printGrid(maxHeight+1)

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
	for _, point := range current {
		if point.y-1 < 0 {
			return false
		}
		for _, set := range t.grid {
			below := point
			below.y -= 1
			if below == set {
				return false
			}
		}
	}
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
			if point.x+1 >= 7 {
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
			if point.x-1 < 0 {
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
		grid:   []Point{},
		jets:   []string{},
		shapes: [][]Point{},
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
		[]Point{{2, 2}, {2, 1}, {2, 0}, {1, 0}, {0, 0}},
		[]Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		[]Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}})

	return t
}

func (t *Task) printGrid(size int) {
	screen := make([][]string, size)
	for i := range screen {
		screen[i] = make([]string, 7)
	}
	for i := range screen {
		for j := range screen[i] {
			screen[i][j] = "."
		}
	}
	for _, point := range t.grid {
		screen[point.y][point.x] = "#"
	}
	for i := len(screen) - 1; i >= 0; i-- {
		fmt.Println(screen[i])
	}
}
