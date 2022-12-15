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

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	t := readInput()

	part1 := t.freeSpaces(2000000)

	part2 := t.tuningFrequency()

	return part1, part2
}

func (t *Task) tuningFrequency() int {
	current := Point{0, 0}
	checkedSensor := Sensor{}

	for {
		match := false
		for _, sensor := range t.sensors {
			if manhatDist(current, sensor.location) <= sensor.distance {
				checkedSensor = sensor
				match = true
				break
			}
		}
		if !match {
			break
		}
        
        jump := checkedSensor.distance - manhatDist(checkedSensor.location, current) + 1
        
        if current.x + jump > 4000000 {
            current.x = 0
            current.y++
        } else {
            current.x += jump
        }
	}

	return current.x * 4000000 + current.y
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
			if point != sensor.nearest && manhatDist(point, sensor.location) <= sensor.distance {
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
