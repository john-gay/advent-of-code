package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	p1Count := 0
	p2Count := 0

	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			mm := t[len(t)-1:]
			switch mm {
			case "X":
				p1Count += 1
			case "Y":
				p1Count += 2
			case "Z":
				p1Count += 3
			}

			switch t {
			case "A X", "B Y", "C Z":
				p1Count += 3
			case "A Y", "B Z", "C X":
				p1Count += 6
			}

			switch t {
			case "A X":
				p2Count += 3
			case "A Y":
				p2Count += 4
			case "A Z":
				p2Count += 8
			case "B X":
				p2Count += 1
			case "B Y":
				p2Count += 5
			case "B Z":
				p2Count += 9
			case "C X":
				p2Count += 2
			case "C Y":
				p2Count += 6
			case "C Z":
				p2Count += 7
			}
		}
	}
	log.Println(fmt.Sprintf("Part 1: %d", p1Count))
	log.Println(fmt.Sprintf("Part 2: %d", p2Count))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
