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

var input = "day7/input.txt"

type Dir struct {
	name        string
	parent      *Dir
	directories []*Dir
	files       []File
	totalSize   int
}

type File struct {
	name string
	size int
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

	dir := createDirectoryStructure()

	freeSpace := 70000000 - dir.totalSize
	toFreeSpace := 30000000 - freeSpace

	return totalLessThan(&dir, 100000), directoryToRemove(&dir, toFreeSpace, dir.totalSize)
}

func createDirectoryStructure() Dir {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := Dir{name: "/"}

	currentDir := &root

	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			p := strings.Split(in, " ")
			if p[0] == "$" {
				if p[1] == "cd" {
					switch p[2] {
					case "/":
						currentDir = &root
					case "..":
						currentDir = currentDir.parent
					default:
						for _, dir := range currentDir.directories {
							if dir.name == p[2] {
								currentDir = dir
							}
						}
					}
				}
			} else {
				switch p[0] {
				case "dir":
					newDir := &Dir{name: p[1], parent: currentDir}
					currentDir.directories = append(currentDir.directories, newDir)
				default:
					size, _ := strconv.Atoi(p[0])
					currentDir.files = append(currentDir.files, File{name: p[1], size: size})
					increaseSize(currentDir, size)
				}
			}
		}
	}
	return root
}

func increaseSize(dir *Dir, size int) {
	dir.totalSize += size
	if dir.parent != nil {
		increaseSize(dir.parent, size)
	}
}

func totalLessThan(dir *Dir, limit int) int {
	sum := 0

	if dir.totalSize <= limit {
		sum += dir.totalSize
	}

	for _, d := range dir.directories {
		size := totalLessThan(d, limit)
		sum += size
	}

	return sum
}

func directoryToRemove(dir *Dir, totalSpace, spaceRequired int) int {
	if dir.totalSize >= totalSpace && dir.totalSize < spaceRequired {
		spaceRequired = dir.totalSize
	}

	for _, d := range dir.directories {
		spaceRequired = directoryToRemove(d, totalSpace, spaceRequired)
	}

	return spaceRequired
}
