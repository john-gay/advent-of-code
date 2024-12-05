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

var input = "2024/day5/input.txt"

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

	rules := [][]int{}
	pages := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			if strings.Contains(in, "|") {
				rs := strings.Split(in, "|")
				rule := []int{}
				for _, r := range rs {
					rulePart, _ := strconv.Atoi(r)
					rule = append(rule, rulePart)
				}
				rules = append(rules, rule)
			} else {
				ps := strings.Split(in, ",")
				page := []int{}
				for _, p := range ps {
					pagePart, _ := strconv.Atoi(p)
					page = append(page, pagePart)
				}
				pages = append(pages, page)
			}
		}
	}

	for _, page := range pages {
		if isValid(rules, page) {
			part1 += page[(len(page)-1)/2]
		} else {
			valid := false
			for !valid {
				valid = true
				for _, rule := range rules {
					ruleIndex1 := indexOf(rule[0], page)
					ruleIndex2 := indexOf(rule[1], page)
					if ruleIndex1 == -1 || ruleIndex2 == -1 {
						continue
					}
					if ruleIndex1 > ruleIndex2 {
						page[ruleIndex1], page[ruleIndex2] = page[ruleIndex2], page[ruleIndex1]
						valid = false
					}
				}
				if valid {
					break
				}
			}
			part2 += page[(len(page)-1)/2]
		}
	}

	return part1, part2
}

func isValid(rules [][]int, page []int) bool {
	valid := true
	for _, rule := range rules {
		ruleIndex1 := indexOf(rule[0], page)
		ruleIndex2 := indexOf(rule[1], page)

		if ruleIndex1 == -1 || ruleIndex2 == -1 {
			continue
		}

		if ruleIndex1 > ruleIndex2 {
			valid = false
		}
	}
	return valid
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
