package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var input = "2023/day7/input.txt"

func main() {
	start := time.Now()

	hands := parse()

	part1 := Play(hands)
	part2 := PlayJokersWild(hands)

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Hand struct {
	Cards []string
	Bid   int
	Score int
}

var CardScore = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
	"W": 0,
}

func parse() []Hand {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hands []Hand

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			bid, _ := strconv.Atoi(p[1])
			cards := strings.Split(p[0], "")
			hands = append(hands, Hand{cards, bid, 0})
		}
	}

	return hands
}

func Play(hands []Hand) int {
	for i, _ := range hands {
		hands[i].Score = calcScore(hands[i].Cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Score == hands[j].Score {
			return LessThan(hands[i].Cards, hands[j].Cards, false)
		} else {
			return hands[i].Score < hands[j].Score
		}
	})

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.Bid
	}

	return totalWinnings
}

func PlayJokersWild(hands []Hand) int {
	for i, _ := range hands {
		hands[i].Score = calcScoreWithJokers(hands[i].Cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Score == hands[j].Score {
			return LessThan(hands[i].Cards, hands[j].Cards, true)
		} else {
			return hands[i].Score < hands[j].Score
		}
	})

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.Bid
	}

	return totalWinnings
}

func LessThan(cardsA []string, cardsB []string, jokers bool) bool {
	for i, _ := range cardsA {
		if cardsA[i] == cardsB[i] {
			continue
		}
		if jokers {
			a := cardsA[i]
			if cardsA[i] == "J" {
				a = "W"
			}
			b := cardsB[i]
			if cardsB[i] == "J" {
				b = "W"
			}
			return CardScore[a] < CardScore[b]
		}
		return CardScore[cardsA[i]] < CardScore[cardsB[i]]
	}
	panic("oh no")
}

func calcScoreWithJokers(cards []string) int {
	jokers := []int{}
	for i, card := range cards {
		if card == "J" {
			jokers = append(jokers, i)
		}
	}
	if len(jokers) == 0 {
		return calcScore(cards)
	}
	score := 0
	for _, card := range cards {
		newCards := make([]string, len(cards))
		copy(newCards, cards)

		if card != "J" {
			for _, joker := range jokers {
				newCards[joker] = card
			}
		}
		newScore := calcScore(newCards)
		if newScore > score {
			score = newScore
		}
	}
	return score
}

func calcScore(cards []string) int {
	cardMap := map[string]int{}
	for _, c := range cards {
		cardMap[c] += 1
	}
	values := []int{}
	for _, val := range cardMap {
		values = append(values, val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	if values[0] == 5 {
		return 7
	}
	if values[0] == 4 {
		return 6
	}
	if values[0] == 3 {
		if values[1] == 2 {
			return 5
		}
		return 4
	}
	if values[0] == 2 {
		if values[1] == 2 {
			return 3
		}
		return 2
	}
	if values[0] == 1 {
		return 1
	}
	return -1
}
