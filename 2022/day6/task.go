package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	marker1 := 0
	marker2 := 0

	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			for i := 0; i < len(in); i++ {
				if marker1 == 0 {
					marker1 = checkForMarker(in, i, i+4)
				}
				if marker2 == 0 {
					marker2 = checkForMarker(in, i, i+14)
				}
				if marker1 != 0 && marker2 != 0 {
					break
				}
			}
		}
	}

	return marker1, marker2
}

func checkForMarker(input string, index, limit int) int {
	if limit == index {
		return index
	}
	for i := 1; i < limit-index; i++ {
		if input[index] == input[index+i] {
			return 0
		}
	}
	return checkForMarker(input, index+1, limit)
}
