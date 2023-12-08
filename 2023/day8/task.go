package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var input = "2023/day8/input.txt"

func main() {
	start := time.Now()

	t := parse()

	part1 := t.followDirections("AAA", "ZZZ")

	ghostNodes := findGhostNodes(t.nextStep)
	part2 := t.followAllDirections(ghostNodes)

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type task struct {
	directions []string
	nextStep   map[string]choice
}

type choice struct {
	left  string
	right string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lowestCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = lowestCommonMultiple(result, integers[i])
	}

	return result
}

func findGhostNodes(steps map[string]choice) []string {
	nodes := []string{}
	for key, _ := range steps {
		chars := strings.Split(key, "")
		if chars[len(chars)-1] == "A" {
			nodes = append(nodes, key)
		}
	}
	return nodes
}

func (t *task) followAllDirections(nodes []string) int {
	stepList := []int{}
	for _, node := range nodes {
		stepList = append(stepList, t.followDirections(node, "Z"))
	}

	return lowestCommonMultiple(stepList[0], stepList[1], stepList...)
}

func (t *task) followDirections(start, end string) int {
	pos := start
	steps := 0
	for {
		for _, d := range t.directions {
			if strings.HasSuffix(pos, end) {
				return steps
			}
			if d == "L" {
				pos = t.nextStep[pos].left
			} else {
				pos = t.nextStep[pos].right
			}
			steps++
		}
	}
}

func parse() task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{nextStep: map[string]choice{}}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			if t.directions == nil {
				t.directions = strings.Split(in, "")
				continue
			}
			p := strings.Split(in, " = ")
			key := p[0]
			p2 := strings.Split(p[1], ", ")
			left := strings.Split(p2[0], "(")[1]
			right := strings.Split(p2[1], ")")[0]
			t.nextStep[key] = choice{left, right}
		}
	}

	return t
}
