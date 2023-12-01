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

var input = "2022/day10/input.txt"

type task struct {
	cycles map[int]int
	cycle  int
	x      int
	screen [][]string
	row    []string
}

func main() {
	start := time.Now()

	part1 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{
		cycles: map[int]int{},
		cycle:  0,
		x:      1,
		screen: [][]string{},
		row:    []string{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			switch p[0] {
			case "noop":
				t.newCycle()
			case "addx":
				v, _ := strconv.Atoi(p[1])
				t.add(v)
			}
		}
	}

	t.printScreen()

	return t.sumStrength()
}

func (t *task) newCycle() {
	t.cycle++
	t.cycles[t.cycle] = t.x
	t.drawPosition()
}

func (t *task) drawPosition() {
	position := t.cycle%40 - 1
	if t.x == position || t.x == position-1 || t.x == position+1 {
		t.row = append(t.row, "#")
	} else {
		t.row = append(t.row, ".")
	}

	if len(t.row) == 40 {
		t.screen = append(t.screen, t.row)
		t.row = []string{}
	}
}

func (t *task) add(x int) {
	for i := 0; i < 2; i++ {
		t.newCycle()
	}
	t.x += x
}

func (t *task) sumStrength() int {
	sum := 0

	for i := 20; i < len(t.cycles); i += 40 {
		sum += i * t.cycles[i]
	}

	return sum
}

func (t *task) printScreen() {
	for _, row := range t.screen {
		fmt.Println(row)
	}
}
