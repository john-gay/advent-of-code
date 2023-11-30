package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var input = "day16/input.txt"

type Task struct {
	valves  map[string]bool
	tunnels map[string][]string
	flow    map[string]int
}

type QueueItem struct {
	path     []string
	paths    [][]string
	pressure int
	opened   map[string]bool
}

type Routes struct {
	path   []string
	opened map[string]bool
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
	t := readInput()

	part1 := t.findHighestPressure()

	part2 := t.findHighestPressureWithElephant()

	return part1, part2
}

func (t *Task) findHighestPressureWithElephant() int {
	queue := []QueueItem{
		{
			paths:    [][]string{{"AA"}, {"AA"}},
			pressure: 0,
			opened:   make(map[string]bool),
		},
	}

	for minute := 1; minute <= 26; minute++ {
		if len(queue) > 10000 {
			sort.Slice(queue, func(i, j int) bool {
				return queue[i].pressure > queue[j].pressure
			})
			queue = queue[:10000]
		}

		var nextQueue []QueueItem
		for _, item := range queue {
			location := item.paths[0][len(item.paths[0])-1]
			elephant := item.paths[1][len(item.paths[1])-1]

			currentPressure := item.pressure
			for open := range item.opened {
				currentPressure += t.flow[open]
			}

			myRoutes := t.getPossibleRoutes(location, item.opened, item.paths[0])
			for _, myRoute := range myRoutes {
				eleRoutes := t.getPossibleRoutes(elephant, myRoute.opened, item.paths[1])
				for _, eleRoute := range eleRoutes {
					paths := [][]string{myRoute.path, eleRoute.path}
					nextQueue = append(nextQueue, QueueItem{
						paths:    paths,
						pressure: currentPressure,
						opened:   eleRoute.opened,
					})
				}
			}
		}
		queue = nextQueue
	}

	return calcMaxPressure(queue)
}

func (t *Task) findHighestPressure() int {
	queue := []QueueItem{
		{
			path:     []string{"AA"},
			pressure: 0,
			opened:   make(map[string]bool),
		},
	}

	for minute := 1; minute <= 30; minute++ {
		if len(queue) > 5000 {
			sort.Slice(queue, func(i, j int) bool {
				return queue[i].pressure > queue[j].pressure
			})
			queue = queue[:5000]
		}

		var nextQueue []QueueItem
		for _, item := range queue {
			location := item.path[len(item.path)-1]
			currentPressure := item.pressure
			for open := range item.opened {
				currentPressure += t.flow[open]
			}

			routes := t.getPossibleRoutes(location, item.opened, item.path)
			for _, route := range routes {
				nextQueue = append(nextQueue, QueueItem{
					path:     route.path,
					pressure: currentPressure,
					opened:   route.opened,
				})
			}
		}
		queue = nextQueue
	}

	return calcMaxPressure(queue)
}

func (t *Task) getPossibleRoutes(location string, opened map[string]bool, path []string) []Routes {
	routes := []Routes{}
	for _, tunnel := range t.tunnels[location] {
		nextOpened := make(map[string]bool)
		for k, v := range opened {
			nextOpened[k] = v
		}
		var nextPath []string
		for _, p := range path {
			nextPath = append(nextPath, p)
		}
		nextPath = append(nextPath, tunnel)
		routes = append(routes, Routes{
			path:   nextPath,
			opened: nextOpened,
		})
	}

	if t.flow[location] > 0 && !opened[location] {
		nextOpened := make(map[string]bool)
		for k, v := range opened {
			nextOpened[k] = v
		}
		nextOpened[location] = true
		var nextPath []string
		for _, p := range path {
			nextPath = append(nextPath, p)
		}
		routes = append(routes, Routes{
			path:   nextPath,
			opened: nextOpened,
		})
	}
	return routes
}

func calcMaxPressure(queue []QueueItem) int {
	maxPressure := 0
	for _, item := range queue {
		if maxPressure < item.pressure {
			maxPressure = item.pressure
		}
	}

	return maxPressure
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := Task{
		valves:  make(map[string]bool),
		tunnels: make(map[string][]string),
		flow:    make(map[string]int),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			var name, to string
			var flow int
			_, _ = fmt.Sscanf(in, "Valve %s has flow rate=%d; tunnel leads to valve %s", &name, &flow, &to)
			t.valves[name] = true
			t.flow[name] = flow
			t.tunnels[name] = []string{to}
			if to == "" {
				var toList []string
				to = strings.Split(in, "to valves ")[1]
				for _, t := range strings.Split(to, ", ") {
					toList = append(toList, t)
				}
				t.tunnels[name] = toList
			}

		}
	}

	return t
}
