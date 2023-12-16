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

var input = "2023/day15/input.txt"

type lens struct {
	key   string
	focus int
}

func main() {
	start := time.Now()

	part1 := reindeerAscii()
	part2 := reindeerHashmap()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func reindeerHashmap() int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hashmap := map[int][]lens{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			for _, p := range strings.Split(in, ",") {
				if strings.Contains(p, "-") {
					parts := strings.Split(p, "-")
					box := 0
					for _, r := range parts[0] {
						box = getReindeerAscii(box, r)
					}
					currentBox := hashmap[box]
					index := -1
					for i, l := range currentBox {
						if l.key == parts[0] {
							index = i
							break
						}
					}
					if index >= 0 {
						hashmap[box] = append(hashmap[box][:index], hashmap[box][index+1:]...)
					}
				} else {
					parts := strings.Split(p, "=")
					box := 0
					for _, r := range parts[0] {
						box = getReindeerAscii(box, r)
					}
					currentBox := hashmap[box]
					focus, _ := strconv.Atoi(parts[1])
					index := -1
					for i, l := range currentBox {
						if l.key == parts[0] {
							index = i
							break
						}
					}
					newLens := lens{parts[0], focus}
					if index >= 0 {
						hashmap[box][index] = newLens
					} else {
						hashmap[box] = append(hashmap[box], newLens)
					}
				}
			}
		}
	}

	sum := 0
	for key, value := range hashmap {
		for i, l := range value {
			sum += (key + 1) * (i + 1) * l.focus
		}
	}

	return sum
}

func reindeerAscii() int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	current := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			for _, r := range in {
				if r == ',' {
					sum += current
					current = 0
				} else {
					current = getReindeerAscii(current, r)
				}
			}
		}
	}

	sum += current

	return sum
}

func getReindeerAscii(current int, r int32) int {
	current += int(r)
	current *= 17
	current = current % 256
	return current
}
