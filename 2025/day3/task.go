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

var input = "2025/day3/input.txt"

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

	grid := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			row := []int{}
			for _, part := range p {
				num, _ := strconv.Atoi(part)
				row = append(row, num)
			}
			grid = append(grid, row)
		}
	}

	part1 := findJoltages(grid, 2)
	part2 := findJoltages(grid, 12)

	return part1, part2
}

func findJoltages(grid [][]int, length int) int {
	total := 0
	for y := 0; y < len(grid); y++ {
		stringValue := ""
		currentIndex := -1
		for i := 0; i < length; i++ {
			max := 0
			for index := currentIndex + 1; index < len(grid[y])-length+i+1; index++ {
				val := grid[y][index]
				if val > max {
					max = val
					currentIndex = index
				}
			}
			stringValue += strconv.Itoa(max)
		}
		v, _ := strconv.Atoi(stringValue)
		total += v
	}
	return total
}
