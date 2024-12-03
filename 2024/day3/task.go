package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var input = "2024/day3/input.txt"

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

	part1 := 0
	part2 := 0
	do := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			pattern := regexp.MustCompile(`\bmul\(\d{1,3},\d{1,3}\)|\bdon't(\(\))|\bdo(\(\))`)
			matches := pattern.FindAllString(in, -1)
			for _, match := range matches {
				if strings.HasPrefix(match, "don't") {
					do = false
					continue
				}
				if strings.HasPrefix(match, "do") {
					do = true
					continue
				}

				p := regexp.MustCompile(`\d{1,3}`)
				numbers := p.FindAllString(match, -1)
				a, b := numbers[0], numbers[1]
				aInt, _ := strconv.Atoi(a)
				bInt, _ := strconv.Atoi(b)
				part1 += aInt * bInt

				if do == true {
					part2 += aInt * bInt
				}
			}
		}
	}

	return part1, part2
}
