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

var input = "2025/day9/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Coord struct {
	x, y int
}

type Edge struct {
	start, stop Coord
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := map[Coord]bool{}
	edges := []Edge{}
	polygon := []Coord{}

	scanner := bufio.NewScanner(file)
	previousCoord := Coord{x: -1, y: -1}
	firstCoord := Coord{x: -1, y: -1}
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, ",")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			coord := Coord{x: x, y: y}
			grid[coord] = true
			polygon = append(polygon, coord)

			if previousCoord.x != -1 {
				edges = append(edges, Edge{start: previousCoord, stop: coord})
			} else {
				firstCoord = coord
			}

			previousCoord = coord
		}
	}
	edges = append(edges, Edge{start: previousCoord, stop: firstCoord})

	part1, part2 := 0, 0

	largestArea := 0
	for coordI := range grid {
		for coordJ := range grid {
			if coordI == coordJ { 
				continue
			}

			dX := positiveDiff(coordI.x, coordJ.x) + 1
			dY := positiveDiff(coordI.y, coordJ.y) + 1
			area := dX * dY

			if (area > largestArea) {
				largestArea = area
			}
		}
	}

	part1 = largestArea

	largestArea = 0
	for coordI := range grid {
		for coordJ := range grid {
			if coordI == coordJ { 
				continue
			}

			minX, maxX := min(coordI.x, coordJ.x), max(coordI.x, coordJ.x)
			minY, maxY := min(coordI.y, coordJ.y), max(coordI.y, coordJ.y)
			
			corners := []Coord{
				{minX, minY},
				{maxX, minY},
				{minX, maxY},
				{maxX, maxY},
			}
			allCornersValid := true
			for _, corner := range corners {
				if !isInsideOrOnPolygon(corner, polygon, edges) {
					allCornersValid = false
					break
				}
			}
			if !allCornersValid {
				continue
			}

			if crossesEdge(coordI, coordJ, edges) {
				continue
			}

			hasOtherVertex := false
			for coord := range grid {
				if coord == coordI || coord == coordJ {
					continue
				}
				if coord.x > minX && coord.x < maxX && coord.y > minY && coord.y < maxY {
					hasOtherVertex = true
					break
				}
			}
			if hasOtherVertex {
				continue
			}

			dX := positiveDiff(coordI.x, coordJ.x) + 1
			dY := positiveDiff(coordI.y, coordJ.y) + 1
			area := dX * dY

			if (area > largestArea) {
				largestArea = area
			}
		}
	}

	part2 = largestArea
	
	return part1, part2
}

func crossesEdge(a, b Coord, edges []Edge) bool {
	minX, maxX := min(a.x, b.x), max(a.x, b.x)
	minY, maxY := min(a.y, b.y), max(a.y, b.y)
	
	for _, edge := range edges {
		if edgeCrossesRectangle(edge, minX, maxX, minY, maxY) {
			return true
		}
	}
	return false
}

func edgeCrossesRectangle(edge Edge, minX, maxX, minY, maxY int) bool {
	if pointStrictlyInside(edge.start, minX, maxX, minY, maxY) || 
	   pointStrictlyInside(edge.stop, minX, maxX, minY, maxY) {
		return true
	}
	
	edgeMinX, edgeMaxX := min(edge.start.x, edge.stop.x), max(edge.start.x, edge.stop.x)
	edgeMinY, edgeMaxY := min(edge.start.y, edge.stop.y), max(edge.start.y, edge.stop.y)
	if edgeMaxX < minX || edgeMinX > maxX || edgeMaxY < minY || edgeMinY > maxY {
		return false
	}
	
	startOn := pointOnBoundary(edge.start, minX, maxX, minY, maxY)
	stopOn := pointOnBoundary(edge.stop, minX, maxX, minY, maxY)
	
	if startOn && stopOn {
		onTopBottom := (edge.start.y == edge.stop.y) && (edge.start.y == minY || edge.start.y == maxY)
		onLeftRight := (edge.start.x == edge.stop.x) && (edge.start.x == minX || edge.start.x == maxX)
		
		if onTopBottom || onLeftRight {
			return false
		}
		return true
	}
	
	rectangleSides := []Edge{
		{Coord{minX, minY}, Coord{maxX, minY}},
		{Coord{minX, maxY}, Coord{maxX, maxY}},
		{Coord{minX, minY}, Coord{minX, maxY}},
		{Coord{maxX, minY}, Coord{maxX, maxY}},
	}
	
	for _, side := range rectangleSides {
		if segmentsProperlyIntersect(side.start, side.stop, edge.start, edge.stop) {
			return true
		}
	}
	
	return false
}

