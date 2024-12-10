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

var maxX, maxY int

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
	maxX, maxY = 0, 0

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

	part1 = calculateAntennas(grid, false)
	part2 = calculateAntennas(grid, true)

	return part1, part2
}

func calculateAntennas(grid map[string][]coord, resonance bool) int {
	antennaMap := map[coord]bool{}
	for _, v := range grid {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				if resonance {
					antennaMap[v[i]] = true
					antennaMap[v[j]] = true
				}
				antenna := getNewCoords(v[i], v[j], resonance)
				for _, a := range antenna {
					antennaMap[a] = true
				}
			}
		}
	}
	return len(antennaMap)
}

func getNewCoords(a coord, b coord, resonance bool) ([]coord) {
	coords := []coord{}

	diffX := getAbsDiff(a.x, b.x)
	diffY := getAbsDiff(a.y, b.y)
	
	i := 0
	for {
		i++
		added := false
		x1, x2, y1, y2 := 0, 0, 0, 0
		if a.x < b.x {
			x1 = a.x - diffX * i
			x2 = b.x + diffX * i
		} else {
			x1 = a.x + diffX * i
			x2 = b.x - diffX * i
		}
		if a.y < b.y {
			y1 = a.y - diffY * i
			y2 = b.y + diffY * i
		} else {
			y1 = a.y + diffY * i
			y2 = b.y - diffY * i
		}

		if x1 >= 0 && x1 < maxX && y1 >= 0 && y1 < maxY {
			coords = append(coords, coord{x: x1, y: y1})
			added = true
		}
		if x2 >= 0 && x2 < maxX && y2 >= 0 && y2 < maxY {
			coords = append(coords, coord{x: x2, y: y2})
			added = true
		}
		if resonance {
			if !added {
				break
			}
		} else {
			break
		}
	}

	return coords
}

func getAbsDiff(a int, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
