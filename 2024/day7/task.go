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

var input = "2024/day7/input.txt"

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

	part1, part2 := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p1 := strings.Split(in, ": ")
			answer, _ := strconv.Atoi(p1[0])
			values := strings.Split(p1[1], " ")

			part1 += sumUp(values, answer, 2)

			part2 += sumUp(values, answer, 3)
		}
	}

	return part1, part2
}

func sumUp(values []string, answer int, base int) int {
	operators := 0
	maxOperatorsStr := "1" + strings.Repeat("0", len(values)-1)
	maxOperators, _ := strconv.ParseInt(maxOperatorsStr, base, 64)
	for {
		operatorsStr := strconv.FormatInt(int64(operators), base)
		operatorsInt, _ := strconv.ParseInt(operatorsStr, base, 64)
		if operatorsInt == maxOperators {
			break
		}

		paddedOperatorStr := fmt.Sprintf("%s", operatorsStr)
		if len(paddedOperatorStr) < len(values)-1 {
			padding := strings.Repeat("0", len(values)-1-len(paddedOperatorStr))
			paddedOperatorStr = padding + paddedOperatorStr
		}

		sum, _ := strconv.Atoi(values[0])
		for i, op := range paddedOperatorStr {
			if i+1 >= len(values) {
				break
			}
			v, _ := strconv.Atoi(values[i+1])
			if op == '0' {
				sum += v
			}
			if op == '1' {
				sum *= v
			}
			if op == '2' {
				sum, _ = strconv.Atoi(fmt.Sprintf("%d%d", sum, v))
			}
			if sum > answer {
				break
			}
		}
		if answer == sum {
			return answer
		}

		operators++
	}
	return 0
}