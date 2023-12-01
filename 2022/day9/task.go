package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var input = "2022/day9/input.txt"

type rope struct {
	knots   []coord
	visited map[string]bool
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

	p1 := tailVisited(2)
	p2 := tailVisited(10)

	return p1, p2
}

func tailVisited(numberOfKnots int) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r := &rope{
		visited: map[string]bool{},
	}

	for i := 0; i < numberOfKnots; i++ {
		r.knots = append(r.knots, coord{0, 0})
	}

	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			times, _ := strconv.Atoi(p[1])
			switch p[0] {
			case "R":
				r.moveRight(times)
			case "L":
				r.moveLeft(times)
			case "U":
				r.moveUp(times)
			case "D":
				r.moveDown(times)
			}
		}
	}

	return len(r.visited)
}

func (r *rope) moveRight(times int) {
	for i := 0; i < times; i++ {
		r.knots[0].x++
		r.updateKnots()
	}
}

func (r *rope) moveLeft(times int) {
	for i := 0; i < times; i++ {
		r.knots[0].x--
		r.updateKnots()
	}
}

func (r *rope) moveUp(times int) {
	for i := 0; i < times; i++ {
		r.knots[0].y++
		r.updateKnots()
	}
}

func (r *rope) moveDown(times int) {
	for i := 0; i < times; i++ {
		r.knots[0].y--
		r.updateKnots()
	}
}

func (r *rope) updateKnots() {
	for i := 1; i < len(r.knots); i++ {
		r.updateKnot(i)
	}
	r.visit(r.knots[len(r.knots)-1].x, r.knots[len(r.knots)-1].y)
}

func (r *rope) updateKnot(index int) {
	head := &r.knots[index-1]
	tail := &r.knots[index]

	dX := tail.x - head.x
	dY := tail.y - head.y

	if magnitude(dX) > 1 || magnitude(dY) > 1 {
		tail.x += move(dX)
		tail.y += move(dY)
	}
}

func (r *rope) visit(x, y int) {
	r.visited[fmt.Sprintf("%d+%d", x, y)] = true
}

func magnitude(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func move(value int) int {
	if value > 0 {
		return -1
	}
	if value == 0 {
		return 0
	}
	return 1
}
