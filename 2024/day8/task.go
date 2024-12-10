package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2024/day8/input.txt"

type coord struct {
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
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1, part2 := 0, 0

	grid := map[string][]coord{}
	maxX, maxY := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			maxX = len(p)
			maxY++
			for i, v := range p {
				if v != "." {
					grid[v] = append(grid[v], coord{x: i, y: maxY-1})
				}
			}
		}
	}

	antennaMap := map[coord]bool{}
	for _, v := range grid {
		for i := 0; i < len(v) - 1; i++ {
			for j := i+1; j < len(v); j++ {
				antenna, _ := getNewCoords(v[i], v[j], maxX, maxY)
				for _, a := range antenna {
					antennaMap[a] = true
				}
			}
		}
	}

	part1 = len(antennaMap)

	antennaMapWithRes := map[coord]bool{}
	for _, v := range grid {
		for i := 0; i < len(v) - 1; i++ {
			for j := i+1; j < len(v); j++ {
				count := 0
				a1, a2 := v[i], v[j]

				resLoop:
				for {
					antennas, diffs := getNewCoords(a1, a2, maxX, maxY)
					count++

					newAntennas := []coord{}
					fmt.Println(antennas)
					for i, a := range antennas {
						antennaMapWithRes[a] = true

						newAntennas = append(newAntennas, coord{x: a.x + diffs[i][0], y: a.y + diffs[i][1]})
						
					}
					fmt.Println("newAntennas", newAntennas)
					antennas = newAntennas
					if len(antennas) == 0 || count == 3 {
						break resLoop
					}
				}
			}
		}
	}

	part2 = len(antennaMapWithRes)

	return part1, part2
}

// Returns all new coords
func getNewCoords(a coord, b coord, maxX, maxY int) ([]coord, [][]int) {
	coords := []coord{}
	diffs := [][]int{}

	diffX := getAbsDiff(a.x, b.x)
	diffY := getAbsDiff(a.y, b.y)
	
	x1, x2, y1, y2 := 0, 0, 0, 0
	dx1, dx2, dy1, dy2 := 0, 0, 0, 0
	if a.x < b.x {
		x1 = a.x - diffX
		x2 = b.x + diffX
		dx1 = -diffX
		dx2 = diffX
	} else {
		x1 = a.x + diffX
		x2 = b.x - diffX
		dx1 = diffX
		dx2 = -diffX
	}
	if a.y < b.y {
		y1 = a.y - diffY
		y2 = b.y + diffY
		dy1 = -diffY
		dy2 = diffY
	} else {
		y1 = a.y + diffY
		y2 = b.y - diffY
		dy1 = diffY
		dy2 = -diffY
	}

	if x1 >= 0 && x1 < maxX && y1 >= 0 && y2 < maxY {
		coords = append(coords, coord{x: x1, y: y1})
		diffs = append(diffs, []int{dx1, dy1})
	}
	
	if x2 >= 0 && x2 < maxX && y2 >= 0 && y2 < maxY {
		coords = append(coords, coord{x: x2, y: y2})
		diffs = append(diffs, []int{dx2, dy2})
	}

	return coords, diffs
}

func getAbsDiff(a int, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
