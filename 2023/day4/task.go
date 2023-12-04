package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var input = "2023/day4/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	cardMap := map[int]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, ":")
			cardParts := strings.Split(p[0], " ")
			cardNum, _ := strconv.Atoi(cardParts[len(cardParts)-1])
			cardMap[cardNum] += 1
			matches := 0
			cards := strings.Split(p[1], "|")
			winningNums := strings.Split(strings.Trim(cards[0], " "), " ")
			numbers := strings.Split(strings.Trim(cards[1], " "), " ")
		winning:
			for _, winning := range winningNums {
				for _, number := range numbers {
					if number != "" && number == winning {
						matches++
						continue winning
					}
				}
			}
			if matches > 0 {
				total += int(math.Pow(2, float64(matches-1)))
				for i := 1; i <= matches; i++ {
					cardMap[cardNum+i] += cardMap[cardNum]
				}
			}
		}
	}
	totalCards := 0
	for _, value := range cardMap {
		totalCards += value
	}

	return total, totalCards
}
