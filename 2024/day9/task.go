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

var input = "2024/day9/input.txt"

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

	diskMap := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			diskMap = strings.Split(in, "")
		}
	}

	f := createFile(diskMap)
	sortFile(f)
	part1 = calculateScore(f)

	f2 := createFile(diskMap)
	sortFileBlocks(f2)
	part2 = calculateScore(f2)

	return part1, part2
}

func calculateScore(file []string) int {
	score := 0
	for i := 0; i < len(file); i++ {
		if file[i] != "." {
			val, _ := strconv.Atoi(file[i])
			score += i * val
		}
	}
	return score
}

func createFile(diskMap []string) []string {
	file := []string{}
	for i := 0; i < len(diskMap); i++ {
		amount, _ := strconv.Atoi(diskMap[i])
		if i % 2 == 0 {
			for j := 0; j < amount; j++ {
				file = append(file, fmt.Sprintf("%d", i / 2))
			}
		} else {
			for j := 0; j < amount; j++ {
				file = append(file, ".")
			}
		}
	}
	return file
}

func sortFileBlocks(file []string) {
	var step int
	for i := len(file) - 1; i > 0; i -= step {
		step = 1
		value := file[i]
		for j := i - 1; j >= 0; j-- {
			if file[j] == value {
				step++
			} else {
				break
			}
		}

		for j := 0; j < i-1; j++ {
			if file[j] == "." {
				count := 0
				for k := 0; k < step; k++ {
					if file[j+k] == "." {
						count++
					}
				}
				if count >= step {
					for k := 0; k < step; k++ {
						file[j+k] = value
						file[i-k] = "."
					}
					break
				}
			}
		}
	}
}

func sortFile(file []string) {
	for i := len(file) - 1; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if file[j] == "." {
				file[j] = file[i]
				file[i] = "."
				break
			}
		}
	}
}