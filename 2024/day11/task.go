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

var input = "2024/day11/input.txt"

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

	stones := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			for _, v := range p {
				n, _ := strconv.Atoi(v)
				stones = append(stones, n)
			}
		}
	}

	part1 = blink(stones, 25)
	part2 = blink(stones, 75)

	return part1, part2
}

func blinkByArray(stones []int, turns int) int {
	for i := 0; i < turns; i++ {
		newStones := make([]int, len(stones))
		padding := 0
		for j := 0; j < len(stones); j++ {
			key := j + padding
			stoneStr := strconv.Itoa(stones[j])
			if stones[j] == 0 {
				newStones[key] = 1
			} else if len(stoneStr)%2 == 0 {
				mid := len(stoneStr) / 2
				left, _ := strconv.Atoi(stoneStr[:mid])
				right, _ := strconv.Atoi(stoneStr[mid:])
				newStones = append(newStones[:key], append([]int{left, right}, newStones[key+1:]...)...)
				padding++
			} else {
				newStones[key] = stones[j] * 2024
			}
		}
		stones = newStones
	}
	return len(stones)
}

func blink(stones []int, turns int) int {
	stoneMap := make(map[int]int)
	for _, value := range stones {
		stoneMap[value]++
	}

	for i := 0; i < turns; i++ {
		newStoneMap := make(map[int]int)

		for value, count := range stoneMap {
			stoneStr := strconv.Itoa(value)

			if value == 0 {
				newStoneMap[1] += count
			} else if len(stoneStr)%2 == 0 {
				mid := len(stoneStr) / 2
				left, _ := strconv.Atoi(stoneStr[:mid])
				right, _ := strconv.Atoi(stoneStr[mid:])

				newStoneMap[left] += count
				newStoneMap[right] += count
			} else {
				newValue := value * 2024
				newStoneMap[newValue] += count
			}
		}

		stoneMap = newStoneMap
	}

	totalPositions := 0
	for _, count := range stoneMap {
		totalPositions += count
	}
	return totalPositions
}
