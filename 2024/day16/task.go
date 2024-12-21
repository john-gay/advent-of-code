package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var input = "2024/day16/input.txt"

type coord struct {
	x, y int
}

type state struct {
	pos coord
	dir coord
}

type queueItem struct {
	state state
	cost  int
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
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1, part2 := 0, 0

	grid := map[coord]string{}
	moves := []string{}
	i := 0
	start, end := coord{0, 0}, coord{0, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "")
			if p[0] == "#" {
				for j := 0; j < len(p); j++ {
					if p[j] == "." {
						continue
					}

					grid[coord{j, i}] = p[j]

					if p[j] == "S" {
						start = coord{j, i}
					} 
					if p[j] == "E" {
						end = coord{j, i}
					}
				}
				i++
			} else {
				moves = append(moves, p...)
			}
		}
	}

	part1 = findMinCostPath(grid, start, end)

	paths := findPathsWithCost(grid, start, end, part1)
	for _, path := range paths {
		for _, p := range path {
			grid[p] = "X"
		}
	}
	if len(paths) > 0 {
		printGrid(grid)
	} else {
		fmt.Println("No paths found")
	}

	for _, val := range grid {
		if val == "X" {
			part2++
		}
	}

	return part1, part2
}

func findMinCostPath(grid map[coord]string, start, end coord) int {
    queue := []queueItem{{
        state: state{pos: start, dir: coord{1, 0}},
        cost:  0,
    }}

    visited := make(map[state]bool)
    costs := make(map[state]int)
    prev := make(map[state]state)

    costs[state{start, coord{0, 0}}] = 0

    directions := []coord{
        {0, 1},
        {0, -1},
        {1, 0},
        {-1, 0},
    }

    for len(queue) > 0 {
        minIdx := 0
        for i := range queue {
            if queue[i].cost < queue[minIdx].cost {
                minIdx = i
            }
        }
        current := queue[minIdx]
        queue = append(queue[:minIdx], queue[minIdx+1:]...)

        if current.state.pos == end {
            return current.cost
        }

        if visited[current.state] {
            continue
        }
        visited[current.state] = true

        for _, d := range directions {
            next := coord{current.state.pos.x + d.x, current.state.pos.y + d.y}
            
            if val, ok := grid[next]; ok && val == "#" {
                continue
            }

            newCost := current.cost + 1
            if current.state.dir != d && 
				(current.state.dir.x*d.x + current.state.dir.y*d.y == 0) {
                newCost += 1000
            }

            nextState := state{next, d}
            if cost, exists := costs[nextState]; !exists || newCost < cost {
                costs[nextState] = newCost
                prev[nextState] = current.state
                queue = append(queue, queueItem{
                    state: nextState,
                    cost:  newCost,
                })
            }
        }
    }

    return -1
}

func findPathsWithCost(grid map[coord]string, start, end coord, targetCost int) [][]coord {
    type queueItem struct {
        state state
        cost  int
        path  []coord
    }

    paths := [][]coord{}
    stateCosts := make(map[state]int)
    queue := []queueItem{{
        state: state{pos: start, dir: coord{1, 0}},
        cost:  0,
        path:  []coord{start},
    }}

    directions := []coord{
        {0, 1},
        {0, -1},
        {1, 0},
        {-1, 0},
    }

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current.state.pos == end && current.cost == targetCost {
            paths = append(paths, current.path)
            continue
        }

        if current.cost > targetCost {
            continue
        }

        if prevCost, exists := stateCosts[current.state]; exists && current.cost > prevCost {
            continue
        }

        stateCosts[current.state] = current.cost

        for _, d := range directions {
            next := coord{current.state.pos.x + d.x, current.state.pos.y + d.y}
            
            if val, ok := grid[next]; ok && val == "#" {
                continue
            }

            newCost := current.cost + 1
            if  current.state.dir != d && 
				(current.state.dir.x*d.x + current.state.dir.y*d.y == 0) {
                newCost += 1000
            }

            if newCost > targetCost {
                continue
            }

            newPath := make([]coord, len(current.path))
            copy(newPath, current.path)
            newPath = append(newPath, next)

            nextState := state{pos: next, dir: d}
            if prevCost, exists := stateCosts[nextState]; !exists || newCost <= prevCost {
                queue = append(queue, queueItem{
                    state: nextState,
                    cost:  newCost,
                    path:  newPath,
                })
            }
        }
    }

    return paths
}

func printGrid(grid map[coord]string) {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for pos := range grid {
		minX = min(minX, pos.x)
		minY = min(minY, pos.y)
		maxX = max(maxX, pos.x)
		maxY = max(maxY, pos.y)
	}

	fmt.Println()

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if val, exists := grid[coord{x, y}]; exists {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}