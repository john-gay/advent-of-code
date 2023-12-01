package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var input = "2022/day18/input.txt"

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

		for next := range getNeighbours(p) {
			if checked[next] {
				continue
			}

			if next.x < -1 || next.x > 20 || next.y < -1 || next.y > 20 || next.z < -1 || next.z > 20 {
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
		nCount := 6
		for n := range getNeighbours(point) {
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
