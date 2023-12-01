package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("2022/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxCals := []int{0, 0, 0}
	current := int(0)

	for scanner.Scan() {
		p := scanner.Text()
		if p != "" {
			c, _ := strconv.ParseInt(p, 10, 32)
			current += int(c)
		} else {
			if maxCals[0] < current {
				maxCals[0] = current
			}
			sort.Ints(maxCals)

			current = 0
		}
	}
	log.Println(fmt.Sprintf("Part 1: %d", maxCals[2]))
	log.Println(fmt.Sprintf("Part 2: %d", maxCals[0]+maxCals[1]+maxCals[2]))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
