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

var input = "2023/day12/input.txt"

type task struct {
	rows []row
}

type row struct {
	seq   []string
	rules []int
	cache map[coord]int
}

type coord struct {
	i int
	j int
}

func main() {
	start := time.Now()

	t := parse(false)
	part1 := t.possibleArrangements()

	t2 := parse(true)
	part2 := t2.possibleArrangements()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (r *row) dp(i, j int) int {
	if i >= len(r.seq) {
		if j < len(r.rules) {
			return 0
		}
		return 1
	}

	_, ok := r.cache[coord{i, j}]
	if ok {
		return r.cache[coord{i, j}]
	}

	res := 0
	if r.seq[i] == "." {
		res = r.dp(i+1, j)
	} else {
		if r.seq[i] == "?" {
			res = r.dp(i+1, j)
		}
		if j < len(r.rules) {
			count := 0
			for k := i; k < len(r.seq); k++ {
				if count == r.rules[j] && r.seq[k] == "?" {
					break
				}
				if count > r.rules[j] || r.seq[k] == "." {
					break
				}
				count += 1
			}

			if count == r.rules[j] {
				skip := i + count
				if skip < len(r.seq) && r.seq[skip] != "#" {
					res += r.dp(skip+1, j+1)
				} else {
					res += r.dp(skip, j+1)
				}
			}
		}
	}

	r.cache[coord{i, j}] = res
	return res
}

func (t *task) possibleArrangements() int {
	sum := 0
	for _, r := range t.rows {
		sum += r.dp(0, 0)
	}

	return sum
}

func parse(unfold bool) task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := task{
		[]row{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			r := row{cache: map[coord]int{}}
			p := strings.Split(in, " ")
			for _, c := range p[0] {
				r.seq = append(r.seq, fmt.Sprintf("%c", c))
			}

			for _, c := range strings.Split(p[1], ",") {
				rule, _ := strconv.Atoi(c)
				r.rules = append(r.rules, rule)
			}
			if unfold {
				newSeq := []string{}
				newRules := []int{}
				for i := 0; i < 5; i++ {
					newSeq = append(newSeq, r.seq...)
					newRules = append(newRules, r.rules...)
					if i != 4 {
						newSeq = append(newSeq, "?")
					}
				}
				r.seq = newSeq
				r.rules = newRules
			}
			t.rows = append(t.rows, r)
		}
	}

	return t
}
