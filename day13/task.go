package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

var input = "day13/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	packets := readInput()

	part1 := 0
	for i := 0; i < len(packets); i += 2 {
		correct := isCorrect(packets[i], packets[i+1])
		if correct {
			part1 += (i + 2) / 2
		}
	}

	part2 := findDecoderKey(packets)

	return part1, part2
}

func findDecoderKey(packets []any) int {
	packet2 := parseString("[[2]]")
	packet6 := parseString("[[6]]")
	packets = append(packets, packet2, packet6)
	packets = sort(packets)

	index2 := 0
	index6 := 0
	for i := 0; i < len(packets); i++ {
		if fmt.Sprintf("%v", packets[i]) == fmt.Sprintf("%v", packet2) {
			index2 = i + 1
		}
		if fmt.Sprintf("%v", packets[i]) == fmt.Sprintf("%v", packet6) {
			index6 = i + 1
		}
		if index2 != 0 && index6 != 0 {
			break
		}
	}
	return index2 * index6
}

func sort(packets []any) []any {
	n := len(packets)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if !isCorrect(packets[i], packets[i+1]) {
				temp := packets[i]
				packets[i] = packets[i+1]
				packets[i+1] = temp
				swapped = true
			}
		}
	}
	return packets
}

func isCorrect(left, right any) bool {
	_, correct := checkCorrect(left, right)
	return correct
}

func checkCorrect(left, right any) (bool, bool) {
	if reflect.TypeOf(left).String() == "float64" && reflect.TypeOf(right).String() == "float64" {
		if left.(float64) < right.(float64) {
			return true, true
		} else if left.(float64) > right.(float64) {
			return true, false
		} else {
			return false, false
		}
	} else if reflect.TypeOf(left).String() != "float64" && reflect.TypeOf(right).String() != "float64" {
		if len(left.([]any)) == 0 && len(right.([]any)) == 0 {
			return false, false
		}
		if len(left.([]any)) == 0 {
			return true, true
		}
		if len(right.([]any)) == 0 {
			return true, false
		}

		for i := 0; i < len(left.([]any)); i++ {
			if i == len(right.([]any)) {
				return true, false
			}
			complete, correct := checkCorrect(left.([]any)[i], right.([]any)[i])
			if complete {
				return complete, correct
			}
		}
		if len(left.([]any)) < len(right.([]any)) {
			return true, true
		}
		return false, false
	} else if reflect.TypeOf(left).String() == "float64" {
		return checkCorrect([]any{left}, right)
	} else {
		return checkCorrect(left, []any{right})
	}
}

func readInput() []any {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var packets []any

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			packets = append(packets, parseString(in))
		}
	}
	return packets
}

func parseString(input string) any {
	var output []any
	_ = json.Unmarshal([]byte(input), &output)
	return output
}
