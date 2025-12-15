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

var input = "2025/day12/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Dimension struct {
	width, height int
	unique int
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	shapeSizes := []int{}
	presents := map[Dimension][]int{}

	scanner := bufio.NewScanner(file)
	shapeSize := 0
	count := 0
	for scanner.Scan() {
		in := scanner.Text()
		count++
		if in != "" {
			parts := strings.Split(in, "")
			if parts[1] == ":" {
				shapeSize = 0
				continue
			}
			if parts[0] == "#" || parts[0] == "." {
				for _, p := range parts {
					if p == "#" {
						shapeSize++
					}
				}
				continue
			}

			dimensionParts := strings.Split(in, ": ")
			widthHeight := strings.Split(dimensionParts[0], "x")
			width, _ := strconv.Atoi(widthHeight[0])
			height, _ := strconv.Atoi(widthHeight[1])
			dim := Dimension{width: width, height: height, unique: count}
			counts := []int{}
			indexParts := strings.Split(dimensionParts[1], " ")
			for _, i := range indexParts {
				idx, _ := strconv.Atoi(i)
				counts = append(counts, idx)
			}
			presents[dim] = counts
		} else {
			if shapeSize > 0 {
				shapeSizes = append(shapeSizes, shapeSize)
				shapeSize = 0
			}
		}
	}

	part1, part2 := 0, 0

	for present, counts := range presents {
		total := present.width * present.height

		size := 0
		for index, count := range counts {
			shapeSize := shapeSizes[index] * count
			size += shapeSize
		}

		if size <= total {
			part1++
		}
	}

	return part1, part2
}
