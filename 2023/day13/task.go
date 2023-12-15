package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2023/day13/input.txt"

type task struct {
	grids []grid
}

type grid struct {
	grid [][]string
}

func main() {
	start := time.Now()

	t := parse()
	part1 := t.sumReflections(false)
	part2 := t.sumReflections(true)

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (t *task) sumReflections(smudge bool) int {
	sum := 0
grid:
	for _, g := range t.grids {
	horizontal:
		for x := 1; x < len(g.grid[0]); x++ {
			count := 0
			mistakes := 0
			for {
				x1 := x + count
				x2 := x - count - 1
				if x1 >= len(g.grid[0]) || x2 < 0 {
					break
				}
				for y := 0; y < len(g.grid); y++ {
					if g.grid[y][x1] != g.grid[y][x2] {
						mistakes++
						if (smudge && mistakes > 1) || !smudge {
							continue horizontal
						}
					}
				}
				count++
			}
			if (smudge && mistakes == 1) || !smudge {
				sum += x
				continue grid
			}
		}
	vertical:
		for y := 1; y < len(g.grid); y++ {
			count := 0
			mistakes := 0
			for {
				y1 := y + count
				y2 := y - count - 1
				if y1 >= len(g.grid) || y2 < 0 {
					break
				}
				for x := 0; x < len(g.grid[0]); x++ {
					if g.grid[y1][x] != g.grid[y2][x] {
						mistakes++
						if (smudge && mistakes > 1) || !smudge {
							continue vertical
						}
					}
				}
				count++
			}
			if (smudge && mistakes == 1) || !smudge {
				sum += y * 100
				continue grid
			}
		}
	}
	return sum
}

func parse() task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{
		[]grid{},
	}

	g := grid{grid: [][]string{}}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			g.grid = append(g.grid, strings.Split(in, ""))
		} else {
			t.grids = append(t.grids, g)
			g = grid{grid: [][]string{}}
		}
	}

	t.grids = append(t.grids, g)

	return t
}
