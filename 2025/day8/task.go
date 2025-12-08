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

var input = "2025/day8/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Coord struct {
	x, y, z int
}

type Edge struct {
	a, b Coord
	dist int
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	coords := []Coord{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, ",")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			z, _ := strconv.Atoi(p[2])
			coords = append(coords, Coord{x, y, z})
		}
	}

	edges := []Edge{}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			dist := int(euclideanDist(coords[i], coords[j]))
			edges = append(edges, Edge{coords[i], coords[j], dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	parent := map[Coord]Coord{}
	size := map[Coord]int{}
	for _, c := range coords {
		parent[c] = c
		size[c] = 1
	}

	var lastEdge Edge
	connections := 0
	part1 := 0
	
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		
		if i == 1000 {
			circuits := map[Coord]int{}
			for _, c := range coords {
				root := find(parent, c)
				circuits[root] = size[root]
			}
			sizes := []int{}
			for _, s := range circuits {
				sizes = append(sizes, s)
			}
			sort.Slice(sizes, func(i, j int) bool {
				return sizes[i] > sizes[j]
			})
			part1 = sizes[0] * sizes[1] * sizes[2]
		}
		
		rootA := find(parent, edge.a)
		rootB := find(parent, edge.b)
		if rootA == rootB {
			continue
		}
		if size[rootA] < size[rootB] {
			parent[rootA] = rootB
			size[rootB] += size[rootA]
		} else {
			parent[rootB] = rootA
			size[rootA] += size[rootB]
		}
		connections++
		lastEdge = edge
		
		root := find(parent, coords[0])
		if size[root] == len(coords) {
			break
		}
	}

	part2 := lastEdge.a.x * lastEdge.b.x

	return part1, part2
}

func euclideanDist(a, b Coord) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return dx*dx + dy*dy + dz*dz
}

func find(parent map[Coord]Coord, c Coord) Coord {
	root := c
	for parent[root] != root {
		root = parent[root]
	}
	for parent[c] != c {
		next := parent[c]
		parent[c] = root
		c = next
	}
	return root
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}