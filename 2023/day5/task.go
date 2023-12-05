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

var input = "2023/day5/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Lookup struct {
	dest   int
	source int
	length int
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seedList := []int{}
	var maps [][]Lookup
	mapIndex := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			dest, err := strconv.Atoi(p[0])
			if err != nil {
				switch p[0] {
				case "seeds:":
					for i := 1; i < len(p); i++ {
						seed, _ := strconv.Atoi(p[i])
						seedList = append(seedList, seed)
					}
				default:
					maps = append(maps, []Lookup{})
					mapIndex++
				}
				continue
			}
			source, _ := strconv.Atoi(p[1])
			length, _ := strconv.Atoi(p[2])
			maps[mapIndex] = append(maps[mapIndex], Lookup{
				dest:   dest,
				source: source,
				length: length,
			})
		}
	}

	part1 := -1
	for _, seed := range seedList {
		part1 = findLowest(seed, maps, part1)
	}

	part2 := -1
	for i := 0; i < len(seedList); i += 2 {
		start := seedList[i]
		end := seedList[i] + seedList[i+1]
		for seed := start; seed < end; seed++ {
			part2 = findLowest(seed, maps, part2)
		}
	}

	return part1, part2
}

func findLowest(seed int, maps [][]Lookup, lowestLocation int) int {
	cur := seed
maps:
	for _, m := range maps {
		for _, lookup := range m {
			if cur >= lookup.source && cur < lookup.source+lookup.length {
				sourceDiff := cur - lookup.source
				cur = lookup.dest + sourceDiff
				continue maps
			}
		}
	}
	if lowestLocation == -1 || cur < lowestLocation {
		lowestLocation = cur
	}
	return lowestLocation
}
