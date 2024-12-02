package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var input = "2024/day2/input.txt"

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
			row := []int{}
			p := strings.Split(in, " ")
			for _, v := range p {
				n, _ := strconv.Atoi(v)
				row = append(row, n)
			}
			grid = append(grid, row)
		}
	}

	safe := 0
	for i := 0; i < len(grid); i++ {
		if isSafe(grid[i]) {
			safe++
		}
	}

	allowOneErrorSafe := 0

OuterLoop:
	for i := 0; i < len(grid); i++ {
		if isSafe(grid[i]) {
			allowOneErrorSafe++
		} else {
			for j := 0; j < len(grid[i]); j++ {
				row := append([]int{}, grid[i]...)
				row = append(row[:j], row[j+1:]...)
				if isSafe(row) {
					allowOneErrorSafe++
					continue OuterLoop
				}
			}
		}
	}

	return safe, allowOneErrorSafe
}

func isSafe(row []int) bool {
	increasing := true
	if row[0] > row[1] {
		increasing = false
	}

	for j := 0; j < len(row)-1; j++ {
		x, y := row[j], row[j+1]
		if math.Abs(float64(x-y)) < 1 || math.Abs(float64(x-y)) > 3 {
			return false
		}
		if increasing {
			if x > y {
				return false
			}
		} else {
			if x < y {
				return false
			}
		}
	}
	return true
}
