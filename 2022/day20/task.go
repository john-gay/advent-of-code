package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var input = "2022/day20/input.txt"

type Task struct {
	items []Item
}

type Item struct {
	value     int
	origIndex int
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
	task1 := readInput()
	part1 := task1.groveCoOrds(1, 1)

	task2 := readInput()
	part2 := task2.groveCoOrds(811589153, 10)

	return part1, part2
}

func (t *Task) groveCoOrds(decryptionKey, times int) int {
	for i := range t.items {
		t.items[i].value *= decryptionKey
	}

	size := len(t.items)
	moves := []int{}
	for i := 0; i < size; i++ {
		moves = append(moves, t.items[i].value)
	}

	for turn := 0; turn < times; turn++ {
		for i := 0; i < len(moves); i++ {
			prevIndex := getIndex(t.items, i)
			prevItem := t.items[prevIndex]
			nextIndex := getNextIndex(prevIndex, moves[i], size)
			t.items = removeIndex(t.items, prevIndex)
			t.items = insertAt(t.items, nextIndex, prevItem)
		}
	}

	zeroIndex := getZero(t.items)
	sum := 0
	for i := 1000; i <= 3000; i += 1000 {
		index := (zeroIndex + i) % size
		fmt.Println(index, t.items[index])
		sum += t.items[index].value
	}
	return sum
}

func getNextIndex(index, move, size int) int {
	newIndex := (index + move) % (size - 1)
	if newIndex <= 0 {
		return (size - 1) + newIndex
	}
	return newIndex
}

func insertAt(items []Item, index int, prev Item) []Item {
	items = append(items[:index+1], items[index:]...)
	items[index] = prev
	return items
}

func removeIndex(items []Item, index int) []Item {
	return append(items[:index], items[index+1:]...)
}

func getIndex(items []Item, index int) int {
	for i, item := range items {
		if item.origIndex == index {
			return i
		}
	}
	return -1
}

func getZero(items []Item) int {
	for i, item := range items {
		if item.value == 0 {
			return i
		}
	}
	return -1
}

func readInput() Task {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var items []Item

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			i, _ := strconv.Atoi(in)
			newItem := Item{value: i, origIndex: count}
			items = append(items, newItem)
			count++
		}
	}

	return Task{items: items}
}
