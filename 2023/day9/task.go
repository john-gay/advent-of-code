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

var input = "2023/day9/input.txt"

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

	sumNext, sumPrev := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			values := []int{}
			for _, s := range p {
				val, _ := strconv.Atoi(s)
				values = append(values, val)
			}
			sumNext += predictNextValue(values)
			sumPrev += predictPrevValue(values)
		}
	}

	return sumNext, sumPrev
}

func predictNextValue(values []int) int {
	diffs := []int{}
	allZero := true
	for i := 0; i < len(values) -1; i++ {
		diff := values[i+1]-values[i]
		diffs = append(diffs, diff)
		if diff != 0 {
			allZero = false
		}
	}
	if allZero {
		return values[len(values)-1]
	}

	return values[len(values)-1] + predictNextValue(diffs)
}

func predictPrevValue(values []int) int {
	diffs := []int{}
	allZero := true
	for i := 0; i < len(values) -1; i++ {
		diff := values[i+1]-values[i]
		diffs = append(diffs, diff)
		if diff != 0 {
			allZero = false
		}
	}
	if allZero {
		return values[0]
	}

	return values[0] - predictPrevValue(diffs)
}