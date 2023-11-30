package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var input = "day21/input.txt"

type Task struct {
	monkeys map[string]any
}

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))

	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (int, int) {
	task := readInput()

	part1 := task.monkeyValue("root")

	task.monkeys["root"] = []string{task.monkeys["root"].([]string)[0], "=", task.monkeys["root"].([]string)[2]}
	task.monkeys["humn"] = "x"

	part2 := task.gradientDescent()

	return part1, part2
}

func (t *Task) monkeyValue(name string) int {
	if reflect.TypeOf(t.monkeys[name]).String() == "int" {
		return t.monkeys[name].(int)
	} else {
		monkeyCalc := t.monkeys[name].([]string)
		switch monkeyCalc[1] {
		case "+":
			return t.monkeyValue(monkeyCalc[0]) + t.monkeyValue(monkeyCalc[2])
		case "-":
			return t.monkeyValue(monkeyCalc[0]) - t.monkeyValue(monkeyCalc[2])
		case "*":
			return t.monkeyValue(monkeyCalc[0]) * t.monkeyValue(monkeyCalc[2])
		case "/":
			return t.monkeyValue(monkeyCalc[0]) / t.monkeyValue(monkeyCalc[2])
		}
	}
	panic(fmt.Sprintf("unhandled: %s", name))
}

func (t *Task) gradientDescent() int {
	result := 1
	t.monkeys["humn"] = 1
	for result != 0 {
		monkeyCalc := t.monkeys["root"].([]string)
		result = t.monkeyValue(monkeyCalc[2]) - t.monkeyValue(monkeyCalc[0])
		t.monkeys["humn"] = t.monkeys["humn"].(int) + int(-0.1*float64(result))
	}
	return t.monkeys["humn"].(int)
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	task := Task{
		monkeys: map[string]any{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			parts := strings.Split(in, ": ")
			name := parts[0]
			value := strings.Split(parts[1], " ")
			if len(value) == 1 {
				intValue, _ := strconv.Atoi(value[0])
				task.monkeys[name] = intValue
			} else {
				task.monkeys[name] = value
			}
		}
	}

	return task
}
