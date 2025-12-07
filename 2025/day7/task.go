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

var input = "2025/day7/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Grid struct {
	x int
	y int
}

var splitters = map[string]Grid{}
var pathSplitLoc = map[string]bool{}
var paths = map[string]int{}
var maxY = 0

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start, y := 0, 0

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        in := scanner.Text()
        if in != "" {
            p := strings.Split(in, "")
			for x, v := range p {
				if v == "^" {
					splitters[fmt.Sprintf("%d,%d", x, y)] = Grid{x: x, y: y}
				}
				if v == "S" {
					start = x
				}
			}
            y++
        }
    }
	maxY = y

	part1 := 0
	part2 := 0

	followPath(start, 0)

	for range pathSplitLoc {
		part1++
	}

	paths[fmt.Sprintf("%d", start)] = 1
	followQuantumPath(0)

	for _, v := range paths {
		part2 += v
	}

	return part1, part2
}

func followPath(x, y int) {
	currentY := y
	for {
		if _, ok := splitters[fmt.Sprintf("%d,%d", x, currentY)]; ok {
			if _, visited := pathSplitLoc[fmt.Sprintf("%d,%d", x, currentY)]; visited {
				return
			}
			pathSplitLoc[fmt.Sprintf("%d,%d", x, currentY)] = true
			followPath(x - 1, currentY)
			followPath(x + 1, currentY)
			return
		}
		
		if currentY > maxY {
			break
		}
		currentY++
	}
}


func followQuantumPath(y int) {
	for y <= maxY {
		newPaths := map[string]int{}
		for p, v := range paths {
			x, _ := strconv.Atoi(p)
			if _, ok := splitters[fmt.Sprintf("%d,%d", x, y)]; ok {
				newPaths[fmt.Sprintf("%d", x - 1)] += v
				newPaths[fmt.Sprintf("%d", x + 1)] += v
			} else {
				newPaths[p] += v
			}	
		}
		paths = newPaths
		y++
	}
}