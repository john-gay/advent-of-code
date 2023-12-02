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

var input = "2023/day2/input.txt"

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

	total, sumPower := 0, 0 
	maxR, maxG, maxB := 12, 13, 14

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			gameMaxR, gameMaxG, gameMaxB := 0, 0, 0
			p := strings.Split(in, ":")
			game, _ := strconv.Atoi(strings.Split(p[0], " ")[1])

			sets := strings.Split(p[1], ";")
			for _, set := range sets {
				setParts := strings.Split(set, ",")
				for _, numColour := range setParts {
					numColourParts := strings.Split(strings.Trim(numColour, " "), " ")
					switch numColourParts[1] {
					case "red":
						gameMaxR = getMaxColour(numColourParts[0], gameMaxR)
					case "green":
						gameMaxG = getMaxColour(numColourParts[0], gameMaxG)
					case "blue":
						gameMaxB = getMaxColour(numColourParts[0], gameMaxB)
					}
				}
			}
			sumPower += gameMaxR * gameMaxG * gameMaxB
			if gameMaxR <= maxR && gameMaxG <= maxG && gameMaxB <= maxB {
				total += game
			}
		}
	}

	return total, sumPower
}

func getMaxColour(numString string, max int) int {
	n, _ := strconv.Atoi(numString)
	if n > max {
		max = n
	}
	return max
}
