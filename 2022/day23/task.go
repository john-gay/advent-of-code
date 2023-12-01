package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

var input = "2022/day23/input.txt"

type Task struct {
	grid      map[Point]bool
	round     int
	direction map[int]string
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
	task1 := readInput()
	part1 := task1.spreadOutRounds(10)

	task2 := readInput()
	part2 := task2.spreadOutFully()

	return part1, part2
}

func (t *Task) spreadOutFully() int {
	for {
		moves := t.spreadOut()
		if moves == 0 {
			break
		}
	}
	return t.round
}

func (t *Task) spreadOutRounds(rounds int) int {
	for i := 0; i < rounds; i++ {
		t.spreadOut()
	}

	return t.calcScore()
}

func (t *Task) spreadOut() int {
	nextMoves := map[Point][]Point{}
	for elf := range t.grid {
		if t.isAlone(elf) {
			continue
		}

		proposedMove, canMove := t.nextMove(elf)
		if canMove {
			nextMoves[proposedMove] = append(nextMoves[proposedMove], elf)
		}
	}

	for nextMove, elfves := range nextMoves {
		if len(elfves) > 1 {
			continue
		}

		t.grid[nextMove] = true
		delete(t.grid, elfves[0])
	}
	t.round++
	return len(nextMoves)
}

func (t *Task) nextMove(p Point) (Point, bool) {
	for dir := 0; dir < len(t.direction); dir++ {
		direction := t.direction[(t.round+dir)%len(t.direction)]
		switch direction {
		case "N":
			if !t.grid[Point{p.x - 1, p.y - 1}] &&
				!t.grid[Point{p.x, p.y - 1}] &&
				!t.grid[Point{p.x + 1, p.y - 1}] {
				return Point{p.x, p.y - 1}, true
			}
		case "S":
			if !t.grid[Point{p.x - 1, p.y + 1}] &&
				!t.grid[Point{p.x, p.y + 1}] &&
				!t.grid[Point{p.x + 1, p.y + 1}] {
				return Point{p.x, p.y + 1}, true
			}
		case "W":
			if !t.grid[Point{p.x - 1, p.y - 1}] &&
				!t.grid[Point{p.x - 1, p.y}] &&
				!t.grid[Point{p.x - 1, p.y + 1}] {
				return Point{p.x - 1, p.y}, true
			}
		case "E":
			if !t.grid[Point{p.x + 1, p.y - 1}] &&
				!t.grid[Point{p.x + 1, p.y}] &&
				!t.grid[Point{p.x + 1, p.y + 1}] {
				return Point{p.x + 1, p.y}, true
			}
		}
	}
	return Point{}, false
}

func (t *Task) isAlone(p Point) bool {
	for _, surroundingPoint := range getSurrounding(p) {
		if t.grid[surroundingPoint] {
			return false
		}
	}
	return true
}

func (t *Task) calcScore() int {
	maxX, minX, maxY, minY := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt
	for elf := range t.grid {
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.x < minX {
			minX = elf.x
		}
		if elf.y > maxY {
			maxY = elf.y
		}
		if elf.y < minY {
			minY = elf.y
		}
	}

	score := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if !t.grid[Point{x, y}] {
				score++
			}
		}
	}
	return score
}

func getSurrounding(p Point) []Point {
	return []Point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y + 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y + 1},
		{p.x - 1, p.y - 1},
	}
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	task := Task{
		grid:      map[Point]bool{},
		direction: map[int]string{0: "N", 1: "S", 2: "W", 3: "E"},
		round:     0,
	}

	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			for i, c := range in {
				if string(c) == "#" {
					task.grid[Point{i, y}] = true
				}
			}
			y++
		}
	}

	return task
}
