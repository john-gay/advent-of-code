package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

var input = "2023/day18/input.txt"

type task struct {
	coords       coords
	colourCoords coords
}

type coords struct {
	coords []coord
}

type coord struct {
	x int
	y int
}

func main() {
	start := time.Now()

	t := parse()

	part1 := t.coords.coordArea()
	part2 := t.colourCoords.coordArea()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func parse() task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{
		coords:       coords{},
		colourCoords: coords{},
	}

	pos := coord{0, 0}
	t.coords.coords = append(t.coords.coords, pos)

	colourPos := coord{0, 0}
	t.colourCoords.coords = append(t.colourCoords.coords, colourPos)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			direction := p[0]
			distance, _ := strconv.Atoi(p[1])
			switch direction {
			case "R":
				pos = coord{pos.x + distance, pos.y}
			case "L":
				pos = coord{pos.x - distance, pos.y}
			case "U":
				pos = coord{pos.x, pos.y + distance}
			case "D":
				pos = coord{pos.x, pos.y - distance}
			}
			t.coords.coords = append(t.coords.coords, pos)

			colour := strings.TrimSuffix(strings.TrimPrefix(p[2], "(#"), ")")
			colourDistance := new(big.Int)
			colourDistance.SetString(colour[:len(colour)-1], 16)
			colourDirection := string(colour[len(colour)-1])
			switch colourDirection {
			case "0":
				colourPos = coord{colourPos.x + int(colourDistance.Int64()), colourPos.y}
			case "1":
				colourPos = coord{colourPos.x, colourPos.y - int(colourDistance.Int64())}
			case "2":
				colourPos = coord{colourPos.x - int(colourDistance.Int64()), colourPos.y}
			case "3":
				colourPos = coord{colourPos.x, colourPos.y + int(colourDistance.Int64())}
			}
			t.colourCoords.coords = append(t.colourCoords.coords, colourPos)
		}
	}

	return t
}

func (c *coords) coordArea() int {
	area := 0
	perimeter := 0
	for i := 0; i < len(c.coords)-1; i++ {
		area += c.coords[i].x*c.coords[i+1].y - c.coords[i+1].x*c.coords[i].y
		perimeter += int(
			math.Sqrt(
				math.Pow(float64(c.coords[i+1].x-c.coords[i].x), 2) +
					math.Pow(float64(c.coords[i+1].y-c.coords[i].y), 2)),
		)
	}

	area = int(math.Abs(float64(area)) / 2)

	return area + perimeter/2 + 1
}
