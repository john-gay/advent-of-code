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

var input = "2025/day2/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Range struct {
	start int
	end   int
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ranges := []Range{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, ",")
			for _, part := range p {
				r := strings.Split(part, "-")
				start, _ := strconv.Atoi(r[0])
				end, _ := strconv.Atoi(r[1])
				ranges = append(ranges, Range{start: start, end: end})
			}
		}
	}

	part1, part2 := 0, 0
	
	for _, r := range ranges {
		for n := r.start; n <= r.end; n++ {
			nAsStr := strconv.Itoa(n)
			chunks := []int{2, len(nAsStr)}
			for i := 3; i < len(nAsStr); i += 2 {
				chunks = append(chunks, i)
			}
			for _, chunk := range chunks {
				if len(nAsStr) % chunk == 0 {
					chunkSize := 0
					if chunk == 1 {
						chunkSize = 1
					} else {
						chunkSize = len(nAsStr) / chunk
					}
					
					allMatch := false
					for i := 1; i < chunk; i++ {
						if nAsStr[0:chunkSize] != nAsStr[i*chunkSize:(i+1)*chunkSize] {
							allMatch = false
							break
						}
						allMatch = true
					}
					if allMatch {
						part2 += n
						if chunk == 2 {
							part1 += n
						}
						break
					}
				}
			}
		}
	}

	return part1, part2
}
