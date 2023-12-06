package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	part1 := race([]int{52, 94, 75, 94}, []int{426, 1374, 1279, 1216})
	part2 := race([]int{52947594}, []int{426137412791216})

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func race(times, distances []int) int {
	sumWins := 1
	for race, time := range times {
		wins := 0
		haveWon := false
		for t := 1; t < time; t++ {
			if t*(time-t) > distances[race] {
				wins++
				haveWon = true
			} else if haveWon {
				break
			}
		}
		sumWins *= wins
	}

	return sumWins
}
