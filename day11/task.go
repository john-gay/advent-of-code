package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var input = "day11/input.txt"

type Monkey struct {	
	items []int
	operation func(old int) int
	test func(item int) int
	base int
	inspected int
}

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	p1 := play(getMonkeys(), 20, 3)
	p2 := play(getMonkeys(), 10000, 1)
	
	return p1, p2
}

func play(monkeys map[int]*Monkey, turns, divisor int) int {
	commonMultiple := 1
	for i := 0; i < len(monkeys); i++ {
		commonMultiple *= monkeys[i].base
	}

	for turn := 0; turn<turns; turn++ {
		for i := 0; i < len(monkeys); i++ {
			size := len(monkeys[i].items)
			for j := 0; j < size; j++ {
				monkeys[i].inspected++
				worry := monkeys[i].operation(monkeys[i].items[0])
				monkeys[i].items = monkeys[i].items[1:]
				worry = int((worry % commonMultiple) / divisor)
				nextMonkey := monkeys[i].test(worry)
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, worry)
			}
		}
	}

	inspectList := []int{}
	for _, monkey := range monkeys {
		inspectList = append(inspectList, monkey.inspected)
	}
	sort.Slice(inspectList, func(i, j int) bool {
		return inspectList[i] > inspectList[j]
	})

	return inspectList[0] * inspectList[1]
}

func getMonkeys() map[int]*Monkey {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	monkeys := map[int]*Monkey{}
	monkey := &Monkey{}

	var test int
	var trueMonkey int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(strings.Trim(in, " "), " ")
			switch p[0] {
			case "Monkey":
				num, _ := strconv.Atoi(p[1][:1])
				monkeys[num] = monkey
			case "Starting":
				list := strings.Split(in, ": ")
				items := strings.Split(list[1], ", ")
				for _, item := range items {
					n, _ := strconv.Atoi(item)
					monkey.items = append(monkey.items, n)
				}
			case "Operation:":
				if p[4] == "+" {
					if p[5] == "old" {
						monkey.operation = func(old int) int {
							return old + old
						}
					} else {
						n, _ := strconv.Atoi(p[5])
						monkey.operation = func(old int) int {
							return old + n
						}
					}
				} else {
					if p[5] == "old" {
						monkey.operation = func(old int) int {
							return old * old
						}
					} else {
						n, _ := strconv.Atoi(p[5])
						monkey.operation = func(old int) int {
							return old * n
						}
					}
				}
			case "Test:":
				n, _ := strconv.Atoi(p[3])
				test = n
			case "If":
				if p[1] == "true:" {
					nt, _ := strconv.Atoi(p[5])
					trueMonkey = nt
				} else {
					nf, _ := strconv.Atoi(p[5])
					tm := trueMonkey
					t := test
					monkey.test = func(item int) int {
						if item % t == 0 {
							return tm
						} else {
							return nf
						}
					}
					monkey.base = t
				}
			}
		} else {
			monkey = &Monkey{}
		}
	}
	return monkeys
}
