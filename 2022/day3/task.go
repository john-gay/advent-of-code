package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

func main() {
	start := time.Now()

	log.Println(fmt.Sprintf("Part 1: %d", part1()))
	log.Println(fmt.Sprintf("Part 2: %d", part2()))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func part1() int {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matching := []rune{}

	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			p1 := t[:len(t)/2]
			p2 := t[len(t)/2:]

			found := []rune{}
			for _, c1 := range p1 {
				for _, c2 := range p2 {
					if string(c1) == string(c2) && !isFound(c1, found) {
						matching = append(matching, c1)
						found = append(found, c1)
					}
				}
			}
		}
	}

	return calcScore(matching)
}

func part2() int {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matching := []rune{}

	groups := [][]string{}
	group := []string{}
	count := 0

	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			if count < 3 {
				group = append(group, t)
				count++
			}
			if count == 3 {
				groups = append(groups, group)
				group = []string{}
				count = 0
			}
		}
	}

	for _, g := range groups {
		matches := []rune{}
		for _, g1 := range g[0] {
			for _, g2 := range g[1] {
				if g1 == g2 && !isFound(g1, matches) {
					matches = append(matches, g1)
				}
			}
		}

		matches2 := []rune{}
		for _, g3 := range g[2] {
			for _, m := range matches {
				if g3 == m && !isFound(g3, matches2) {
					matching = append(matching, g3)
					matches2 = append(matches2, g3)
				}
			}
		}
	}

	return calcScore(matching)
}

func calcScore(matching []rune) int {
	sum := 0
	for _, m := range matching {
		if unicode.IsUpper(m) {
			sum += int(m - 64 + 26)
		} else {
			sum += int(m - 96)
		}
	}
	return sum
}

func isFound(check rune, found []rune) bool {
	for _, f := range found {
		if check == f {
			return true
		}
	}
	return false
}
