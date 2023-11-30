package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

var input = "day25/input.txt"

func main() {
	start := time.Now()

	part1 := run()

	log.Println(fmt.Sprintf("Part 1: %s", part1))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() string {
	snafus := readInput()

	part1 := snafuSum(snafus)

	return part1
}

func snafuSum(snafus []string) string {
	decimalSum := 0
	for _, snafu := range snafus {
		decimal := snafuToDecimal(snafu)
		decimalSum += decimal
	}

	return decimalToSnafu(decimalSum)
}

func decimalToSnafu(decimal int) string {
	snafu := ""
	for decimal > 0 {
		switch (decimal + 2) % 5 {
		case 4:
			snafu = "2" + snafu
		case 3:
			snafu = "1" + snafu
		case 2:
			snafu = "0" + snafu
		case 1:
			snafu = "-" + snafu
		case 0:
			snafu = "=" + snafu
		}
		decimal = (decimal + 2) / 5
	}
	return snafu
}

func snafuToDecimal(snafu string) int {
	sum := 0
	size := len(snafu)
	for i := size - 1; i >= 0; i-- {
		pow := int(math.Pow(5, float64(size-i-1)))
		switch string(snafu[i]) {
		case "2":
			sum += pow * 2
		case "1":
			sum += pow * 1
		case "0":
			sum += pow * 0
		case "-":
			sum += pow * -1
		case "=":
			sum += pow * -2
		}
	}
	return sum
}

func readInput() []string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	snafus := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			snafus = append(snafus, in)
		}
	}
	return snafus
}
