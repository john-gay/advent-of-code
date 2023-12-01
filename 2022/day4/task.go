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

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	file, err := os.Open("2022/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count1 := 0
	count2 := 0

	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			p := strings.Split(t, ",")
			e1 := strings.Split(p[0], "-")
			e1p1, _ := strconv.ParseInt(e1[0], 10, 64)
			e1p2, _ := strconv.ParseInt(e1[1], 10, 64)

			e2 := strings.Split(p[1], "-")
			e2p1, _ := strconv.ParseInt(e2[0], 10, 64)
			e2p2, _ := strconv.ParseInt(e2[1], 10, 64)

			if (e1p1 <= e2p1 && e1p2 >= e2p2) || (e1p1 >= e2p1 && e1p2 <= e2p2) {
				count1++
			}

			if !(e1p2 < e2p1 || e1p1 > e2p2) {
				count2++
			}
		}
	}

	return count1, count2
}
