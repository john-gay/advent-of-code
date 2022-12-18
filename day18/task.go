package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var input = "day18/input.txt"

type Point struct {
	x int
	y int
	z int
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
	points := readInput()

	part1 := calcSurfaceArea(points)

	part2 := calcOuterSurfaceArea(points)

	return part1, part2
}

func calcOuterSurfaceArea(points map[Point]bool) int {
	area := 0
	queue := []Point{{0, 0, 0}}
	checked := map[Point]bool{}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, next := range []Point{
			{p.x + 1, p.y, p.z},
			{p.x - 1, p.y, p.z},
			{p.x, p.y - 1, p.z},
			{p.x, p.y + 1, p.z},
			{p.x, p.y, p.z + 1},
			{p.x, p.y, p.z - 1}} {

			if checked[next] {
				continue
			}

			if next.x < -10 || next.x > 30 || next.y < -10 || next.y > 30 || next.z < -10 || next.z > 30 {
				continue
			}

			if points[next] {
				area++
			} else {
				checked[next] = true
				queue = append(queue, next)
			}
		}
	}

	return area
}

func calcSurfaceArea(points map[Point]bool) int {
	area := 0
	for point := range points {
		neighbours := getNeighbours(point)
		nCount := 6
		for n := range neighbours {
			for otherPoint := range points {
				if otherPoint == n {
					nCount--
					break
				}
			}

		}
		area += nCount
	}

	return area
}

func getNeighbours(point Point) map[Point]bool {
	return map[Point]bool{
		{point.x + 1, point.y, point.z}: true,
		{point.x - 1, point.y, point.z}: true,
		{point.x, point.y + 1, point.z}: true,
		{point.x, point.y - 1, point.z}: true,
		{point.x, point.y, point.z + 1}: true,
		{point.x, point.y, point.z - 1}: true,
	}
}

func readInput() map[Point]bool {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := map[Point]bool{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			point := Point{}
			fmt.Sscanf(in, "%d,%d,%d", &point.x, &point.y, &point.z)
			points[point] = true
		}
	}

	return points
}
