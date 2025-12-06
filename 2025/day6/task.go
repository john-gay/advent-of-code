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

var input = "2025/day6/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d. Code lost", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Sum struct {
	values []string
	operator string
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sums := []Sum{}
    lines := []string{}
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        in := scanner.Text()
        if in != "" {
            lines = append(lines, in)
        }
    }

	lastLine := lines[len(lines)-1]
	blocks := []int{}
	gap := 0
	firstPassed := false
	for _, char := range lastLine {
		gap++
		if char == '+' || char == '*' {
			if firstPassed {
				blocks = append(blocks, gap-1)
			} else {
				firstPassed = true
			}
			gap = 0
		}
	}
	blocks = append(blocks, gap+1)
	blockStart := 0
	for _, columns := range blocks {
		lineIndex := 0
		sum := Sum{values: make([]string, columns)}
		for i := blockStart; i < blockStart+columns; i++ {
			for _, line := range lines {
				lineParts := strings.Split(line, "")
				if lineParts[i] == "+" || lineParts[i] == "*" {
					sum.operator = lineParts[i]
				} else if lineParts[i] != " " {
					sum.values[i-blockStart] += lineParts[i]
				}
			}
			lineIndex++
		}
		blockStart += columns + 1
		sums = append(sums, sum)
	}

	part1 := 0
	part2 := 0

	for _, sum := range sums {
		total := 0
		if sum.operator == "+" {
			for _, v := range sum.values {
				val, _ := strconv.Atoi(v)
				total += val
			}
		} else if sum.operator == "*" {
			total = 1
			for _, v := range sum.values {
				val, _ := strconv.Atoi(v)
				total *= val
			}
		}
		part2 += total
	}

	return part1, part2
}
