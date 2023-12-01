package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("2023/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1, part2 := 0, 0

	for scanner.Scan() {
		t := scanner.Text()
		part1 += sumFirstAndLast(t, false)
		part2 += sumFirstAndLast(t, true)
	}
	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func sumFirstAndLast(t string, textNumbers bool) int {
	f, l := 0, 0
	var err error
	for i, ch := range t {
		f, err = strconv.Atoi(fmt.Sprintf("%c", ch))
		if err != nil {
			if !textNumbers {
				continue
			}
			n, err := checkForTextNumber(t, i)
			if err != nil || n == 0 {
				continue
			}
			f = n
		}
		break
	}
	for i := len(t) - 1; i >= 0; i-- {
		l, err = strconv.Atoi(fmt.Sprintf("%c", t[i]))
		if err != nil {
			if !textNumbers {
				continue
			}
			n, err := checkForTextNumber(t, i)
			if err != nil || n == 0 {
				continue
			}
			l = n
		}
		break
	}
	num, err := strconv.Atoi(fmt.Sprintf("%d%d", f, l))
	if err != nil {
		panic(err)
	}
	return num
}

func checkForTextNumber(t string, i int) (int, error) {
	switch fmt.Sprintf("%c", t[i]) {
	case "o":
		if len(t) > i+2 {
			check := fmt.Sprintf("%c%c%c", t[i], t[i+1], t[i+2])
			if check == "one" {
				return 1, nil
			}
		}
	case "t":
		if len(t) > i+2 {
			check := fmt.Sprintf("%c%c%c", t[i], t[i+1], t[i+2])
			if check == "two" {
				return 2, nil
			}
		}
		if len(t) > i+4 {
			check := fmt.Sprintf("%c%c%c%c%c", t[i], t[i+1], t[i+2], t[i+3], t[i+4])
			if check == "three" {
				return 3, nil
			}
		}
	case "f":
		if len(t) > i+3 {
			check := fmt.Sprintf("%c%c%c%c", t[i], t[i+1], t[i+2], t[i+3])
			if check == "four" {
				return 4, nil
			}
		}
		if len(t) > i+3 {
			check := fmt.Sprintf("%c%c%c%c", t[i], t[i+1], t[i+2], t[i+3])
			if check == "five" {
				return 5, nil
			}
		}
	case "s":
		if len(t) > i+2 {
			check := fmt.Sprintf("%c%c%c", t[i], t[i+1], t[i+2])
			if check == "six" {
				return 6, nil
			}
		}
		if len(t) > i+4 {
			check := fmt.Sprintf("%c%c%c%c%c", t[i], t[i+1], t[i+2], t[i+3], t[i+4])
			if check == "seven" {
				return 7, nil
			}
		}
	case "e":
		if len(t) > i+4 {
			check := fmt.Sprintf("%c%c%c%c%c", t[i], t[i+1], t[i+2], t[i+3], t[i+4])
			if check == "eight" {
				return 8, nil
			}
		}
	case "n":
		if len(t) > i+3 {
			check := fmt.Sprintf("%c%c%c%c", t[i], t[i+1], t[i+2], t[i+3])
			if check == "nine" {
				return 9, nil
			}
		}
	}

	return 0, nil
}
