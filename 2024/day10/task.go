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

var input = "2024/day10/input.txt"

type coord struct {
	x int
	y int
}

var nines = map[coord]int{}

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

	part1, part2 := 0, 0

	grid := map[coord]int{}
	i := 0

	startingPoints := []coord{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			for j := 0; j < len(p); j++ {
				num, _ := strconv.Atoi(p[j])
				grid[coord{i, j}] = num
				if num == 0 {
					startingPoints = append(startingPoints, coord{i, j})
				}
			}
			i++
		}
	}

	for _, sp := range startingPoints {
		nines = map[coord]int{}
		findTrails(grid, sp, 0)
		part1 += len(nines)
		for _, v := range nines {
			part2 += v
		}
	}

	return part1, part2
}

func findTrails(grid map[coord]int, c coord, value int) {
	if value == 9 {
        nines[c] += 1
		return
    }

    directions := []coord{
        {c.x, c.y + 1},
        {c.x, c.y - 1},
        {c.x + 1, c.y},
        {c.x - 1, c.y},
    }

    for _, dir := range directions {
        if grid[dir] == value+1 {
            findTrails(grid, dir, value+1)
        }
    }

    return
}

func printGrid(grid map[coord]int) {
	maxX, maxY := 0, 0
	for c := range grid {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			fmt.Print(grid[coord{i, j}])
		}
		fmt.Println()
	}
}
