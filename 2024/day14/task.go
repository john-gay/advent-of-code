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

var input = "2024/day14/input.txt"

type Robot struct {
	px, py, vx, vy int
}

type coord struct {
	x, y int
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

	xMax, yMax := 101, 103

	robots := []Robot{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			parts := strings.Split(in, " ")
			pos := strings.TrimPrefix(parts[0], "p=")
			posparts := strings.Split(pos, ",")
			px, _ := strconv.Atoi(posparts[0])
			py, _ := strconv.Atoi(posparts[1])
			vel := strings.TrimPrefix(parts[1], "v=")
			velparts := strings.Split(vel, ",")
			vx, _ := strconv.Atoi(velparts[0])
			vy, _ := strconv.Atoi(velparts[1])

			robots = append(robots, Robot{
				px: px,
				py: py,
				vx: vx,
				vy: vy,
			})
		}
	}

	robotsCopy1 := copyRobots(robots)
	part1 = simulate(robotsCopy1, xMax, yMax, 100, false)

	robotsCopy2 := copyRobots(robots)
	part2 = simulate(robotsCopy2, xMax, yMax, 10000, true)

	printGrid(robotsCopy2, xMax, yMax)

	return part1, part2
}

func copyRobots(robots []Robot) []Robot {
	copied := make([]Robot, len(robots))
	copy(copied, robots)
	return copied
}

func calculateSafety(robots []Robot, xMax, yMax int) int {
	nw, ne, sw, se := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.px == (xMax-1)/2 || robot.py == (yMax-1)/2 {
			continue
		}
		if robot.px < xMax/2 && robot.py < yMax/2 {
			nw++
		} else if robot.px >= xMax/2 && robot.py < yMax/2 {
			ne++
		} else if robot.px < xMax/2 && robot.py >= yMax/2 {
			sw++
		} else {
			se++
		}
	}
	return nw * ne * sw * se
}

func simulate(robots []Robot, xMax, yMax, turns int, findTree bool) int {
	for i := 0; i < turns; i++ {
		for j := 0; j < len(robots); j++ {
			robots[j].px += robots[j].vx
			if robots[j].px < 0 {
				robots[j].px += xMax
			} else if robots[j].px >= xMax {
				robots[j].px -= xMax
			}

			robots[j].py += robots[j].vy
			if robots[j].py < 0 {
				robots[j].py += yMax
			} else if robots[j].py >= yMax {
				robots[j].py -= yMax
			}
		}
		if findTree {
			if hasEightInRow(robots) {
				return i + 1
			}
		}
	}
	return calculateSafety(robots, xMax, yMax)
}

func hasEightInRow(robots []Robot) bool {
	positions := make(map[coord]bool)
	for _, r := range robots {
		positions[coord{r.px, r.py}] = true
	}

	for _, r := range robots {
		hCount := 1
		for i := 1; i < 8; i++ {
			if positions[coord{r.px + i, r.py}] {
				hCount++
			} else {
				break
			}
		}
		for i := 1; i < 8; i++ {
			if positions[coord{r.px - i, r.py}] {
				hCount++
			} else {
				break
			}
		}
		if hCount >= 8 {
			return true
		}
	}
	return false
}

func printGrid(robots []Robot, xMax, yMax int) {
	grid := make([][]string, yMax)
	for i := range grid {
		grid[i] = make([]string, xMax)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for _, robot := range robots {
		if robot.px >= 0 && robot.px < xMax && robot.py >= 0 && robot.py < yMax {
			if grid[robot.py][robot.px] == "." {
				grid[robot.py][robot.px] = "1"
			} else {
				val, _ := strconv.Atoi(grid[robot.py][robot.px])
				grid[robot.py][robot.px] = strconv.Itoa(val + 1)
			}
		}
	}

	fmt.Print("  ")
	for x := 0; x < xMax; x++ {
		fmt.Printf("%d", x%10)
	}
	fmt.Println()

	for y := 0; y < yMax; y++ {
		fmt.Printf("%d ", y%10)
		for x := 0; x < xMax; x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}
