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

var input = "2025/day5/input.txt"

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
	fruits := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, "-")
			if len(p) == 2 {
				start, _ := strconv.Atoi(p[0])
				end, _ := strconv.Atoi(p[1])
				ranges = append(ranges, Range{start: start, end: end})
			} else {
				num, _ := strconv.Atoi(p[0])
				fruits = append(fruits, num)
			}
		}
	}

	part1 := 0
	part2 := 0

	for _, fruit := range fruits {
		for _, r := range ranges {
			if fruit >= r.start && fruit <= r.end {
				part1++
				break
			}
		
		}
	}

	reducedRanges := map[string]Range{}
	for _, r := range ranges {
		key := fmt.Sprintf("%d-%d", r.start, r.end)
		reducedRanges[key] = r
	}
	changed := true
	count := 0
    for changed {
        count++
        changed = false
        for _, r := range reducedRanges {
            key := fmt.Sprintf("%d-%d", r.start, r.end)
            for k, existing := range reducedRanges {
                if key != k && (
					(r.start >= existing.start && r.start <= existing.end) || 
					(r.end >= existing.start && r.end <= existing.end) || 
					(existing.start >= r.start && existing.start <= r.end)) {
                    unionRange := Range{
                        start: existing.start,
                        end:   existing.end,
                    }
                    if r.start < existing.start {
                        unionRange.start = r.start
                    }
                    if r.end > existing.end {
                        unionRange.end = r.end
                    }
                    delete(reducedRanges, k)
                    delete(reducedRanges, key)
                    reducedRanges[fmt.Sprintf("%d-%d", unionRange.start, unionRange.end)] = unionRange
                    changed = true
                    break
                }
            }
            if changed {
                break
            }
        }
        if !changed {
            break
        }
    }

	for _, r := range reducedRanges {
		part2 += r.end - r.start + 1
	}

	return part1, part2
}
