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

var input = "2025/day1/input.txt"

type Move struct {
	direction string
	steps     int
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
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moves := []Move{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			d := p[0]
			steps, _ := strconv.Atoi(strings.Join(p[1:], ""))
			moves = append(moves, Move{direction: d, steps: steps})
		}
	}

	start := 50
	min, max := 0, 99
	size := max - min + 1
	current := start
	part1, part2 := 0, 0

	for _, move := range moves {
		clicks := 0
		
		if move.direction == "L" {
			for i := 1; i <= move.steps; i++ {
				testPos := (current - i + size) % size
				if testPos == 0 {
					clicks++
				}
			}
			current = (current - move.steps%size + size) % size
		} else {
			for i := 1; i <= move.steps; i++ {
				testPos := (current + i) % size
				if testPos == 0 {
					clicks++
				}
			}
			current = (current + move.steps) % size
		}

		if current == 0 {
			part1++
		}
		part2 += clicks
	}

	return part1, part2
}
