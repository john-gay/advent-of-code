package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var input = "2022/day19/input.txt"

type Blueprint struct {
	number int
	robots map[string]Robot
	max    map[string]int
}

type Robot struct {
	costs map[string]int
}

type State struct {
	minute    int
	construct string
	robots    map[string]int
	collected map[string]int
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
	blueprints := readInput()

	part1 := totalQualityLevel(blueprints, 24)

	part2 := productQualityLevel(blueprints, 32)

	return part1, part2
}

func productQualityLevel(blueprints []Blueprint, minutes int) int {
	qualityLevel := 1

	for _, blueprint := range blueprints[:3] {
		maxGeodes := getMaxGeodes(minutes, blueprint)

		fmt.Println(qualityLevel, maxGeodes, blueprint.number)
		qualityLevel *= maxGeodes
	}

	return qualityLevel
}

func totalQualityLevel(blueprints []Blueprint, minutes int) int {
	qualityLevel := 0

	for _, blueprint := range blueprints {
		maxGeodes := getMaxGeodes(minutes, blueprint)

		fmt.Println(qualityLevel, maxGeodes, blueprint.number)
		qualityLevel += maxGeodes * blueprint.number
	}

	return qualityLevel
}

func getMaxGeodes(minutes int, blueprint Blueprint) int {
	queue := []State{{
		minute:    0,
		construct: "",
		robots:    map[string]int{"ore": 1},
		collected: map[string]int{},
	}}
	checked := map[string]bool{}

	maxGeodes := 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.minute > minutes {
			continue
		}

		if maxGeodes < state.collected["geode"] {
			maxGeodes = state.collected["geode"]
		}

		key := state.getKey()
		if checked[key] {
			continue
		}
		checked[key] = true

		for _, construct := range []string{"geode", "obsidian", "clay", "ore"} {
			state.collected[construct] += state.robots[construct]
		}

		if state.construct != "" {
			state.robots[state.construct]++
			state.construct = ""
		}

		nextStates := nextSteps(blueprint, state)

		queue = append(queue, nextStates...)
	}
	return maxGeodes
}

func nextSteps(blueprint Blueprint, state State) []State {
	states := []State{}
	for _, construct := range []string{"geode", "obsidian", "clay", "ore"} {
		canBuild := true
		for required, cost := range blueprint.robots[construct].costs {
			if state.collected[required] < cost {
				canBuild = false
			}
		}
		should := state.shouldBuild(blueprint, construct)
		if canBuild && should {
			nextCollected := copyMap(state.collected)
			for required, cost := range blueprint.robots[construct].costs {
				nextCollected[required] -= cost
			}
			states = append(states, State{
				minute:    state.minute + 1,
				construct: construct,
				robots:    copyMap(state.robots),
				collected: nextCollected,
			})
			if construct == "geode" || construct == "obsidian" {
				break
			}
		}
	}

	if !(len(states) == 1 && states[0].construct == "geode") {
		emptyState := state
		emptyState.construct = ""
		emptyState.minute++
		states = append(states, emptyState)
	}

	return states
}

func copyMap[K comparable, V any](input map[K]V) map[K]V {
	output := make(map[K]V)
	for k, v := range input {
		output[k] = v
	}
	return output
}

func readInput() []Blueprint {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var blueprints []Blueprint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			blueprint := Blueprint{
				robots: map[string]Robot{},
				max:    map[string]int{},
			}
			var ore, clay, obsidianOre, obsidianClay, geodeOre, geodeObsidian int
			_, _ = fmt.Sscanf(in, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
				&blueprint.number, &ore, &clay, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)
			blueprint.robots["ore"] = Robot{map[string]int{"ore": ore}}
			blueprint.robots["clay"] = Robot{map[string]int{"ore": clay}}
			blueprint.robots["obsidian"] = Robot{map[string]int{"ore": obsidianOre, "clay": obsidianClay}}
			blueprint.robots["geode"] = Robot{map[string]int{"ore": geodeOre, "obsidian": geodeObsidian}}

			blueprint.max["ore"] = maximum(ore, obsidianOre, geodeOre)
			blueprint.max["clay"] = maximum(clay, obsidianClay)
			blueprint.max["obsidian"] = geodeObsidian

			blueprints = append(blueprints, blueprint)
		}
	}

	return blueprints
}

func maximum(integers ...int) int {
	max := integers[0]

	for _, i := range integers {
		if max < i {
			max = i
		}
	}

	return max
}

func (s *State) shouldBuild(blueprint Blueprint, construct string) bool {
	if construct == "geode" {
		return true
	}

	return s.robots[construct] < blueprint.max[construct]
}

func (s *State) getKey() string {
	return fmt.Sprintf("m: %d, c: %s, or: %d, cr: %d, obr: %d, gr: %d, ore %d, clay %d, obs: %d, geode: %d",
		s.minute,
		s.construct,
		s.robots["ore"],
		s.robots["clay"],
		s.robots["obsidian"],
		s.robots["geode"],
		s.collected["ore"],
		s.collected["clay"],
		s.collected["obsidian"],
		s.collected["geode"],
	)
}
