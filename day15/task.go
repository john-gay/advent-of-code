package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

var input = "day15/input.txt"

type Task struct {
	sensors map[Point]Sensor
	row     []int
}

type Sensor struct {
	location Point
	nearest  Point
	distance int
}

type Point struct {
	x int
	y int
}

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d, test = 26, not 4137843", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	t := readInput()

	//fmt.Println(fmt.Sprintf("%+v", t))

	//part1 := t.freeSpaces(10)
	part1 := t.freeSpaces(2000000)

	part2 := 0

	return part1, part2
}

func (t *Task) freeSpaces(y int) int {
	total := 0

	minX := math.MaxInt
	maxX := math.MinInt
	for _, sensor := range t.sensors {
		minX = minimum(sensor.location.x-sensor.distance-1, minX)
		maxX = maximum(sensor.location.x+sensor.distance+1, maxX)
		minX = minimum(sensor.nearest.x-sensor.distance-1, minX)
		maxX = maximum(sensor.nearest.x+sensor.distance+1, maxX)
	}

outerLoop:
	for x := minX; x <= maxX; x++ {
		point := Point{x, y}
		for _, sensor := range t.sensors {
			//fmt.Println(point, sensor.location, manhatDist(point, sensor.location), sensor.distance)
			if point != sensor.nearest && manhatDist(point, sensor.location) <= sensor.distance {
				//fmt.Println("not empty")
				total++
				continue outerLoop
			}
		}
	}

	return total
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := Task{
		sensors: map[Point]Sensor{},
	}

	sensor := Sensor{}
	minX := math.MaxInt
	maxX := math.MinInt

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			var lx, ly, nx, ny int
			_, _ = fmt.Sscanf(in, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &lx, &ly, &nx, &ny)
			sensor.location = Point{lx, ly}
			sensor.nearest = Point{nx, ny}
			sensor.distance = manhatDist(sensor.location, sensor.nearest)
			t.sensors[sensor.location] = sensor
			sensor = Sensor{}

			minX = minimum(minX, nx)
			minX = minimum(minX, lx)
			maxX = maximum(maxX, nx)
			maxX = maximum(maxX, lx)
		}
	}
	t.row = []int{minX, maxX}

	return t
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func manhatDist(start, end Point) int {
	dx := math.Abs(float64(start.x - end.x))
	dy := math.Abs(float64(start.y - end.y))
	return int(dx + dy)
}
