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

//    [V] [G]             [H]
//[Z] [H] [Z]         [T] [S]
//[P] [D] [F]         [B] [V] [Q]
//[B] [M] [V] [N]     [F] [D] [N]
//[Q] [Q] [D] [F]     [Z] [Z] [P] [M]
//[M] [Z] [R] [D] [Q] [V] [T] [F] [R]
//[D] [L] [H] [G] [F] [Q] [M] [G] [W]
//[N] [C] [Q] [H] [N] [D] [Q] [M] [B]
//1   2   3   4   5   6   7   8   9

type task struct {
	containers [][]string
}

var containers = [][]string{
	{"N", "D", "M", "Q", "B", "P", "Z"},
	{"C", "L", "Z", "Q", "M", "D", "H", "V"},
	{"Q", "H", "R", "D", "V", "F", "Z", "G"},
	{"H", "G", "D", "F", "N"},
	{"N", "F", "Q"},
	{"D", "Q", "V", "Z", "F", "B", "T"},
	{"Q", "M", "T", "Z", "D", "V", "S", "H"},
	{"M", "G", "F", "P", "N", "Q"},
	{"B", "W", "R", "M"},
}

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %s", part1))
	log.Println(fmt.Sprintf("Part 2: %s", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (string, string) {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	t1 := task{
		containers: containers,
	}

	cpy := make([][]string, len(containers))
	copy(cpy, containers)
	t2 := task{
		containers: cpy,
	}

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			parts := strings.Split(text, " ")
			m, _ := strconv.ParseInt(parts[1], 10, 16)
			move := int(m)
			f, _ := strconv.ParseInt(parts[3], 10, 16)
			from := int(f) - 1
			to16, _ := strconv.ParseInt(parts[5], 10, 16)
			to := int(to16) - 1

			t1.PerformMovesFIFO(move, from, to)
			t2.PerformMovesMAO(move, from, to)
		}
	}

	return t1.GetTopContainers(), t2.GetTopContainers()
}

func (t task) PerformMovesFIFO(move, from, to int) {
	for i := 0; i < move; i++ {
		t.containers[to] = append(t.containers[to], t.containers[from][len(t.containers[from])-1])
		t.containers[from] = t.containers[from][:len(t.containers[from])-1]
	}
}

func (t task) PerformMovesMAO(move, from, to int) {
	t.containers[to] = append(t.containers[to], t.containers[from][len(t.containers[from])-move:]...)
	t.containers[from] = t.containers[from][:len(t.containers[from])-move]
}

func (t task) GetTopContainers() string {
	output := ""
	for _, c := range t.containers {
		output += c[len(c)-1]
	}
	return output
}