func pointIsCorner(p Coord, minX, maxX, minY, maxY int) bool {
	return (p.x == minX || p.x == maxX) && (p.y == minY || p.y == maxY)
}

func pointOnBoundary(p Coord, minX, maxX, minY, maxY int) bool {
	onHorizontal := (p.y == minY || p.y == maxY) && p.x >= minX && p.x <= maxX
	onVertical := (p.x == minX || p.x == maxX) && p.y >= minY && p.y <= maxY
	return onHorizontal || onVertical
}

func segmentsProperlyIntersect(p1, p2, q1, q2 Coord) bool {
	if max(p1.x, p2.x) < min(q1.x, q2.x) || min(p1.x, p2.x) > max(q1.x, q2.x) {
		return false
	}
	if max(p1.y, p2.y) < min(q1.y, q2.y) || min(p1.y, p2.y) > max(q1.y, q2.y) {
		return false
	}

	line1Horizontal := p1.y == p2.y
	line2Horizontal := q1.y == q2.y
	if line1Horizontal == line2Horizontal {
		return false
	}
	
	if line1Horizontal {
		horizY := p1.y
		vertX := q1.x
		horizMinX, horizMaxX := min(p1.x, p2.x), max(p1.x, p2.x)
		vertMinY, vertMaxY := min(q1.y, q2.y), max(q1.y, q2.y)
		return horizY > vertMinY && horizY < vertMaxY &&
			   vertX > horizMinX && vertX < horizMaxX
	}
	
	vertX := p1.x
	horizY := q1.y
	vertMinY, vertMaxY := min(p1.y, p2.y), max(p1.y, p2.y)
	horizMinX, horizMaxX := min(q1.x, q2.x), max(q1.x, q2.x)
	return vertX > horizMinX && vertX < horizMaxX &&
		   horizY > vertMinY && horizY < vertMaxY
}

func pointStrictlyInside(p Coord, minX, maxX, minY, maxY int) bool {
	return p.x > minX && p.x < maxX && p.y > minY && p.y < maxY
}

func positiveDiff(a, b int) int {
	if (a < b) {
		return b - a
	}
	return a - b
}

func isInsideOrOnPolygon(point Coord, polygon []Coord, edges []Edge) bool {
	for _, edge := range edges {
		if pointOnEdge(point, edge) {
			return true
		}
	}
	
	for _, vertex := range polygon {
		if point == vertex {
			return true
		}
	}
	
	inside := false
	n := len(polygon)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		if ((polygon[i].y > point.y) != (polygon[j].y > point.y)) &&
			(point.x < (polygon[j].x-polygon[i].x)*(point.y-polygon[i].y)/(polygon[j].y-polygon[i].y)+polygon[i].x) {
			inside = !inside
		}
	}
	return inside
}

func pointOnEdge(point Coord, edge Edge) bool {
	minX, maxX := min(edge.start.x, edge.stop.x), max(edge.start.x, edge.stop.x)
	minY, maxY := min(edge.start.y, edge.stop.y), max(edge.start.y, edge.stop.y)
	
	if point.x < minX || point.x > maxX || point.y < minY || point.y > maxY {
		return false
	}
	
	if edge.start.x == edge.stop.x {
		return point.x == edge.start.x
	} else if edge.start.y == edge.stop.y {
		return point.y == edge.start.y
	}
	
	return false
}

