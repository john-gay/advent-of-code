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

var input = "2024/day13/input.txt"

type equation struct {
	x1, y1, x2, y2, xp, yp int
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

	equations := []equation{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Button A:") {
			aCoords := strings.Split(strings.TrimPrefix(line, "Button A: "), ",")
			ax, _ := strconv.Atoi(strings.TrimPrefix(aCoords[0], "X+"))
			ay, _ := strconv.Atoi(strings.TrimPrefix(aCoords[1], " Y+"))

			scanner.Scan()
			line = scanner.Text()
			bCoords := strings.Split(strings.TrimPrefix(line, "Button B: "), ",")
			bx, _ := strconv.Atoi(strings.TrimPrefix(bCoords[0], "X+"))
			by, _ := strconv.Atoi(strings.TrimPrefix(bCoords[1], " Y+"))

			scanner.Scan()
			line = scanner.Text()
			prizeCoords := strings.Split(strings.TrimPrefix(line, "Prize: "), ",")
			xp, _ := strconv.Atoi(strings.TrimPrefix(prizeCoords[0], "X="))
			yp, _ := strconv.Atoi(strings.TrimPrefix(prizeCoords[1], " Y="))

			equations = append(equations, equation{
				x1: ax,
				y1: ay,
				x2: bx,
				y2: by,
				xp: xp,
				yp: yp,
			})
		}
	}

	for _, equation := range equations {
		part1 += solveWithCramers(equation, 0) // 37128
		part2 += solveWithCramers(equation, 10000000000000)
	}

	return part1, part2
}

func solveWithCramers(eq equation, offset int) int {
	xp, yp := (eq.xp + offset), (eq.yp + offset)

	det := float64(eq.x1*eq.y2 - eq.y1*eq.x2)
	x := float64(xp*eq.y2-yp*eq.x2) / det
	y := float64(yp*eq.x1-xp*eq.y1) / det

	_, xFrac := math.Modf(x)
	_, yFrac := math.Modf(y)
	if xFrac != 0 || yFrac != 0 {
		return 0
	}

	return 3*int(x) + int(y)
}
