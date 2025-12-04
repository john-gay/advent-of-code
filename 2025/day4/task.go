package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2025/day4/input.txt"

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

	grid := make(map[string][]int)
	x, y := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			for _, part := range p {
				if part == "@" {
					grid[fmt.Sprintf("%d,%d", x, y)] = []int{x, y}
				}
				x++
			}
			y++
			x = 0
		}
	}

	part1 := 0
	part2 := 0
	
	i := 0
	for {
		reduced := false
		found := make(map[string][]int)
		for _, val := range grid {
			adjacent := [][2]int{
				{val[0]-1, val[1]},
				{val[0]+1, val[1]}, 
				{val[0], val[1]-1},
				{val[0], val[1]+1},
				{val[0]-1, val[1]-1},
				{val[0]-1, val[1]+1},
				{val[0]+1, val[1]-1},
				{val[0]+1, val[1]+1},
			}
			count := 0
			for _, a := range adjacent {
				key := fmt.Sprintf("%d,%d", a[0], a[1])
				if _, exists := grid[key]; exists {
					count++
				}
			}
			if count < 4 {
				if i == 0 {
					part1++
				}
				part2++
				found[fmt.Sprintf("%d,%d", val[0], val[1])] = val
			}
		}
		if len(found) > 0 {
			reduced = true
			for k := range found {
				delete(grid, k)
			}
		}
		if !reduced {
			break
		}
		i++
	}	

	return part1, part2
}
