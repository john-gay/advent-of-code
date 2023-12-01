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

var input = "2022/day8/input.txt"

type Dir struct {
	name        string
	parent      *Dir
	directories []*Dir
	files       []File
	totalSize   int
}

type File struct {
	name string
	size int
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

	grid := readInput()

	p1 := countVisible(grid)
	p2 := bestScenicScore(grid)

	return p1, p2
}

func readInput() [][]int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]int{}

	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			row := []int{}
			cells := strings.Split(in, "")
			for _, cell := range cells {
				n, _ := strconv.Atoi(cell)
				row = append(row, n)
			}
			grid = append(grid, row)
		}
	}
	return grid
}

func countVisible(grid [][]int) int {
	count := 0

	for i, row := range grid {
		for j, _ := range row {
			if isVisible(i, j, grid) {
				count++
			}
		}
	}

	return count
}

func isVisible(x, y int, grid [][]int) bool {
	if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[0])-1 {
		return true
	}

	cell := grid[x][y]

	for i := 0; i < x; i++ {
		if cell > grid[i][y] {
			if x == i+1 {
				return true
			}
		} else {
			break
		}
	}

	for i := len(grid[0]) - 1; x < i; i-- {
		if cell > grid[i][y] {
			if x == i-1 {
				return true
			}
		} else {
			break
		}
	}

	for j := 0; j < y; j++ {
		if cell > grid[x][j] {
			if y == j+1 {
				return true
			}
		} else {
			break
		}
	}

	for j := len(grid[0]) - 1; y < j; j-- {
		if cell > grid[x][j] {
			if y == j-1 {
				return true
			}
		} else {
			break
		}
	}

	return false
}

func bestScenicScore(grid [][]int) int {
	scenicScore := 0

	for i, row := range grid {
		for j, _ := range row {
			newScore := calcScenicScore(i, j, grid)
			if newScore > scenicScore {
				scenicScore = newScore
			}
		}
	}

	return scenicScore
}

func calcScenicScore(x, y int, grid [][]int) int {
	ls, rs, us, ds := 0, 0, 0, 0
	cell := grid[x][y]

	for i := x - 1; i >= 0; i-- {
		us++
		if cell <= grid[i][y] {
			break
		}
	}

	for i := x + 1; i < len(grid[0]); i++ {
		ds++
		if cell <= grid[i][y] {
			break
		}
	}

	for j := y - 1; j >= 0; j-- {
		ls++
		if cell <= grid[x][j] {
			break
		}
	}

	for j := y + 1; j < len(grid); j++ {
		rs++
		if cell <= grid[x][j] {
			break
		}
	}
	return ls * rs * us * ds
}
